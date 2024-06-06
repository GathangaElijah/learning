package statistics

import (
	"fmt"
	"os"
)

// This functions fetches the string at the commandline
func FileFetcher() string {
	var args = os.Args
	if len(args) == 1 || len(args) > 2 {
		fmt.Println(`Usage: "go run main.go data.txt"`)
		os.Exit(1)
	}
	fileName := args[1]
	return fileName
}
