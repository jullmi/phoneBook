/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

func (entity PhoneBook) Len() int {
	return len(entity)
}

func (entity PhoneBook) Less(i, j int) bool {
	if entity[i].Surname == entity[j].Surname {
		return entity[i].Name < entity[i].Name
	}
	return entity[i].Surname < entity[j].Surname
}

func (entity PhoneBook) Swap(i, j int) {
	entity[i], entity[j] = entity[j], entity[i]
}

// PrettyPrintJSONstream pretty prints the contents of the phone book
func prettyPrintJSONstream(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "\t")

	err := encoder.Encode(data)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func list() {
	sort.Sort(DATA)
	text, err := prettyPrintJSONstream(DATA)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(text)
	fmt.Printf("%d records in total.\n", len(DATA))
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all entries",
	Long:  `This command lists all entries in the phone book application`,
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
