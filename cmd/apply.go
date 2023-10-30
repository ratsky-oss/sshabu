/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	// "fmt"
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

		err = shabu.Boil()
		cobra.CheckErr(err)

		buf := new(bytes.Buffer)
		err = sshabu.RenderTemplate(shabu, buf)
		cobra.CheckErr(err)

		// TESTED BY ssh -G -F destination.txt host1 
		err = os.WriteFile(opensshDestconfigFile, buf.Bytes(), 0600)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
}
