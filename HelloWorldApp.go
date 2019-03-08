package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	errNoError int = 0 + iota
	errArgumentMissing
	errFileNotExist
	errFileNotReadable
)

func main() {
	fileName, err := getArgument()
	if err != nil {
		fmt.Println(err)
		os.Exit(errArgumentMissing)
	}

	fileExists := fileExists(fileName)
	if !fileExists {
		fmt.Printf("File '%s' does not exist\n", fileName)
		os.Exit(errFileNotExist)
	}

	fileIsReadable := fileIsReadable(fileName)
	if !fileIsReadable {
		fmt.Printf("File '%s' is not readable\n", fileName)
		os.Exit(errFileNotReadable)
	}

	encoded := encodeFile(fileName)
	fmt.Println(encoded)

	os.Exit(errNoError)

}

func encodeFile(filename string) string {
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}

	return base64.StdEncoding.EncodeToString(fileContent)
}

func getArgument() (string, error) {
	if len(os.Args) > 1 {
		return os.Args[1], nil
	}

	return "", errors.New("No argument was defined")
}

func fileExists(filename string) bool {
	// Stat returns file info. It will return
	// an error if there is no file.
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}

func fileIsReadable(filename string) bool {
	_, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil && os.IsPermission(err) {
		return false
	}

	return true
}
