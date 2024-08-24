package lemin

import (
	"fmt"
)

func MoveAnts(pathfinder *PathFinder, data *LeminData) {
	instCount := 0
	turnCount := 0
	occupiedTunnels := make(map[string]int)
	initialThreshold := int(0.3 * float64(len(data.AntList)))

	robinetMode := isDirectlyConnected(data.StartRoom, data.EndRoom, pathfinder.AllPaths)

	for {
		for i := range data.AntList {
			ant := &data.AntList[i]

			if *ant.OccupyingRoom == data.EndRoom {
				continue
			}

			var nextMove *Room

			if robinetMode && data.StartRoom.AntNb > uint(initialThreshold) {
				// Répartition cyclique des fourmis initiales
				pathIndex := i % len(pathfinder.AllPaths)
				path := pathfinder.AllPaths[pathIndex]
				// Trouver la salle suivante sur le chemin
				nextMove = getNextRoomOnPath(path, ant.OccupyingRoom)
			} else {
				//fmt.Printf("Ant \"%s\" in room \"%s\":\n", data.AntList[i].Name, data.AntList[i].OccupyingRoom.Name)
				nextMove = data.NextBestMove(pathfinder, ant.OccupyingRoom)
			}

			if nextMove == ant.OccupyingRoom || nextMove == nil {
				//fmt.Println("It cannot proceed. It waits its turn.")
				continue
			}

			tunnelKey := fmt.Sprintf("%s-%s", ant.OccupyingRoom.Name, nextMove.Name)

			if occupiedTunnels[tunnelKey] > 0 {
				continue
			}

			occupiedTunnels[tunnelKey] = 1

			//fmt.Printf("Best move is \"%s\" -> \"%s\".\n", data.AntList[i].OccupyingRoom.Name, nextMove.Name)
			if data.GetRoomIndexFromName(ant.OccupyingRoom.Name) >= 0 {
				data.OtherRooms[data.GetRoomIndexFromName(ant.OccupyingRoom.Name)].Occupied = false
			}
			data.AntList[i].OccupyingRoom = nextMove
			//fmt.Printf("Ant \"%s\" moved to \"%s\".\n", data.AntList[i].Name, data.AntList[i].OccupyingRoom.Name)
			if data.GetRoomIndexFromName(ant.OccupyingRoom.Name) >= 0 {
				data.OtherRooms[data.GetRoomIndexFromName(ant.OccupyingRoom.Name)].Occupied = true
			} else if *ant.OccupyingRoom == data.EndRoom {
				data.EndRoom.AntNb++
				//fmt.Printf("And it has arrived to its destination. It's #%d.\n", data.EndRoom.AntNb)
			}

			fmt.Print(ant.Name + "-" + ant.OccupyingRoom.Name + " ")
			instCount++
		}

		fmt.Println()
		turnCount++

		for key := range occupiedTunnels {
			occupiedTunnels[key] = 0
		}

		if int(data.EndRoom.AntNb) == len(data.AntList) {
			break
		}
	}

	fmt.Printf("Nombre d'instructions: %d\nNombre de Tours: %d\n", instCount, turnCount)
}
