package internal

import (
	"encoding/csv"
	"os"
)

func SaveCSVfile(filepath string) error {

	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)

	for _, row := range DATA {
		temp := []string{row.Name, row.Surname, row.Tel, row.LastAccess}
		_ = csvwriter.Write(temp)

	}
	csvwriter.Flush()

	return nil

}
