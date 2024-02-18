/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"sshabu/pkg"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

// soakCmd represents the soak command
var soakCmd = &cobra.Command{
	Use:   "soak",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		
		var shabu sshabu.Shabu
		
		// read config and convert it to struct
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		opensshDefConfigFile := home+"/.ssh/config"
		newHosts, err := sshabu.ConvertToShabuStruct(opensshDefConfigFile)
		cobra.CheckErr(err)
		
		// 
		err = viper.UnmarshalExact(&shabu)
		cobra.CheckErr(err)
		err = shabu.Boil()
		cobra.CheckErr(err)
		for _, v := range newHosts {
			_ , err = sshabu.AreKeysInOption(v)
			cobra.CheckErr(err)
			// TODO: Get rid of Params manipulation after ConvertToShabuStruct()
			v["name"] = v["Host"]
			delete(v,"Host")
			host := sshabu.CreateHost(v)
			err = shabu.AddHost(host.(sshabu.Host))
			if err != nil {
				fmt.Println("Warning:", err)
			}
		}

		y, err := yaml.Marshal(shabu)
		if err != nil {
			fmt.Printf("err: %v\n", err) // shabu add host srv-1 -o "Hostname: jkjkjk, "
			return
		}
		fmt.Println(string(y))
	},
}

func init() {
	RootCmd.AddCommand(soakCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// soakCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// soakCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
