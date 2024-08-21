package main

import (
	"fmt"
	"lemin"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Print("Usage with 'go run':\ngo run . LEMIN_FILE\nUsage with built object:\n./OBJECT_NAME LEMIN_FILE\n\n")
		return
	}

	leminData, errParse := lemin.ParseLeminFile(os.Args[1])
	if errParse != nil {
		fmt.Fprintf(os.Stderr, "ERROR - couldn't parse %s:\n%s\n", os.Args[1], errParse.Error())
		os.Exit(1)
	}

	fmt.Print(leminData.FileContent, "\n\n")

	err := lemin.Resolve(leminData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR finding solutions :\n%v\n", err)
		os.Exit(1)
	}
}
