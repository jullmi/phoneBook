package internal

import (
	"encoding/csv"
	"os"
)

var PATH = "/Users/jullmi/Documents/development/go/learn_golang/phoneBook/data.csv"

func ReadCSVFile(filePath string) error {

	_, err := os.Stat(filePath)
	if err != nil {
		return err
	}
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return err
	}

	for _, line := range lines {
		temp := &Subscriber{line[0], line[1], line[2], line[3]}
		Data = append(Data, *temp)
	}

	return nil

}
