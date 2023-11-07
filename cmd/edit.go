/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
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
	Long: `Edit command sshabu config file with editor.
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
		fmt.Println("卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌")
		fmt.Println("⸫ Using config file:", cfgFile)
		fmt.Println("卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌")
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
		cmd := exec.Command("sshabu", "apply")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Failed to run apply: %v\n", err)
			return
		}
	} else {
		fmt.Println("Ok.(╥﹏╥)")
		fmt.Println("Changes was not applied.")
	}
}


func init() {
	rootCmd.AddCommand(editCmd)
}
