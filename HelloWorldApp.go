package main

import (
	"fmt"
	"os"
)

func getArgument() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	} else {
		return "no args"
	}
}
func main() {
	fmt.Println(getArgument())

}
