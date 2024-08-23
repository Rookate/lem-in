package lemin

import "fmt"

func MoveAnts(pathfinder *PathFinder, data *LeminData, maxTurns int) {
	occupiedRoom := make(map[*Room]bool)
	occupiedtunnels := make(map[*Connection]bool)
	var nextRoom *Room
	hasArrived := 0
	instCount := 0
	turnCount := 0
	antsOnPath := make([]int, len(pathfinder.AllPaths))

	for turnCount < maxTurns {
		fmt.Println("New Turn")
		var moves []string

		for i := 0; i < len(data.AntList); i++ {
			ant := &data.AntList[i]

			if ant.OccupyingRoom == &data.EndRoom {
				continue
			}

			moveMade := false
			var path []*Room
			var pathIndex int

			if data.StartRoom.AntNb == 1 {
				pathIndex = FindBestPath(pathfinder, data, antsOnPath)
				path = pathfinder.AllPaths[pathIndex]
			} else {
				pathIndex = i % len(pathfinder.AllPaths)
				path = pathfinder.AllPaths[pathIndex]
			}

			fmt.Printf("Ant %s in %s considering path %d\n", ant.Name, ant.OccupyingRoom.Name, pathIndex)
			currentRoom := ant.OccupyingRoom

			for j := 0; j < len(path)-1; j++ {
				if path[j] == currentRoom {
					nextRoom = path[j+1]

					fmt.Printf("Ant %s considering moving to %s\n", ant.Name, nextRoom.Name)
					// Vérifier si la salle suivante est la salle de fin et si elle est disponible
					if nextRoom == &data.EndRoom {
						continue
					}

					tunnel := data.GetTunnel(currentRoom, nextRoom)
					if tunnel == nil {
						fmt.Printf("Tunnel between %s and %s not found!\n", currentRoom.Name, nextRoom.Name)
						continue
					}

					if !occupiedRoom[nextRoom] && !occupiedtunnels[tunnel] {
						fmt.Printf("Ant %s moving from %s to %s\n", ant.Name, currentRoom.Name, nextRoom.Name)
						if currentRoom == &data.StartRoom {
							data.StartRoom.AntNb--
						}
						ant.OccupyingRoom.Occupied = false
						ant.OccupyingRoom = nextRoom

						if ant.OccupyingRoom == &data.EndRoom {
							hasArrived++
							fmt.Printf("Ant %s has arrived at the end room!\n", ant.Name)
						} else {
							ant.OccupyingRoom.Occupied = true
							occupiedRoom[nextRoom] = true
							occupiedtunnels[tunnel] = true
						}

						moves = append(moves, fmt.Sprintf("%s-%s", ant.Name, nextRoom.Name))
						moveMade = true
						instCount++
						break
					} else {
						tunnel := data.GetTunnel(currentRoom, nextRoom)
						fmt.Printf("Room %s or tunnel %s->%s is occupied! Room occupied: %v, Tunnel occupied: %v\n",
							nextRoom.Name, currentRoom.Name, nextRoom.Name,
							occupiedRoom[nextRoom], occupiedtunnels[tunnel])
					}
				}
			}
			if moveMade {
				// Incrémenter le nombre de fourmis sur ce chemin
				antsOnPath[pathIndex]++
			} else {
				fmt.Printf("Ant %s could not move\n", ant.Name)
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

		// Réinitialiser la map occupiedRoom à la fin de chaque tour
		occupiedRoom = map[*Room]bool{}
		occupiedtunnels = map[*Connection]bool{}
	}
	fmt.Printf("Number of instructions: %d\nNumber of Turns: %d\n", instCount, turnCount)
}
