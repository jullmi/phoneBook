package main

import (
	"fmt"
	"os"
	"path"
	"phoneBook/internal"
)


func main() {
	data := internal.PopulateData(100)

	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("Usage: %s search|list|populate <arguments>\n", exe)
		return
	}

	switch arguments[1] {
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
		}
		if len(arguments) < 3 {
			fmt.Println("Enter search value and try it again")
		}
		res := internal.Search(data, arguments[2])
		if res != nil {
			fmt.Println(*res)
		} else {
			fmt.Println("Subscriber not found:", arguments[2])
		}

	case "list":
		internal.List(data)
	default:
		fmt.Println("Not a valid option")
	}
}
