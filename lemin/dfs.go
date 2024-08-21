package lemin

func DFS(currentRoom *Room, endRoom *Room, visited map[string]bool, path []*Room, pathFinder *PathFinder, graph *Graph) {
	visited[currentRoom.Name] = true
	path = append(path, currentRoom)

	if currentRoom != endRoom {
		for _, neighbor := range graph.GetNeighbors(currentRoom) {
			if !visited[neighbor.To.Name] {
				DFS(neighbor.To, endRoom, visited, path, pathFinder, graph)
			}
		}
	} else {
		// Copy the path and append it to AllPaths
		copiedPath := append([]*Room{}, path...)
		pathFinder.AllPaths = append(pathFinder.AllPaths, copiedPath)

		totalDistance := CalculatePathDistance(copiedPath, graph)
		pathFinder.AllDistancePaths = append(pathFinder.AllDistancePaths, totalDistance)
	}

	visited[currentRoom.Name] = false
}
