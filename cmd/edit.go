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
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit sshabu config file",
	Long: `Edit the sshabu configuration file with editor.
If no editor command found, ask you to choose between vim and nano.

After editing you will be promted if you'd like to use 'sshabu apply'`,
	Run: func(cmd *cobra.Command, args []string) {
		editFile(cfgFile) 
	},
}

func editFile(filePath string) {
	cmd := exec.Command("editor", filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {

		fmt.Println("⸫ Using config file:", cfgFile)

		editor := ""
		fmt.Println("Editor is not installed.")
		fmt.Println("Choose an editor [nano/vim or press Enter]: ")
		reader := bufio.NewReader(os.Stdin)
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		switch choice {
		case "nano":
			editor = "nano"
		case "vim":
			editor = "vim"
		default:
			fmt.Println("Vim is the right choice!")
			editor = "vim"
		}
		cmd := exec.Command(editor, filePath)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Failed to open editor: %v\n", err)
			return
		}
	}
	

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Would you like sshabu to apply changes? [y/n]: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)


	if strings.ToLower(text) == "y" {
            if err := RunApply([]string{}); err != nil {
                cobra.CheckErr(err)
            }
	} else {
		fmt.Println("Ok.(╥﹏╥)")
		fmt.Println("Changes was not applied.")
	}
}


func init() {
	RootCmd.AddCommand(editCmd)
}
