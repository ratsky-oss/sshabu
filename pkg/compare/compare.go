package compare

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"strings"
	// "sort"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
	Added ChangeType = iota
    Deleted
    Modified
)

type ChangeType int

type Difference struct {
	lineNumber int
    line       string
    changeType ChangeType
}

type Bite struct {
    number  int
    content string
}

type Bites struct {
	length       int
	content      []Bite
}

// External functions
func (bites *Bites) TakeBites(path string) {
	var lineArray []Bite

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++
		if !isWhitespaceOrEmpty(line) {
			bite := Bite{
				number:  lineNumber,
				content: line,
			}
			lineArray = append(lineArray, bite)
		}
	}

	bites.content = lineArray
	bites.length = len(lineArray)
}

func PrintCompareStrings(firstBites Bites, secondBites Bites) {
    differences := diffBites(firstBites, secondBites)
    differences = removeDuplicateDifferences(differences)

    addedCount := 0
    deletedCount := 0
    modifiedCount := 0

    for _, diff := range differences {
        switch diff.changeType {
        case Added:
            addedCount++
        case Deleted:
            deletedCount++
        case Modified:
            modifiedCount++
        }
    }

    // Определяем максимальную длину строки в firstBites.content
    maxStringLen := 0
    for _, bite := range firstBites.content {
        if len(bite.content) > maxStringLen {
            maxStringLen = len(bite.content)
        }
    }

    fmt.Println("-----------")
    fmt.Printf("Total Added Lines: %d\n", addedCount)
    fmt.Printf("Total Deleted Lines: %d\n", deletedCount)
    fmt.Printf("Total Modified Lines: %d\n", modifiedCount/2) // Since a modification is represented by two entries
    fmt.Println("-----------")
    fmt.Println("Changes in lines:")

    for _, diff := range differences {
        switch diff.changeType {
        case Added:
            fmt.Println(strconv.Itoa(diff.lineNumber) + ": " + Green + diff.line + White)
        case Deleted:
            fmt.Println(strconv.Itoa(diff.lineNumber) + ": " + Red + diff.line + White)
        case Modified:
            firstline := ""
            if diff.lineNumber <= len(firstBites.content) {
                firstline = firstBites.content[diff.lineNumber-1].content
            }
            secondline := ""
            if diff.lineNumber <= len(secondBites.content) {
                secondline = secondBites.content[diff.lineNumber-1].content
            }

            fmt.Println(strconv.Itoa(diff.lineNumber) + ": " + Red + firstline + White + strings.Repeat(" ", maxStringLen-len(firstline)) + " --> " + Green + secondline + White)
        }
    }
}

// Internal functions
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func indent(input string, indent int) string {
	padding := indent + len(input)
	return fmt.Sprintf("% "+strconv.Itoa(padding)+"s", input)
}

func diffBites(bites1, bites2 Bites) []Difference {
    var differences []Difference
    i, j := 0, 0

    for i < bites1.length && j < bites2.length {
        if bites1.content[i].content == bites2.content[j].content {
            i++
            j++
            continue
        }

        added := false
        deleted := false

        for k := j; k < bites2.length && !added; k++ {
            if bites1.content[i].content == bites2.content[k].content {
                for l := j; l < k; l++ {
                    differences = append(differences, Difference{lineNumber: bites2.content[l].number, line: bites2.content[l].content, changeType: Added})
                    j++
                }
                added = true
            }
        }

        for k := i; k < bites1.length && !deleted && !added; k++ {
            if bites1.content[k].content == bites2.content[j].content {
                for l := i; l < k; l++ {
                    differences = append(differences, Difference{lineNumber: bites1.content[l].number, line: bites1.content[l].content, changeType: Deleted})
                    i++
                }
                deleted = true
            }
        }

        if !added && !deleted {
            differences = append(differences, Difference{lineNumber: bites1.content[i].number, line: bites1.content[i].content, changeType: Modified})
            i++
            differences = append(differences, Difference{lineNumber: bites2.content[j].number, line: bites2.content[j].content, changeType: Modified})
            j++
        }
    }

    for ; i < bites1.length; i++ {
        differences = append(differences, Difference{lineNumber: bites1.content[i].number, line: bites1.content[i].content, changeType: Deleted})
    }

    for ; j < bites2.length; j++ {
        differences = append(differences, Difference{lineNumber: bites2.content[j].number, line: bites2.content[j].content, changeType: Added})
    }

    return differences
}


func isWhitespaceOrEmpty(s string) bool {
    return len(strings.TrimSpace(s)) == 0
}

func removeDuplicateDifferences(diffs []Difference) []Difference {
    var uniqueDiffs []Difference
    seenModified := make(map[int]bool)

    for _, diff := range diffs {
        if diff.changeType == Modified && seenModified[diff.lineNumber] {
            continue
        }
        if diff.changeType == Modified {
            seenModified[diff.lineNumber] = true
        }
        uniqueDiffs = append(uniqueDiffs, diff)
    }

    return uniqueDiffs
}



