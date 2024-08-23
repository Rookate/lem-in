package lemin

import "fmt"

func MoveAnts(pathfinder *PathFinder, data *LeminData, maxTurns int) {
	occupiedRoom := make(map[*Room]bool)
	occupiedTunnels := make(map[string]int) // Utilisation de chaînes de caractères comme clés pour les tunnels
	var nextRoom *Room
	hasArrived := 0
	instCount := 0
	turnCount := 0
	antsOnPath := make([]int, len(pathfinder.AllPaths))

	for turnCount < maxTurns {
		var moves []string

		for i := 0; i < len(data.AntList); i++ {
			ant := &data.AntList[i]

			if ant.OccupyingRoom == &data.EndRoom {
				continue
			}

			var path []*Room

			// Choix du chemin basé sur l'index de la fourmi
			pathIndex := i % len(pathfinder.AllPaths)
			path = pathfinder.AllPaths[pathIndex]
			currentRoom := ant.OccupyingRoom
			for j := 0; j < len(path)-1; j++ {
				if path[j] == currentRoom {
					nextRoom = path[j+1]

					tunnel := data.GetTunnel(currentRoom, nextRoom)
					if tunnel == nil {
						continue
					}

					// Création de la clé pour le tunnel
					tunnelKey := fmt.Sprintf("%s-%s", currentRoom.Name, nextRoom.Name)

					// Vérification de l'occupation
					if !occupiedRoom[nextRoom] && occupiedTunnels[tunnelKey] == 0 {
						// Mise à jour des occupations
						if currentRoom == &data.StartRoom {
							data.StartRoom.AntNb--
							// Incrémenter le compteur de fourmis sur ce chemin lorsqu'une fourmi quitte la salle de départ
							antsOnPath[pathIndex]++
						}

						ant.OccupyingRoom.Occupied = false
						ant.OccupyingRoom = nextRoom
						occupiedTunnels[tunnelKey] = 1

						if ant.OccupyingRoom == &data.EndRoom {
							hasArrived++
							// Décrémenter le compteur de fourmis sur ce chemin lorsqu'une fourmi atteint la salle de fin
							antsOnPath[pathIndex]--
						} else {
							ant.OccupyingRoom.Occupied = true
							occupiedRoom[nextRoom] = true
						}
						moves = append(moves, fmt.Sprintf("%s-%s", ant.Name, nextRoom.Name))
						instCount++
						break
					}
				}
			}
		}
		turnCount++

		for _, move := range moves {
			fmt.Printf("%s ", move)
		}
		fmt.Println()

		for key := range occupiedTunnels {
			occupiedTunnels[key] = 0
		}

		occupiedRoom = make(map[*Room]bool)

		if hasArrived == len(data.AntList) {
			break
		}
	}

	fmt.Printf("Nombre d'instructions: %d\nNombre de Tours: %d\n", instCount, turnCount)
}
