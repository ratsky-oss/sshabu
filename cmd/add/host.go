/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package add

import (
	"fmt"
	sshabu "sshabu/pkg"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
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
		// if shabu.AreAllUnique(){
		// 	fmt.Println("YAML seems OK")
		// 	}  else {
		// 	fmt.Println("Error: 'Name' Fields must be unique")
		// 	os.Exit(1)
		// }
		// names := sshabu.FindNamesInShabu(shabu)
		
		// host_options, _ := cmd.Flags().GetString("options")
		err = shabu.Boil()
		cobra.CheckErr(err)
		// y2, err := yaml.JSONToYAML(j)
		// if err != nil {
			// 	fmt.Printf("err: %v\n", err)
			// 	return
			// }
		hostParams := stringsToMap(DefOptions)
		_ , err = sshabu.AreKeysInOption(hostParams)
		cobra.CheckErr(err)
		hostParams["name"] = args[0]
		host := sshabu.CreateHost(hostParams)
		shabu.AddHost(host.(sshabu.Host))
		y, err := yaml.Marshal(shabu)
		if err != nil {
			fmt.Printf("err: %v\n", err) // shabu add host srv-1 -o "Hostname: jkjkjk, "
			return
		}
		fmt.Println(string(y))
		fmt.Println(host)
	},
}

func init() {
	addCmd.AddCommand(hostCmd)
	hostCmd.Flags().StringSliceVar(&DefOptions, "option", []string{}, "Host option definition")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hostCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hostCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
