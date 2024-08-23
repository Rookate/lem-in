package lemin

import (
	"fmt"
)

func MoveAnts(pathfinder *PathFinder, data *LeminData) {
	instCount := 0
	turnCount := 0

	for {
		for i := range data.AntList {
			if *data.AntList[i].OccupyingRoom == data.EndRoom {
				continue
			}

			//fmt.Printf("Ant \"%s\" in room \"%s\":\n", data.AntList[i].Name, data.AntList[i].OccupyingRoom.Name)
			nextMove := data.NextBestMove(pathfinder, data.AntList[i].OccupyingRoom)

			if nextMove == data.AntList[i].OccupyingRoom {
				//fmt.Println("It cannot proceed. It waits its turn.")
				continue
			}

			//fmt.Printf("Best move is \"%s\" -> \"%s\".\n", data.AntList[i].OccupyingRoom.Name, nextMove.Name)
			if data.GetRoomIndexFromName(data.AntList[i].OccupyingRoom.Name) >= 0 {
				data.OtherRooms[data.GetRoomIndexFromName(data.AntList[i].OccupyingRoom.Name)].Occupied = false
			}
			data.AntList[i].OccupyingRoom = nextMove
			//fmt.Printf("Ant \"%s\" moved to \"%s\".\n", data.AntList[i].Name, data.AntList[i].OccupyingRoom.Name)
			if data.GetRoomIndexFromName(data.AntList[i].OccupyingRoom.Name) >= 0 {
				data.OtherRooms[data.GetRoomIndexFromName(data.AntList[i].OccupyingRoom.Name)].Occupied = true
			} else if *data.AntList[i].OccupyingRoom == data.EndRoom {
				data.EndRoom.AntNb++
				//fmt.Printf("And it has arrived to its destination. It's #%d.\n", data.EndRoom.AntNb)
			}

			fmt.Print(data.AntList[i].Name + "-" + data.AntList[i].OccupyingRoom.Name + " ")
			instCount++
		}

		fmt.Println()
		turnCount++

		if int(data.EndRoom.AntNb) == len(data.AntList) {
			break
		}
	}

	fmt.Printf("Nombre d'instructions: %d\nNombre de Tours: %d\n", instCount, turnCount)
}
