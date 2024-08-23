package lemin

func DFS(data *LeminData, currentRoom *Room, endRoom *Room, visited map[string]bool, path []*Room, pathFinder *PathFinder) {
	visited[currentRoom.Name] = true
	path = append(path, currentRoom)

	if currentRoom != endRoom {
		neighbors := data.GetNeighbors(currentRoom)
		for _, neighbor := range neighbors {
			if !visited[neighbor.Name] {
				DFS(data, neighbor, endRoom, visited, path, pathFinder)
			}
		}
	} else {
		// Copy the path
		copiedPath := make([]*Room, len(path))
		copy(copiedPath, path)
		if !Doublon(copiedPath, pathFinder.AllPaths) {
			pathFinder.AllPaths = append(pathFinder.AllPaths, copiedPath)
		}

		totalDistance := data.CalculateDistanceRooms(copiedPath)
		pathFinder.AllDistancePaths = append(pathFinder.AllDistancePaths, totalDistance)
	}

	visited[currentRoom.Name] = false
}

func Doublon(path []*Room, pathsList [][]*Room) bool {
	for _, singlePath := range pathsList {
		if len(path) != len(singlePath) {
			continue
		}

		count := 0
		for i := range singlePath {
			if i == 0 {
				continue
			}

			if Compare(path[i], singlePath[i]) {
				count++
			}
		}

		if count == len(singlePath)-1 {
			return true
		}
	}
	return false
}
