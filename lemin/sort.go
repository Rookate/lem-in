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

func SortPaths(pathFinder *PathFinder) {
	indices := make([]int, len(pathFinder.AllDistancePaths))
	for i := range indices {
		indices[i] = i
	}

	sort.Slice(indices, func(i, j int) bool {
		return pathFinder.AllDistancePaths[indices[i]] < pathFinder.AllDistancePaths[indices[j]]
	})

	sortedPaths := make([][]*Room, len(pathFinder.AllPaths))
	sortedDistances := make([]float64, len(pathFinder.AllDistancePaths))

	for i, index := range indices {
		sortedPaths[i] = pathFinder.AllPaths[index]
		sortedDistances[i] = pathFinder.AllDistancePaths[index]
	}
	pathFinder.AllPaths = sortedPaths
	pathFinder.AllDistancePaths = sortedDistances
}
