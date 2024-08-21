package lemin

import (
	"fmt"
)

func Resolve(data *LeminData) error {
	{
		dataOk, errMsg := data.IsValidData()
		if !dataOk {
			return fmt.Errorf("invalid data structure:\n%s", errMsg)
		}
	}

	buffer := make([]Path, 0)
	solutions := make([][]Path, 0)
	visited := make(map[Room]bool, len(data.OtherRooms)+2)

	for _, val := range data.OtherRooms {
		visited[val] = false
	}

	visited[data.StartRoom] = false
	visited[data.EndRoom] = false

	_, _, solutions, _ = data.GetSolutions(&data.StartRoom, buffer, solutions, visited)

	if len(solutions) == 0 {
		return fmt.Errorf("no solution was found with this configuration")
	}

	fmt.Println("Solutions:")
	for i, paths := range solutions {
		fmt.Print("Solution ", i+1, ":\n")
		for i, path := range paths {
			fmt.Printf("- Step %d: \"%s\" (%d;%d) -> \"%s\" (%d;%d)\n",
				i+1,
				path.From.Name, path.From.X, path.From.Y,
				path.To.Name, path.To.X, path.To.Y)
		}
		fmt.Printf("Total distance : %.4f\n", GetTotalDistance(paths))
	}

	return nil
}

func (data *LeminData) GetSolutions(currentRoom *Room, tmp []Path, solutions [][]Path, visited map[Room]bool) (bool, []Path, [][]Path, map[Room]bool) {
	// This room is now visited
	visited[*currentRoom] = true

	// If this room is the finish line, it is a correct solution.
	if currentRoom == &data.EndRoom {
		return true, tmp, solutions, visited
	}

	possibles := make([]*Path, 0)

	/*
		All possibilities are:
		Paths where initial room is the current room
			OR where source is the end room AND destination is the current room
	*/
	for _, p := range data.Paths {
		if !currentRoom.Compare(p.From) && (!data.EndRoom.Compare(p.From) || !currentRoom.Compare(p.To)) {
			continue
		}
		possibles = append(possibles, &p)
	}

	// If there is no possibility, it is a dead end.
	if len(possibles) == 0 {
		return false, tmp, solutions, visited
	}

	for _, p := range possibles {
		// To avoid loops, evaluate next path if the path's destination has already been visited.
		if visited[*p.To] {
			continue
		}

		/*
			Path is added to the list,
			and next call starts from its destination.
		*/
		tmp = append(tmp, *p)
		status, newtmp, newsolutions, visited := data.GetSolutions(p.To, tmp, solutions, visited)

		/*
			If next call returns false, it's either a dead end or a loop.
			In that case, backtrack and evaluate the next path.
		*/
		if !status {
			visited[*tmp[len(tmp)-1].To] = false
			tmp = tmp[:len(tmp)-1]

			if p == possibles[len(possibles)-1] {
				return status, tmp, solutions, visited
			}
		} else {
			/*
				Otherwise, it's a solution.
				In that case,
				if the path's source is the start room,
				THEN the program found a solution in previous path, so try the other connections.

				If the destination is the end room,
				THEN the solution is a single path between start and end. Add it to the solutions.
			*/
			if p.From == &data.StartRoom {
				if p.To == &data.EndRoom {
					newsolutions = append(newsolutions, tmp)
				}
				solutions = newsolutions
				continue
			}

			/*
				If the path's source is NOT the start room
				AND its destination is NOT the finish line,
				THEN keep backtracking until start room.
			*/
			if p.To != &data.EndRoom {
				return status, newtmp, newsolutions, visited
			}

			/*
				If the path's source is NOT the start room
				AND its destination is the finish line,
				THEN a solution has been found. Add it to the list and start backtracking to start room.
			*/
			newsolutions = append(newsolutions, newtmp)
			return status, []Path{}, newsolutions, visited
		}
	}

	return false, tmp, solutions, visited
}
