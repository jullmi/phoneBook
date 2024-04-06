/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

func matchTel(tel string) bool {
	t := []byte(tel)
	re := regexp.MustCompile(`\d+$`)
	return re.Match(t)
}

func search(key string) *Entry {
	i, ok := Index[key]

	if !ok {
		return nil
	}
	return &DATA[i]
}

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search of a phone number",
	Long:  `The command which search a pnone humber into the phone book`,
	Run: func(cmd *cobra.Command, args []string) {
		searchKey, _ := cmd.Flags().GetString("key")
		if searchKey == "" {
			fmt.Println("Not a valid a search key: ", searchKey)
		}

		t := strings.ReplaceAll(searchKey, "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid phone number: ", t)
			return
		}

		temp := search(t)
		if temp == nil {
			fmt.Println("Phone number not found:", t)
			return
		}

		fmt.Println(*temp)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringP("key", "k", "", "key to search")
}
