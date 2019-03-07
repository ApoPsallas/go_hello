package main

import (
	"fmt"
	"os"
)

var (
	fileInfo *os.FileInfo
	err      error
)

func getArgument() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	} else {
		return "no args"
	}
}
func fileExists(filename string) bool {
	// Stat returns file info. It will return
	// an error if there is no file.
	fileInfo, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	_ = fileInfo
	return true
}
func main() {
	fmt.Println(fileExists(getArgument()))

}
