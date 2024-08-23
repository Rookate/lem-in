package lemin

import (
	"fmt"
	"testing"
)

func TestMoveAnts(t *testing.T) {
	// Création des salles
	roomA := &Room{Name: "0"}
	roomB := &Room{Name: "1"}
	roomC := &Room{Name: "2"}
	roomD := &Room{Name: "3"}
	roomE := &Room{Name: "4"}

	// Création des connexions
	connAB := Connection{From: roomA, To: roomB, Distance: 1}
	connAD := Connection{From: roomA, To: roomD, Distance: 1}
	connBC := Connection{From: roomB, To: roomC, Distance: 1}
	connDC := Connection{From: roomD, To: roomC, Distance: 1}
	connAE := Connection{From: roomA, To: roomE, Distance: 1}
	connED := Connection{From: roomE, To: roomD, Distance: 1}

	// Création des chemins pour PathFinder
	allPaths := [][]*Room{
		{roomA, roomB, roomC, roomD}, // Chemin 0
		{roomA, roomE, roomD},        // Chemin 1
		{roomA, roomD},
	}

	allDistances := []float64{3, 2, 1} // Exemple de distances associées

	// Initialisation de PathFinder
	pathFinder := &PathFinder{
		AllPaths:         allPaths,
		AllDistancePaths: allDistances,
	}

	// Création des fourmis
	ants := []Ant{
		{Name: "Ant1", OccupyingRoom: roomA},
		{Name: "Ant2", OccupyingRoom: roomA},
	}

	// Initialisation des données
	data := &LeminData{
		AntList:        ants,
		StartRoom:      *roomA,
		EndRoom:        *roomD,
		OtherRooms:     []Room{*roomB, *roomC, *roomE},
		ConnectionList: []Connection{connAB, connAD, connBC, connDC, connAE, connED},
	}

	// Appel de la fonction à tester avec une limite de 20 tours
	SortPaths(pathFinder)
	const maxTurns = 20
	MoveAnts(pathFinder, data, maxTurns)

	// Vérification que toutes les fourmis ont atteint la salle de fin
	for _, ant := range data.AntList {
		if ant.OccupyingRoom != &data.EndRoom {
			t.Errorf("%s should be in the End room, but is in %s", ant.Name, ant.OccupyingRoom.Name)
		}
	}

	// Vérification que toutes les fourmis ont effectivement bougé
	if len(data.AntList) != 2 {
		t.Errorf("Expected 2 ants, but found %d", len(data.AntList))
	}
}

func TestGetTunnel(t *testing.T) {
	roomA := &Room{Name: "A"}
	roomB := &Room{Name: "B"}
	roomC := &Room{Name: "C"}

	data := &LeminData{
		ConnectionList: []Connection{
			{From: roomA, To: roomB},
			{From: roomB, To: roomC},
		},
	}

	tunnel := data.GetTunnel(roomA, roomB)
	if tunnel == nil {
		t.Error("Expected tunnel between roomA and roomB")
	}
}

func TestCreateAnts(t *testing.T) {
	startRoom := Room{Name: "StartRoom", X: 0, Y: 0}
	data := LeminData{
		StartRoom: startRoom,
	}

	// Appel de la fonction à tester avec le nombre de fourmis à créer
	data.CreateAnts(3)

	// Vérifications
	if len(data.AntList) != 3 {
		t.Fatalf("Expected 3 ants, but got %d", len(data.AntList))
	}

	for i, ant := range data.AntList {
		expectedName := fmt.Sprintf("L%d", i+1)
		if ant.Name != expectedName {
			t.Errorf("L%d: Expected name %s, but got %s", i+1, expectedName, ant.Name)
		}
		if *ant.OccupyingRoom != startRoom {
			t.Errorf("Ant %d: Expected occupying room %v, but got %v", i+1, startRoom, ant.OccupyingRoom)
		}
	}
}

