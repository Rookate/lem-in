package lemin

import "fmt"

func CreateAnts(antCount *LeminData, startRoom *Room) []Ant {

	var ants []Ant

	for i := 0; i < int(antCount.AntAmount); i++ {
		ant := Ant{
			Name:          fmt.Sprintf("L%d", i+1),
			OccupyingRoom: startRoom,
		}
		ants = append(ants, ant)
	}
	return ants
}

func MoveAnts(pathfinder *PathFinder, antCount *LeminData, ants []Ant) {
	occupiedRoom := make(map[*Room]bool)
	var nextRoom *Room
	count := 0
	counTurn := 0

	for {

		allAntsAtEnd := true
		var moves []string
		for i := 0; i < int(antCount.AntAmount); i++ {
			ant := &ants[i]

			if ant.OccupyingRoom == &antCount.EndRoom {
				continue
			}

			moveMade := false

			for _, path := range pathfinder.AllPaths {
				currentRoom := ant.OccupyingRoom

				for j := 0; j < len(path)-1; j++ {
					if path[j] == currentRoom {
						nextRoom = path[j+1]

						if !occupiedRoom[nextRoom] {
							ant.OccupyingRoom.Occupied = false
							ant.OccupyingRoom = nextRoom
							if ant.OccupyingRoom != &antCount.EndRoom {
								ant.OccupyingRoom.Occupied = true
							}
							occupiedRoom[nextRoom] = true

							moves = append(moves, fmt.Sprintf("%s-%s", ant.Name, nextRoom.Name))
							moveMade = true
							count++

							break
						}
					}
					if moveMade {
						break
					}
				}
				if moveMade {
					break
				}

				if ant.OccupyingRoom != &antCount.EndRoom {
					allAntsAtEnd = false
				}
			}
		}
		counTurn++

		for _, move := range moves {
			fmt.Printf("%s ", move)
		}
		fmt.Println()

		if allAntsAtEnd {
			break
		}

		occupiedRoom = make(map[*Room]bool)
	}
	fmt.Printf("Number of instruction: %d\nNumber of Turn: %d\n", count, counTurn)
}
