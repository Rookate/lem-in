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

	{
		dataOk, errMsg := leminData.IsValidData()
		if !dataOk {
			fmt.Fprintf(os.Stderr, "ERROR - invalid data structure:\n%s\n", errMsg)
			os.Exit(1)
		}
	}

	fmt.Println(leminData.FileContent + "\n\nParsed:")
	fmt.Printf("Amount of ants: %d\nStart room: %v\nEnd room: %v\nCheckpoint rooms:\n%v\nPaths:\n%v\n", leminData.AntAmount, leminData.StartRoom, leminData.EndRoom, leminData.OtherRooms, leminData.Paths)
}
