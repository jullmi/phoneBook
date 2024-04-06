/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

type Entry struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Tel        string `json:"tel"`
	LastAccess string `json:"lastccess"`
}

var PATH = "/Users/jullmi/Documents/development/go/learn_golang/phoneBook/data.json"

func NewEntry(name, surname, tel string) *Entry {
	if tel == "" || surname == "" {
		return nil
	}

	lastAccess := strconv.FormatInt(time.Now().Unix(), 10)
	return &Entry{Name: name, Surname: surname, Tel: tel, LastAccess: lastAccess}
}

type PhoneBook []Entry

var DATA = PhoneBook{}

// Serialize serializes a slice with JSON records
func Serialize(slice interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(slice)
}

// DeSerialize decodes a serialized slice with JSON records
func Deserialize(slice interface{}, r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(slice)
}

var Index map[string]int

func CreateIndex() error {
	Index = make(map[string]int)

	for i, k := range DATA {
		key := k.Tel
		Index[key] = i
	}
	return nil
}

func SaveJSONFile(filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer file.Close()

	err = Serialize(&DATA, file)
	if err != nil {
		return err
	}
	return nil
}

func readJSONFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	file, err := os.Open(filepath)
	if err != nil {
		return err
	}

	defer file.Close()

	err = Deserialize(&DATA, file)
	if err != nil {
		return err
	}

	return nil

}

func setJSONFile() error {
	filepath := os.Getenv("PHONEBOOK")
	if filepath != "" {
		PATH = filepath
	}

	_, err := os.Stat(PATH)
	if err != nil {
		fmt.Println("CREATING FILE...")
		file, err := os.Create(PATH)
		if err != nil {
			file.Close()
			return err
		}

		file.Close()
	}

	fileinfo, err := os.Stat(PATH)

	if err != nil {
		return err
	}

	mode := fileinfo.Mode()

	if !mode.IsRegular() {
		return fmt.Errorf("%s is not a regular file", PATH)
	}

	return nil

}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "phoneBook",
	Short: "The phone book application",
	Long:  `This is application Phone book which uses JSON records`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	err := setJSONFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = readJSONFile(PATH)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		return
	}

	CreateIndex()

	cobra.CheckErr(rootCmd.Execute())
}

func init() {
}
