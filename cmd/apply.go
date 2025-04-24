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
	"bytes"
	"fmt"
	"os"
	"sshabu/pkg"
	"sshabu/pkg/compare"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RunApply(args []string) error {
		// Create a fresh command instance
		runApplyCmd := &cobra.Command{
			Use: applyCmd.Use,
			Run: applyCmd.Run,
		}
		
		// Copy all flags
		runApplyCmd.Flags().AddFlagSet(applyCmd.Flags())
		runApplyCmd.SetArgs(args)
		
		err := runApplyCmd.Execute()
		return err
}      
// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Transform sshabu.yaml to openssh.config",
	Long: `Apply the generate openssh_config according to yaml specification.
Command is going to ask you confirmation before overriding destination openssh.config.
openssh.config file is located right next to the used sshabu.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		
		fmt.Println("⸫ Using config file:", cfgFile)
		
		var shabu sshabu.Shabu
		err := viper.UnmarshalExact(&shabu)
		cobra.CheckErr(err)
		// if shabu.AreAllUnique(){
		// 	fmt.Println("YAML seems OK")
		// 	}  else {
		// 	fmt.Println("Error: 'Name' Fields must be unique")
		// 	os.Exit(1)
		// }
		// names := sshabu.FindNamesInShabu(shabu)
		
		err = shabu.Boil()
		cobra.CheckErr(err)

		buf := new(bytes.Buffer)
		err = sshabu.RenderTemplate(shabu, buf)
		cobra.CheckErr(err)

		err = os.WriteFile(opensshTmpFile, buf.Bytes(), 0600)
		cobra.CheckErr(err)
		sshabu.OpensshCheck(opensshTmpFile)
		
		var (
			destFile compare.Bites
			tmpFile compare.Bites
		)
	
		destFile.TakeBites(opensshDestconfigFile)
		tmpFile.TakeBites(opensshTmpFile)

		differences := compare.DiffBites(destFile, tmpFile)

		if len(differences) == 0{
			fmt.Println("---------------------")
			fmt.Println("No changes! ʕっ•ᴥ•ʔっ")
			fmt.Println("---------------------")
			return
		} 
		
		
		
		if !forceApply {
			
			resultStrings := compare.TransformDifferencesToReadableFormat(differences, destFile, tmpFile)
			
			for _,line := range(resultStrings) {
				fmt.Println(line)
			}

			fmt.Println("\nDo you really want to apply changes? (yes/no): ")
			if !sshabu.AskForConfirmation() {
				fmt.Println("Aborted")
				return
			}
		}
		
		err = os.WriteFile(opensshDestconfigFile, []byte(strings.Join(tmpFile.Content, "\n")), 0644)
		os.Remove(opensshTmpFile)
		if err != nil {
			fmt.Println("Error overwriting the file:", err)
			return
		}
		fmt.Println("Yep-Yep-Yep! It's time for shabu! ʕ •́؈•̀)")
	},
}

var forceApply bool

func init() {
	applyCmd.Flags().BoolVarP(&forceApply, "force", "f", false, "Apply configuration without confirmation")
	rootCmd.AddCommand(applyCmd)
}