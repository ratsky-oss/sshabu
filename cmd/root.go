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
	PersistentPreRun: func(cmd *cobra.Command, args []string){
		if cmd.Name() == "completion"{
			return
		}
		initConfig()
	},
	Use:   "sshabu",
	Version: "0.0.1-alpha",
	Short: "Is a robust SSH client management tool",
	Long: `is a robust SSH client management tool designed to streamline the process of connecting to multiple servers effortlessly. 
This tool leverages OpenSSH and offers a user-friendly interface to enhance the overall SSH experience. 
With Sshabu, managing SSH configurations becomes more intuitive, allowing users to organize and connect to their servers efficiently.

Sshabu works with sshabu.yaml and openssh.config file.
openssh.config will be created next to sshabu.yaml

sshabu.yaml location - $HOME (user's home dir)
`,
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
	// cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "manully override config file path")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
		if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.SetConfigType("yaml")
		viper.SetConfigName("sshabu")
		// viper.AddConfigPath("$PWD")
		viper.AddConfigPath(home+"/.sshabu")
	}
	
	if err := viper.ReadInConfig(); err == nil {
		cfgFile = viper.ConfigFileUsed()
		cfgPath := filepath.Dir(cfgFile)
		opensshTmpFile = cfgPath+"/openssh.tmp"
		opensshDestconfigFile = cfgPath+"/openssh.config"
		os.OpenFile(opensshTmpFile, os.O_RDONLY|os.O_CREATE, 0666)
		os.OpenFile(opensshDestconfigFile, os.O_RDONLY|os.O_CREATE, 0666)
		} else {
			fmt.Printf("(╯°□°)╯︵ ɹoɹɹƎ\n%s\n$HOME/.sshabu/sshabu.yaml\n",err)
			os.Exit(1)
	}
}


func SetVersionInfo(version, commit, date string) {
	rootCmd.Version = fmt.Sprintf("%s \nBuilt on %s from Git SHA %s)", version, date, commit)
}