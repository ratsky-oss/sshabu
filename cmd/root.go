/*
Copyright © 2023 alvtsky github.com/Ra-sky
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	// "sshabu/pkg"
)

var cfgFile string
var opensshTmpFile string
var opensshDestconfigFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sshabu",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}


// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sshabu.yaml)")

	opensshDestconfigFile = "/Users/alivitskiy/Documents/Code/sshabu/.config.tmp"
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigType("yaml")
		viper.SetConfigName("sshabu")
		viper.AddConfigPath("$PWD")
		viper.AddConfigPath("$HOME/.sshabu")
	}

	if err := viper.ReadInConfig(); err == nil {
		cfgFile = viper.ConfigFileUsed()
		fmt.Fprintln(os.Stderr, "Using config file:", cfgFile)
		cfgPath := filepath.Dir(cfgFile)
		opensshTmpFile = cfgPath+"/openssh.tmp"
		opensshDestconfigFile = cfgPath+"/openssh.config"
		os.OpenFile(opensshTmpFile, os.O_RDONLY|os.O_CREATE, 0666)
		os.OpenFile(opensshDestconfigFile, os.O_RDONLY|os.O_CREATE, 0666)
	}
}
