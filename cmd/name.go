package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

// nameCmd represents the name command
var nameCmd = &cobra.Command{
	Use:   "name",
	Short: "",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("name called")
	},
}

func hexToName(args []string) {
	var hm map[string]string

	// read json file
	data, err := ioutil.ReadFile("colornames.min.json")

	if err != nil {
		fmt.Printf("Error while reading the file %v", err)
	}

	_ = json.Unmarshal(data, &hm)

	name, ok := hm[args[0]]
	if ok {
		fmt.Printf("Name: %s, Hex: %s\n", name, args[0])
	} else {
		fmt.Println("Name Not Found.")
	}
}

func init() {
	rootCmd.AddCommand(nameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
