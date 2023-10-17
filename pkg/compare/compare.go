package compare

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	// "strings"
	"sort"
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
)

type Difference struct {
	lineNumber int
	line       string
}

type Bites struct {
	length       int
	content      []string
}

// External functions
func (bites *Bites) TakeBites(path string){
	var lineArray    []string

	file, err := os.Open(path)
	check(err)
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		lineArray = append(lineArray, scanner.Text())
    }

	bites.content      = lineArray
	bites.length       = len(lineArray)
}

func PrintCompareStrings(firstBites Bites, secondBites Bites) {
	differences := lcsDifference(firstBites, secondBites)
	// Определяем максимальную длину строки в differences
    maxStringLen := 0
    for _, diff := range differences {
        if len(diff.line) > maxStringLen {
            maxStringLen = len(diff.line)
        }
    }
	// Sort the list of differences by line numbers in ascending order
	sort.Slice(differences, func(i, j int) bool {
		return differences[i].lineNumber < differences[j].lineNumber
	})

	fmt.Println("-----------")
	fmt.Println("Changes in lines:")

	for _, diff := range differences {
		if diff.lineNumber <= len(firstBites.content) {
			firstline := firstBites.content[diff.lineNumber-1]
			secondline := ""
			if diff.lineNumber <= len(secondBites.content) {
				secondline = secondBites.content[diff.lineNumber-1]
			}

			fmt.Println(strconv.Itoa(diff.lineNumber) + ": " + Red + firstline + White + indent(" --> ", maxStringLen-len(firstline)) + Green + secondline + White)
		} else {
			fmt.Println(strconv.Itoa(diff.lineNumber) + ": " + Green + diff.line + White)
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

func lcsDifference(bites1, bites2 Bites) []Difference {
	m := bites1.length
	n := bites2.length

	L := make([][]int, m+1)
	for i := range L {
		L[i] = make([]int, n+1)
	}
	
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 || j == 0 {
				L[i][j] = 0
			} else if bites1.content[i-1] == bites2.content[j-1] {
				L[i][j] = L[i-1][j-1] + 1
			} else {
				L[i][j] = max(L[i-1][j], L[i][j-1])
			}
		}
	}

	var differences []Difference
	i, j := m, n
	for i > 0 && j > 0 {
		if bites1.content[i-1] == bites2.content[j-1] {
			// Lines match, just move up diagonally
			i--
			j--
		} else if L[i-1][j] > L[i][j-1] {
			// The line from bites1 is missing in bites2
			differences = append(differences, Difference{lineNumber: i, line: bites1.content[i-1]})
			i--
		} else {
			// The line from bites2 is new compared to bites1
			differences = append(differences, Difference{lineNumber: j, line: bites2.content[j-1]})
			j--
		}
	}

	// If there are leftover lines in bites1
	for i > 0 {
		differences = append(differences, Difference{lineNumber: i, line: bites1.content[i-1]})
		i--
	}

	// If there are leftover lines in bites2
	for j > 0 {
		differences = append(differences, Difference{lineNumber: j, line: bites2.content[j-1]})
		j--
	}

	return differences
}
