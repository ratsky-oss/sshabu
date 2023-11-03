/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
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

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Transform .sshabu.yaml to openssh_config",
	Long: `sshabu apply - generate openssh_config according to yaml specification.
Command is going to ask you confirmation before applying`,
	Run: func(cmd *cobra.Command, args []string) {
		
		var shabu sshabu.Shabu
		err := viper.UnmarshalExact(&shabu)
		cobra.CheckErr(err)
		if shabu.AreAllUnique(){
			fmt.Println("YAML seems OK")
			}  else {
			fmt.Println("Error: 'Name' Fields must be unique")
			os.Exit(1)
		}
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
		compare.PrintCompareStrings(destFile, tmpFile)	
			
		fmt.Println("\nDo you really want to apply changes? (yes/no): ")
		if sshabu.AskForConfirmation() {
			err := os.WriteFile(opensshDestconfigFile, []byte(strings.Join(tmpFile.Content, "\n")), 0644)
			os.Remove(opensshTmpFile)
			if err != nil {
				fmt.Println("Error overwriting the file:", err)
				return
			}
			fmt.Println("Yep-Yep-Yep! Time for shabu!")
		} else {
			fmt.Println("Aborted")
		}
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
}