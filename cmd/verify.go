/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	sshabu "sshabu/pkg"

	// "sshabu/pkg"
	"github.com/spf13/cobra"
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error{
		file, err := os.Open(opensshDestconfigFile)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return err
		}
		defer file.Close()

		hostValues, err := sshabu.DestinationHosts(file)
		if err != nil {
			fmt.Println("Error parsing file:", err)
			return err
		}
		for _, value := range hostValues {
			fmt.Println(value)
			cmd := exec.Command("bash","-c","ssh -G -F .config.tmp " + value)
			if err := cmd.Run(); err != nil {
				// TODO more info about error
				return err
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)
	
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// verifyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// verifyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}