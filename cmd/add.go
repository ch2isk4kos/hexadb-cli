package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

// add hexadecimal command
func add(args []string) {
	hex := args[0]
	color := args[1]

	// read json
	data, err := ioutil.ReadFile("colornames.min.js")
	if err != nil {
		fmt.Printf("Error Writing File: %v", err)
	}

	// map hex to to color
	var hexMap map[string] string
	
	// unmarshall the json file
	_ = json.Unmarshal(data, &hexMap)

	// check if hex exists
	hexName, ok := hexMap[hex]
	if ok {
		fmt.Printf("\n\nHex Already Rxists: %s\n", hexName)
	} else {
		hexMap[hex] = color

		hexJSON, _ := json.Marshal(hexMap)

		err = ioutil.WriteFile("colornames.min.json", hexJSON, 0777)
		if err != nil {
			fmt.Printf("\nError Writing File: %v\n\n", err)
		}

		fmt.Printf("\nHex Added: %v\n\n", color)
	}
}
