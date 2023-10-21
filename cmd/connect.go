/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	sshabu "sshabu/pkg"

	"github.com/spf13/cobra"
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		file, err := os.Open(opensshDestconfigFile)
		if err != nil {
			fmt.Println("Error reading file:", err)
		}
		defer file.Close()

		hostValues, err := sshabu.DestinationHosts(file)
		if err != nil {
			fmt.Println("Error parsing file:", err)
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return hostValues, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("HI MAN")
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