func TestDFS(t *testing.T) {
	// Création des salles
	roomA := &Room{Name: "0"}
	roomB := &Room{Name: "1"}
	roomC := &Room{Name: "2"}
	roomD := &Room{Name: "3"}
	roomE := &Room{Name: "4"}

	// Création des connexions
	connAB := Connection{From: roomA, To: roomB, Distance: 1}
	connAD := Connection{From: roomA, To: roomD, Distance: 1}
	connBC := Connection{From: roomB, To: roomC, Distance: 1}
	connDC := Connection{From: roomD, To: roomC, Distance: 1}
	connAE := Connection{From: roomA, To: roomE, Distance: 1}
	connED := Connection{From: roomE, To: roomD, Distance: 1}

	// Initialisation des connexions dans LeminData
	data := &LeminData{
		ConnectionList: []Connection{connAB, connAD, connBC, connDC, connAE, connED},
		EndRoom:        *roomD,
	}

	// Initialisation du PathFinder et appel DFS
	pathFinder := &PathFinder{}
	visited := make(map[string]bool)
	var path []*Room

	DFS(data, roomA, roomD, visited, path, pathFinder)

	// Imprimer les chemins trouvés
	fmt.Println("Chemins trouvés :")
	for i, p := range pathFinder.AllPaths {
		var pathNames []string
		for _, r := range p {
			pathNames = append(pathNames, r.Name)
		}
		fmt.Printf("Chemin %d: %v\n", i+1, pathNames)
	}

	// Vérification du nombre de chemins trouvés
	expectedPaths := 3 // Il y a deux chemins possibles : A -> B -> D et A -> C -> D
	if len(pathFinder.AllPaths) != expectedPaths {
		t.Errorf("Nombre de chemins incorrect : attendu %d, obtenu %d", expectedPaths, len(pathFinder.AllPaths))
	}

	// Vérification des chemins spécifiques trouvés
	expectedPath1 := []string{"0", "3"}
	expectedPath2 := []string{"0", "1", "2", "3"}
	expectedPath3 := []string{"0", "4", "3"}

	for _, p := range pathFinder.AllPaths {
		var pathNames []string
		for _, r := range p {
			pathNames = append(pathNames, r.Name)
		}
		if !equalPath(pathNames, expectedPath1) && !equalPath(pathNames, expectedPath2) && !equalPath(pathNames, expectedPath3) {
			t.Errorf("Chemin inattendu trouvé : %v", pathNames)
		}
	}
}

// Fonction d'aide pour comparer deux chemins
func equalPath(path1, path2 []string) bool {
	if len(path1) != len(path2) {
		return false
	}
	for i := range path1 {
		if path1[i] != path2[i] {
			return false
		}
	}
	return true
}

func TestGetNeighbors(t *testing.T) {
	// Création des salles pour le test
	room0 := &Room{Name: "0"}
	room1 := &Room{Name: "1"}
	room2 := &Room{Name: "2"}
	room3 := &Room{Name: "3"}

	// Création des connexions pour le test
	conn01 := Connection{From: room0, To: room1, Distance: 1.0}
	conn02 := Connection{From: room1, To: room2, Distance: 1.0}
	conn03 := Connection{From: room0, To: room3, Distance: 2.0}
	conn04 := Connection{From: room2, To: room3, Distance: 1.0}

	// Initialisation des données de test
	data := LeminData{
		ConnectionList: []Connection{conn01, conn02, conn03, conn04},
	}

	// Fonction pour tester les voisins
	checkNeighbors := func(room *Room, expectedNeighbors []*Room) {
		neighbors := data.GetNeighbors(room)
		if len(neighbors) != len(expectedNeighbors) {
			t.Errorf("Expected %d neighbors for room %s, got %d", len(expectedNeighbors), room.Name, len(neighbors))
		}

		// Vérification que chaque voisin attendu est présent
		for _, expectedNeighbor := range expectedNeighbors {
			found := false
			for _, neighbor := range neighbors {
				if neighbor == expectedNeighbor {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Expected neighbor %s for room %s not found", expectedNeighbor.Name, room.Name)
			}
		}
	}

	// Tests des voisins pour chaque salle
	checkNeighbors(room0, []*Room{room1, room3})
	checkNeighbors(room1, []*Room{room0, room2})
	checkNeighbors(room2, []*Room{room1, room3})
	checkNeighbors(room3, []*Room{room0, room2})
}
