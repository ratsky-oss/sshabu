/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package del

import (
	"fmt"
	"sshabu/cmd"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var delCmd = &cobra.Command{
	Hidden: true,
	Use:   "del",
	Short: "delete <host> or <group> by name",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {

	cmd.RootCmd.AddCommand(delCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
