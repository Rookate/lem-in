package lemin

import (
	"sort"
)

func (data *LeminData) CalculateDistanceRooms(path []*Room) float64 {
	totalDistance := 0.0

	for i := 0; i < len(path)-1; i++ {
		for _, tunnel := range data.TunnelList {
			if (*tunnel.From == *path[i] && *tunnel.To == *path[i+1]) || (*tunnel.To == *path[i] && *tunnel.From == *path[i+1]) {
				totalDistance += tunnel.Distance
				break
			}
		}
	}

	return totalDistance
}

func SortbyPaths(pathFinder *PathFinder) {
	indices := make([]int, len(pathFinder.AllPaths))
	for i := range indices {
		indices[i] = i
	}

	sort.Slice(indices, func(i, j int) bool {
		return len(pathFinder.AllPaths[indices[i]]) < len(pathFinder.AllPaths[indices[j]])
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

func SortbyDistance(pathFinder *PathFinder) {
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
