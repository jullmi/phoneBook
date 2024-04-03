package internal

import (
	"fmt"
	"os"
)

func SetCSVFile() error {
	filepath := os.Getenv("PHONEBOOK")
	if filepath != "" {
		PATH = filepath
	}
	_, err := os.Stat(PATH)
	if err != nil {
		fmt.Println("CREATING CSV FILE...")
		file, err := os.Create(PATH)
		if err != nil {
			file.Close()
			return err
		}
		file.Close()
	}

	fileInfo, err := os.Stat(PATH)

	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		return fmt.Errorf(PATH, "is not a regular file")
	}

	return nil
}
