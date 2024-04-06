/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func deleteEntry(key string) error {
	i, ok := Index[key]
	if !ok {
		return fmt.Errorf("%s cannot be found", key)
	}
	DATA = append(DATA[:i], DATA[i+1:]...)
	delete(Index, key)

	err := SaveJSONFile(PATH)
	if err != nil {
		return err
	}
	fmt.Println("Success: entry was deleted")
	return nil
}

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete an entry",
	Long:  `delete an entry the phone book application`,
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetString("key")
		if key == "" {
			fmt.Println("Not a valid key:", key)
		}
		err := deleteEntry(key)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().String("key", "", "key to delete")
}
