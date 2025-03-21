package main

import (
	"fmt"
	"lemin/lemin"
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

	// fmt.Println(leminData.FileContent + "\n\nParsed:")

	// for _, rooms := range leminData.Paths {
	// 	fmt.Printf("Path from %s to %s : distance: %.4f\n", rooms.From.Name, rooms.To.Name, rooms.Distance)
	// }

	// for _, otherRooms := range leminData.OtherRooms {
	// 	fmt.Printf("Room: %s (%d;%d)\n",
	// 		otherRooms.Name,
	// 		otherRooms.X,
	// 		otherRooms.Y,
	// 	)
	// }

	// fmt.Printf("Amount of ants: %d\nStart room: %v\nEnd room: %v\nCheckpoint rooms:\n%v\n",
	// 	leminData.AntAmount,
	// 	leminData.StartRoom.Name,
	// 	leminData.EndRoom.Name,
	// 	leminData.OtherRooms,
	// )

	visited := make(map[string]bool)
	var path []*lemin.Room
	pathfinder := lemin.PathFinder{}

	lemin.DFS(leminData, &leminData.StartRoom, &leminData.EndRoom, visited, path, &pathfinder)
	initialAntCount := len(leminData.AntList)
	initialPathfinder := lemin.PathFinder{
		AllPaths:         pathfinder.AllPaths,
		AllDistancePaths: pathfinder.AllDistancePaths,
	}

	lemin.SortbyDistance(&pathfinder)
	movesByDistance, turnCountByDistance := lemin.MoveAnts(&pathfinder, leminData)

	leminData.StartRoom.AntNb = uint(len(leminData.AntList))
	leminData.EndRoom.AntNb = 0
	leminData.AntList = []lemin.Ant{}
	leminData.CreateAnts(initialAntCount)
	pathfinder.AllPaths = initialPathfinder.AllPaths
	pathfinder.AllDistancePaths = initialPathfinder.AllDistancePaths

	// Tri par nombre de salles (path)
	lemin.SortbyPaths(&pathfinder)
	movesByPaths, turnCountByPaths := lemin.MoveAnts(&pathfinder, leminData)

	// Tri par distance (distance)

	if turnCountByDistance < turnCountByPaths {
		//fmt.Println("Le tri par distance a pris moins de tours.")
		for _, move := range movesByDistance {
			fmt.Printf("%s\n", move)
		}

		fmt.Printf("Number of turns: %d\n", turnCountByDistance)
	} else {
		//fmt.Println("Le tri par salle a pris moins de tours.")
		for _, move := range movesByPaths {
			fmt.Printf("%s\n", move)
		}

		fmt.Printf("Number of turns: %d\n", turnCountByPaths)
	}

	// distance, turncountbyDistance := lemin.MoveAnts(&pathfinder, leminData)

	// if turnCountbyPath < turncountbyDistance {
	// 	for _, step := range paths {
	// 		fmt.Printf("%s ", step)
	// 	}
	// 	fmt.Printf("Number of Turns: %d\n", turnCountbyPath)
	// } else {
	// 	for _, step := range distance {
	// 		fmt.Printf("%s ", step)
	// 	}
	// 	fmt.Printf("Number of Turns: %d\n", turncountbyDistance)
	// }

	// for _, ant := range ants {
	// 	fmt.Printf("Ant name: %s\n", ant.Name)
	// }

	// count := 1
	// for _, distance := range pathfinder.AllDistancePaths {
	// 	fmt.Println(distance)
	// 	count++
	// }

	// fmt.Printf("Tous les chemins trouvés : %d\n", count-1)
	// count = 1
	// for _, p := range pathfinder.AllPaths {
	// 	fmt.Printf("Chemin %d: ", count)
	// 	for _, r := range p {
	// 		fmt.Printf("%s -> ", r.Name)
	// 	}
	// 	fmt.Println("Fin")
	// 	count++
	// }
}
