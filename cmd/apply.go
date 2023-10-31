/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"sshabu/pkg"
    "os"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Transform .sshabu.yaml to openssh_config",
	Long: `sshabu apply - generate openssh_config according to yaml specification.
	command is going to ask you confirmation before applying`,
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
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
}
