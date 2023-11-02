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

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "A brief description of your command",
	Long: `...`,
	Run: func(cmd *cobra.Command, args []string) {
		editFile(cfgFile) 
	},
}

func editFile(filePath string) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		fmt.Print("The EDITOR environment variable is not set. Choose an editor [nano/vim]: ")
		reader := bufio.NewReader(os.Stdin)
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		switch choice {
		case "nano":
			editor = "nano"
		case "vim":
			editor = "vim"
		default:
			fmt.Println("Invalid choice. Exiting.")
			return
		}
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

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want to apply changes? [y/n]: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	if strings.ToLower(text) == "y" {
		applyCmd.Execute()
	} else {
		fmt.Println("Changes not applied.")
	}
}


func init() {
	rootCmd.AddCommand(editCmd)
}
