package sshabu

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
    "fmt"
)

func OpensshCheck(cfg string){
    fmt.Println("Verifing...")
    vcmd := exec.Command("bash","-c","ssh -G -F " + cfg)
    vcmd.Stderr = os.Stderr
    if err := vcmd.Run(); err == nil{
        fmt.Println("Seems legit to me")
    }
}

func DestinationHosts(r io.Reader) ([]string, error) {
    scanner := bufio.NewScanner(r)

    // Slice to store values after "Host "
    hostValues := []string{}

    for scanner.Scan() {
        line := scanner.Text()

        // Check if the line starts with "Host " and doesn't contain "*" or "!"
        if strings.HasPrefix(line, "Host ") && !strings.Contains(line, "*") && !strings.Contains(line, "!") {
            hostValue := strings.TrimPrefix(line, "Host ")

            // Split hostValue by spaces and add the resulting entities to hostValues
            entities := strings.Fields(hostValue)
            hostValues = append(hostValues, entities...)
        }
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return hostValues, nil
}

func AskForConfirmation() bool {
    var response string
    _, err := fmt.Scanln(&response)
    if err != nil {
        fmt.Println("Please enter 'yes' or 'no'.")
        return false
    }
    response = strings.ToLower(response)
	switch response {
	case "yes", "y":
		return true
	case "no", "n":
		return false
	default:
		fmt.Println("Please enter 'yes' or 'no'.")
        return false
	}
}