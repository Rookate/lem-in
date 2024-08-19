package lemin

func (data *LeminData) GetPaths(currentRoom *Room) bool {
	//visited[currentRoom] = true

	possibles := make([]*Path, 0)

	for _, p := range data.Paths {
		if p.From != currentRoom {
			continue
		}
		possibles = append(possibles, &p)
	}

	for _, p := range possibles {

		done := data.GetPaths(p.To)

		if done {

		}
	}
}
