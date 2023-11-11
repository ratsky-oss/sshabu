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
	"fmt"
	"os"
	"os/exec"
	sshabu "sshabu/pkg"

	"github.com/spf13/cobra"
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect [flags] [user@]name_of_host",
	Short: "Just a wrapper around ssh command",
	Long: `Generally just a wrapper around ssh command with autocompletion from sshabu config.

Base usage:
~ sshabu connect some_host
# Command above wll be transformed to the following
# ssh -F $HOME/.sshabu/openssh.config some_host

Optionally you could pass openssh parametrs or override user
~ sshabu connect -o "-p 2222 -i /path/to/dir" user@host_example
# ssh -F $HOME/.sshabu/openssh.config -p 2222 -i /path/to/dir user@host_example
`,
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

		args = append(args, extraOptions)

		sshArgs := append([]string{"-F", opensshDestconfigFile}, args...)

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
