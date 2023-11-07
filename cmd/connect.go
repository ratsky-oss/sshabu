/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	sshabu "sshabu/pkg"

	"github.com/spf13/cobra"
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Just a wrapper around ssh command",
	Long: `Generally just a wrapper around ssh command with autocompletion from sshabu config`,
ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		file, _ := os.Open(opensshDestconfigFile)
 
		defer file.Close()

		hostValues, err := sshabu.DestinationHosts(file)
		if err != nil {
			file.Close()
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return hostValues, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
			// Construct the ssh command with -I option
		var sshArgs []string
		if len(extraOptions) > 0 {
			sshArgs = append([]string{"-F", opensshDestconfigFile, extraOptions}, args...)
		} else {
			sshArgs = append([]string{"-F", opensshDestconfigFile}, args...)
		}
		
		fmt.Println("Running SSH command:", "ssh", sshArgs)

		// Execute the SSH command
		scmd := exec.Command("ssh", sshArgs...)
		scmd.Stdout = os.Stdout
		scmd.Stderr = os.Stderr
		scmd.Stdin = os.Stdin
		if err := scmd.Run(); err != nil {
			fmt.Println("Error executing SSH command:", err)
			os.Exit(1)
		}
	},
}

var extraOptions string

func init() {
	connectCmd.Flags().StringVarP(&extraOptions, "options", "o", "", "openssh options passed to ssh command")
	rootCmd.AddCommand(connectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
