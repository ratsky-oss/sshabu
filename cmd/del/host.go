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
		// err = shabu.Boil()
		cobra.CheckErr(err)
		// fmt.Println(shabu)
		fmt.Println(shabu.Hosts)
		err = shabu.FuncSshabuSlice(func(slicePtr interface{}) error {
			// 1. Проверяем, что нам передали именно *[]Host
			// ptr, ok := slicePtr.(**[]sshabu.Host) // Обратите внимание на двойной указатель
			// if !ok {
			// 	return fmt.Errorf("expected *[]Host, got %T", slicePtr)
			// }
			
			// // 2. Разыменовываем указатель на слайс
			// hosts := *ptr
			
			// // 3. Фильтруем элементы
			// filtered := make([]sshabu.Host, 0, len(*hosts))
			// for _, h := range *hosts {
			// 	if h.Name != args[0] {
			// 		filtered = append(filtered, h)
			// 	}
			// }
			
			// // 4. Меняем исходный слайс через указатель
			// *ptr = &filtered
			ptr, ok := slicePtr.(*[]sshabu.Host)
			if !ok {
				return fmt.Errorf("invalid type %T", slicePtr)
			}
			
			// Теперь работаем с *[]Host
			filtered := make([]sshabu.Host, 0)
			for _, h := range *ptr {
				if h.Name != args[0] {
					filtered = append(filtered, h)
				}
			}
			*ptr = filtered
			fmt.Println(ptr)
			return nil
			}, args[0])
		// err = shabu.DelHost("milanr-pi-new-01")
		cobra.CheckErr(err)
		y, err := yaml.Marshal(shabu)
		if err != nil {
			fmt.Printf("err: %v\n", err) 
			return
		}
		fmt.Println(string(y))
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
