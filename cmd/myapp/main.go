package main

import (
	"fmt"
	"os"
	"path"
	"phoneBook/internal"
	"strings"
)

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("Usage: %s search|list|insert|delete <arguments>\n", exe)
		return
	}

	_, err := os.Stat(internal.PATH)
	if err != nil {
		fmt.Println("CREATING CSV FILE...")
		file, err := os.Create(internal.PATH)
		if err != nil {
			fmt.Println(err)
		}
		file.Close()
	}

	fileInfo, err := os.Stat(internal.PATH)

	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		fmt.Println(internal.PATH, "is not a regular file")
		return
	}

	err = internal.ReadCSVFile(internal.PATH)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = internal.CreateIndex()
	if err != nil {
		fmt.Println("Cannot create index")
		return
	}

	switch arguments[1] {
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Phone number")
		}

		t := strings.ReplaceAll(arguments[2], "-", "")
		if !internal.MatchTel(t) {
			fmt.Println("Not a valid phone number:", t)
		}

		res := internal.Search(t)
		if res != nil {
			fmt.Println(*res)
		} else {
			fmt.Println("Subscriber not found:", arguments[2])
		}

	case "list":
		internal.List()
	case "insert":
		if len(arguments) != 5 {
			fmt.Println("Usage: insert Name Surname Phone")
			return
		}
		t := strings.ReplaceAll(arguments[4], "-", "")
		if !internal.MatchTel(t) {
			fmt.Println("Phone number is not valid")
			return
		}
		sub := internal.NewSubscriber(arguments[2], arguments[3], t)
		if sub != nil {
			err = internal.Insert(sub)
			if err != nil {
				fmt.Println(err)
			}
		}
	case "delete":
		if len(arguments) != 3 {
			fmt.Println("Usage: delete Phone number")
		}

		t := strings.ReplaceAll(arguments[2], "-", "")
		if !internal.MatchTel(t) {
			fmt.Println("Phone number is not valid: ", t)
		}
		err = internal.DeleteSubscriber(t)
		if err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("Not a valid option")
	}
}
