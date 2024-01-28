/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package add

import (
	"fmt"
	sshabu "sshabu/pkg"
	
	"gopkg.in/yaml.v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var DefOptions string

var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
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
		shabu.AddHost(sshabu.Host{Name: "pipupu", Options: sshabu.Options{}})
		// host_options, _ := cmd.Flags().GetString("options")
		err = shabu.Boil()
		cobra.CheckErr(err)
		y, err := yaml.Marshal(shabu)
		if err != nil {
			fmt.Printf("err: %v\n", err) // shabu add host srv-1 -o "Hostname: jkjkjk, "
			return
		}
		// y2, err := yaml.JSONToYAML(j)
		// if err != nil {
		// 	fmt.Printf("err: %v\n", err)
		// 	return
		// }
		fmt.Println(string(y))
		stri := sshabu.GetAvaliableOptions()
		fmt.Println(stri)
	},
}

func init() {
	addCmd.AddCommand(hostCmd)
	addCmd.Flags().StringVarP(&DefOptions, "options", "o", "", "Host option definition")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hostCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hostCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
