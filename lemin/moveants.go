package lemin

import "fmt"

func MoveAntss(pathfinder *PathFinder, data *LeminData, ants []Ant) {
	occupiedRoom := make(map[*Room]bool)
	endRoomCooldown := make(map[int]bool)
	antsOnPath := make([]int, len(pathfinder.AllPaths))
	var nextRoom *Room
	count := 0
	turnCount := 0

	for {
		allAntsAtEnd := true
		var moves []string

		for i := 0; i < len(data.AntList); i++ {
			ant := &ants[i]

			if ant.OccupyingRoom == &data.EndRoom {
				continue
			}

			moveMade := false
			currentRoom := ant.OccupyingRoom

			// Trouver le meilleur chemin pour cette fourmi
			bestPathIndex := FindBestPath(pathfinder, data, antsOnPath)
			bestPath := pathfinder.AllPaths[bestPathIndex]

			for j := 0; j < len(bestPath)-1; j++ {
				if bestPath[j] == currentRoom {
					nextRoom = bestPath[j+1]

					// Vérifier si la salle suivante est la EndRoom et si elle est en cooldown
					if nextRoom == &data.EndRoom && endRoomCooldown[bestPathIndex] {
						continue
					}

					if !occupiedRoom[nextRoom] {
						ant.OccupyingRoom.Occupied = false
						ant.OccupyingRoom = nextRoom

						// Définir le cooldown si la salle suivante est la EndRoom
						if ant.OccupyingRoom == &data.EndRoom {
							endRoomCooldown[bestPathIndex] = true
						} else {
							ant.OccupyingRoom.Occupied = true
							occupiedRoom[nextRoom] = true
						}

						moves = append(moves, fmt.Sprintf("%s-%s", ant.Name, nextRoom.Name))
						moveMade = true
						count++
						break
					}
				}
			}

			// Si une fourmi n'a pas bougé, elle n'est pas encore arrivée à la salle finale
			if !moveMade {
				allAntsAtEnd = false
			}
		}

		// Décrémenter le cooldown pour chaque chemin
		for pathIndex := range endRoomCooldown {
			if endRoomCooldown[pathIndex] {
				endRoomCooldown[pathIndex] = false
			}
		}

		turnCount++

		for _, move := range moves {
			fmt.Printf("%s ", move)
		}
		fmt.Println()

		if allAntsAtEnd {
			break
		}

		// Réinitialiser la map occupiedRoom à la fin de chaque tour
		occupiedRoom = make(map[*Room]bool)
	}
	fmt.Printf("Number of instructions: %d\nNumber of Turns: %d\n", count, turnCount)
}

func FindBestPath(pathfinder *PathFinder, data *LeminData, antsOnPath []int) int {
	bestPathIndex := 0
	bestPathScore := float64(len(pathfinder.AllPaths[0])) + float64(antsOnPath[0])

	for i := 1; i < len(pathfinder.AllPaths); i++ {
		currentPathScore := float64(len(pathfinder.AllPaths[i])) + float64(antsOnPath[i])
		if currentPathScore < bestPathScore {
			bestPathScore = currentPathScore
			bestPathIndex = i
		}
	}

	return bestPathIndex
}

func MoveAntsss(pathfinder *PathFinder, data *LeminData) {
	occupiedRoom := make(map[*Room]bool)
	endRoomCooldown := make(map[int]int) // Map to track the cooldown for each path
	var nextRoom *Room
	hasArrived := 0
	instCount := 0
	turnCount := 0

	for {
		var moves []string

		for i := 0; i < len(data.AntList); i++ {
			ant := &data.AntList[i]

			if ant.OccupyingRoom == &data.EndRoom {
				continue
			}

			moveMade := false

			for pathIndex, path := range pathfinder.AllPaths {
				currentRoom := ant.OccupyingRoom

				for j := 0; j < len(path)-1; j++ {
					if path[j] == currentRoom {
						nextRoom = path[j+1]

						// Check if the nextRoom is the endRoom and if it's available
						if nextRoom == &data.EndRoom && endRoomCooldown[pathIndex] > 0 {
							continue
						}

						if !occupiedRoom[nextRoom] {
							ant.OccupyingRoom.Occupied = false
							ant.OccupyingRoom = nextRoom

							if ant.OccupyingRoom == &data.EndRoom {
								hasArrived++
								// Set cooldown for the endRoom for this path
								endRoomCooldown[pathIndex] = 1
							} else {
								ant.OccupyingRoom.Occupied = true
								occupiedRoom[nextRoom] = true
							}

							moves = append(moves, fmt.Sprintf("%s-%s", ant.Name, nextRoom.Name))
							moveMade = true
							instCount++
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
			}
		}

		// Decrement the cooldown for the endRoom for each path
		for pathIndex := range endRoomCooldown {
			if endRoomCooldown[pathIndex] > 0 {
				endRoomCooldown[pathIndex]--
			}
		}

		turnCount++

		for _, move := range moves {
			fmt.Printf("%s ", move)
		}
		fmt.Println()

		if hasArrived == len(data.AntList) {
			break
		}

		// Reset the occupiedRoom map at the end of each turn
		occupiedRoom = make(map[*Room]bool)
	}
	fmt.Printf("Number of instructions: %d\nNumber of Turns: %d\n", instCount, turnCount)
}

/*

	- Find Best path possible
	- we need to add the number of ant in the actual path and the number of rooms. For exemple if we have 3 ants on the path and we have 5 rooms before we get the endroom.
	The ant will take 8 instructions to get endroom
	- Then we have to compare path to get the best path.


	First idea : Get the Best path for each ant then move them.
	 - Get Start Room then decrement the number of ant when then leave start Room
	 - Savoir si la fourmi quitte la room de départ
*/
