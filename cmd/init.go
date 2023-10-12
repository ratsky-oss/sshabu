/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	// "sshabu/pkg"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	
	viper.SetConfigName(".sshabu") // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/appname/")   // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname")  // call multiple times to add many search paths
	viper.AddConfigPath(".")               // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		panic(err)
	}
	fmt.Println(viper.AllSettings())
	// startupConfig := sshabu.Shabu{
	// 	Options: sshabu.Options{
	// 		User: "user",
	// 	},
	// 	Hosts: []sshabu.Host{
	// 		sshabu.Host{
	// 			Name: "host1",
	// 			Options: sshabu.Options{Hostname: "192.168.1.1"},
	// 		},
	// 		sshabu.Host{
	// 			Name: "host2",
	// 			Options: sshabu.Options{Hostname: "192.168.1.2"},
	// 		},
	// 	},
	// 	Groups: []sshabu.Group{
	// 		sshabu.Group{

	// 			Name: "general", 
	// 			Options: sshabu.Options{
	// 				IdentityFile: "~/.ssh/id_rsa",
	// 			},
	// 			Hosts: []sshabu.Host{
	// 				sshabu.Host{
	// 					Name: "host3",
	// 					Options: sshabu.Options{Hostname: "192.168.1.3"},
	// 				},
	// 				sshabu.Host{
	// 					Name: "host4",
	// 					Options: sshabu.Options{Hostname: "192.168.1.4"},
	// 				},
	// 			},
	// 			Subgroups: []sshabu.Group{
	// 				sshabu.Group{
	// 					Name: "work", 
	// 					Options: sshabu.Options{
	// 						User: "alivitskiy",
	// 						IdentityFile: "~/.ssh/id_rsa",
	// 					},
	// 					Hosts: []sshabu.Host{
	// 						sshabu.Host{
	// 							Name: "workstation1",
	// 							Options: sshabu.Options{Hostname: "10.0.1.1"},
	// 						},
	// 						sshabu.Host{
	// 							Name: "workstation3",
	// 							Options: sshabu.Options{Hostname: "10.0.1.2"},
	// 						},
	// 					},
	// 					Subgroups: []sshabu.Group{
	// 						sshabu.Group{
								
	// 						},
	// 				},
	// 			},
	// 		},
	// 	},

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
