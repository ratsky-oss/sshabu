/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	// "os"
	"sshabu/pkg"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("apply called")
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
	
	viper.SetConfigType("yaml")  // Set the config file type
    viper.SetConfigFile(".sshabu") // Set the config file name
    if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
        return
    }
	
	var shabu sshabu.Shabu
	err := viper.Unmarshal(&shabu)
	// fmt.Printf("%+v",shabu)
    check(err)
	buf := new(bytes.Buffer)
	err = sshabu.RenderTemplate(shabu, buf)
	check(err)
	fmt.Println(buf.String())
	// shabu := sshabu.Shabu{
	// 	Hosts: []sshabu.Host{
	// 		{
	// 			Name: "ExampleHost",
	// 			Options: &sshabu.Options{
	// 				AddKeysToAgent: true,
	// 				AddressFamily:  "inet",
	// 				BatchMode:      false,
	// 				BindAddress:    "192.168.1.1",
	// 			},
	// 		},
	// 	},
	// 	Groups: []sshabu.Group{
	// 		{
	// 			Name: "ExampleGroup",
	// 			Options: &sshabu.Options{
	// 				AddressFamily:  "inet6",
	// 				BatchMode:      true,
	// 			},
	// 			Hosts: []sshabu.Host{
	// 				{
	// 					Name: "GroupHost1",
	// 					Options: &sshabu.Options{
	// 						AddKeysToAgent: true,
	// 					},
	// 				},
	// 				{
	// 					Name: "GroupHost2",
	// 					Options: &sshabu.Options{
	// 						AddressFamily: "inet",
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }

	


	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// applyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// applyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
