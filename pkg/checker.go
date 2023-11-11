// Copyright (C) 2023  Shovra Nikita, Livitsky Andrey

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sshabu

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func OpensshCheck(openssh_cfg string) error {
    fmt.Println("Verifing result...")
    vcmd := exec.Command("bash","-c","ssh -GTF " + openssh_cfg + " test")
    vcmd.Stderr = os.Stderr
    vcmd.Stdin = nil
    if err := vcmd.Run(); err != nil{
        return err
    }
    fmt.Println("Seems legit to me")
    return nil
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