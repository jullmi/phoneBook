/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


func insert(pS *Entry) error {
	_, ok := Index[(*pS).Tel]
	if ok {
		return fmt.Errorf("%s already exists", pS.Tel)
	}
	DATA = append(DATA, *pS)

	_ = CreateIndex()

	err := SaveJSONFile(PATH)
	if err != nil {
		return err
	}
	fmt.Println("Success: entry was inserted")
	return nil
}

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "insert new data",
	Long: `The command inserts new data into the phone book application`,
	Run: func(cmd *cobra.Command, args []string) {

		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			fmt.Println("Not a valid name: ", name)
			return
		}

		surname, _ := cmd.Flags().GetString("surname")
		if surname == "" {
			fmt.Println("Not a valid surname: ", surname)
		}

		tel, _ :=cmd.Flags().GetString("tel")
		if tel == "" {
			fmt.Println("Not a valid phone number: ", tel)
		}
		temp := NewEntry(name, surname, tel)

		err := insert(temp)
		if err != nil {
			fmt.Println(err)
			return 
		}
	},
}

func init() {
	rootCmd.AddCommand(insertCmd)
	insertCmd.Flags().StringP("name", "n", "", "name value")
	insertCmd.Flags().StringP("surname", "s", "", "surname value")
	insertCmd.Flags().StringP("tel", "t", "", "phone value")
}
