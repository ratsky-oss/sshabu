/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package del

import (
	// "fmt"
	"fmt"
	sshabu "sshabu/pkg"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	// "gopkg.in/yaml.v3"
)

var DefOptions []string

func stringsToMap(input []string) map[string]string {
	result := make(map[string]string)

	for _, str := range input {
		// Split each string into key and value using ":" as a separator
		keyValue := strings.Split(str, ":")
		if len(keyValue) == 2 {
			// Trim any leading or trailing whitespaces from key and value
			key := strings.TrimSpace(keyValue[0])
			value := strings.TrimSpace(keyValue[1])
			result[key] = value
		}
	}
	return result
}

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
		err = shabu.Boil()
		cobra.CheckErr(err)
		// fmt.Println(shabu)
		err = shabu.FuncSshabuObj(func(i interface{}) error {
			fmt.Println(i)
			return nil
			}, args[0])
		// fmt.Println(host)
		// err = shabu.DelHost("milanr-pi-new-01")
		cobra.CheckErr(err)
		// y, err := yaml.Marshal(shabu)
		// if err != nil {
		// 	fmt.Printf("err: %v\n", err) 
		// 	return
		// }
		// fmt.Println(string(y))
	},
}

func init() {
	delCmd.AddCommand(hostCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hostCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hostCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
