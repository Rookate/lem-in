package lemin

import "sort"

func CalculatePathDistance(path []*Room, graph *Graph) float64 {
	totalDistance := 0.0

	for i := 0; i < len(path)-1; i++ {
		from := path[i]
		to := path[i+1]

		// Cherche la distance dans les connexions de la salle actuelle
		for _, conn := range graph.GetNeighbors(from) {
			if conn.To == to {
				totalDistance += conn.Distance
				break
			}
		}
	}

	return totalDistance
}

// fonction qui tri toutes les distances de la plus courte à la plus longue
func SortPaths(pathFinder *PathFinder) {
	sort.Slice(pathFinder.AllDistancePaths, func(i, j int) bool {
		return pathFinder.AllDistancePaths[i] < pathFinder.AllDistancePaths[j]
	})
}
