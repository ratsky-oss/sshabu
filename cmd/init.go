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
	Short: "Create default directories",
	Long: `Init command search for  $HOME/.sshabu/ directory.
If no directory found, init will create it and create default $HOME/.sshabu/sshabu.yaml config.`,
	Run: func(cmd *cobra.Command, args []string) {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		if _, err := os.Stat(home+"/.sshabu/"); os.IsNotExist(err) {
			fmt.Println("Creating base paths")
			err = os.MkdirAll(home+"/.sshabu/", 0750)
			cobra.CheckErr(err)
			err = os.WriteFile(home+"/.sshabu/sshabu.yaml", []byte(sshabu.ConfigExample()), 0660)
			fmt.Println("Success ʕ♥ᴥ♥ʔ")
			cobra.CheckErr(err)
			} else {
				fmt.Println("Base sshabu path already exists")
				fmt.Println("Doing nothing ಠ_ಠ")
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
