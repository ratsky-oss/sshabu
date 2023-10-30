/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"sshabu/pkg/compare"
	"github.com/spf13/cobra"
	"os"
    "strings"
)


// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("-----------------------------------")
	
    var (
        firstFile compare.Bites
        secondFile compare.Bites
    )

    firstFile.TakeBites("./test.txt")
    secondFile.TakeBites("./test2.txt")
    fmt.Println(secondFile)
	compare.PrintCompareStrings(firstFile, secondFile)	
		
    fmt.Println("\nDo you want to overwrite? (yes/no): ")
    if askForConfirmation() {
        err := os.WriteFile("./test.txt", []byte(strings.Join(secondFile.Content, "\n")), 0644)
        if err != nil {
            fmt.Println("Error overwriting the file:", err)
            return
        }
        fmt.Println("First file overwritten successfully!")
    } else {
        fmt.Println("Action canceled.")
    }

    fmt.Println("-----------------------------------")
},
}

func askForConfirmation() bool {
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


func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
