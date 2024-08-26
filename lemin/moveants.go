package lemin

import (
	"fmt"
	"strings"
)

func MoveAnts(pathfinder *PathFinder, data *LeminData) ([]string, int) {
	turnCount := 0
	occupiedTunnels := make(map[string]int)
	initialThreshold := int(0.1 * float64(len(data.AntList)))
	var moves []string

	robinetMode := data.isDirectlyConnected(data.StartRoom, data.EndRoom)

	for {
		// Initialiser la liste des mouvements pour ce tour
		var currentTurnMoves []string

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
				// Choisir le meilleur mouvement basé sur les critères
				nextMove = data.NextBestMove(pathfinder, ant.OccupyingRoom)
			}

			if nextMove == ant.OccupyingRoom || nextMove == nil {
				continue
			}

			tunnelKey := fmt.Sprintf("%s-%s", ant.OccupyingRoom.Name, nextMove.Name)

			if occupiedTunnels[tunnelKey] > 0 {
				continue
			}

			occupiedTunnels[tunnelKey] = 1

			if ant.OccupyingRoom == &data.StartRoom {
				data.StartRoom.AntNb--
			}

			if data.GetRoomIndexFromName(ant.OccupyingRoom.Name) >= 0 {
				data.OtherRooms[data.GetRoomIndexFromName(ant.OccupyingRoom.Name)].Occupied = false
			}
			ant.OccupyingRoom = nextMove
			if data.GetRoomIndexFromName(ant.OccupyingRoom.Name) >= 0 {
				data.OtherRooms[data.GetRoomIndexFromName(ant.OccupyingRoom.Name)].Occupied = true
			} else if *ant.OccupyingRoom == data.EndRoom {
				data.EndRoom.AntNb++
			}

			// Ajouter le mouvement à la liste pour ce tour
			currentTurnMoves = append(currentTurnMoves, fmt.Sprintf("%s-%s", ant.Name, ant.OccupyingRoom.Name))
		}

		// Ajouter les mouvements du tour actuel à la liste globale
		if len(currentTurnMoves) > 0 {
			moves = append(moves, strings.Join(currentTurnMoves, " "))
		}

		turnCount++

		// Réinitialiser les tunnels occupés pour le prochain tour
		for key := range occupiedTunnels {
			occupiedTunnels[key] = 0
		}

		if int(data.EndRoom.AntNb) == len(data.AntList) {
			break
		}
	}

	return moves, turnCount
}

func (data *LeminData) NextBestMove(pathfinder *PathFinder, currentRoom *Room) *Room {
	if currentRoom == nil || pathfinder == nil {
		return nil
	}

	for _, path := range pathfinder.AllPaths {
		for i := range path {

			if *path[i] == *currentRoom {
				if data.GetRoomIndexFromName(path[i+1].Name) == -2 {
					return &data.EndRoom
				}

				if !data.OtherRooms[data.GetRoomIndexFromName(path[i+1].Name)].Occupied {
					return path[i+1]
				}
			}
		}
	}

	return currentRoom
}

// Fonction pour vérifier si startRoom est directement reliée à endRoom
func (data *LeminData) isDirectlyConnected(r1 Room, r2 Room) bool {
	for _, path := range data.TunnelList {
		if *path.From == r1 && *path.To == r2 || *path.From == r2 && *path.To == r1 {
			return true
		}
	}
	return false
}

func getNextRoomOnPath(path []*Room, currentRoom *Room) *Room {
	for i := 0; i < len(path)-1; i++ {
		if path[i] == currentRoom {
			return path[i+1]
		}
	}
	return nil
}
