/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	// "fmt"
	// "log"
	_ "embed"
	"fmt"
	"os"
	sshabu "sshabu/pkg"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		if _, err := os.Stat(home+"/.sshabu/"); os.IsNotExist(err) {
			fmt.Println("Creating base paths")
			err = os.MkdirAll(home+"/.sshabu/", 0750)
			cobra.CheckErr(err)
			err = os.WriteFile(home+"/.sshabu/sshabu.yaml", []byte(sshabu.ConfigExample()), 0660)
			fmt.Println("Success")
			cobra.CheckErr(err)
			} else {
				fmt.Println("Base sshabu path already exists")
				fmt.Println("Doing nothing")
		}

		},
	}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
