package lemin

import "fmt"

func CreateAnts(antCount *LeminData, startRoom *Room) []Ant {

	var ants []Ant

	for i := 0; i < int(antCount.AntAmount); i++ {
		ant := Ant{
			Name:          fmt.Sprintf("Ant %d", i+1),
			OccupyingRoom: startRoom,
		}
		ants = append(ants, ant)
	}
	return ants
}

func MoveAnts(pathfinder *PathFinder, anCount int, ants []*Ant) {

}
