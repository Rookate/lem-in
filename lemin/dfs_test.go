package lemin

import (
	"fmt"
	"testing"
)

func TestMoveAnts(t *testing.T) {
	// Initialisation des salles et chemins
	startRoom := &Room{Name: "Start", Occupied: true}
	endRoom := &Room{Name: "End", Occupied: false}
	midRoom1 := &Room{Name: "Mid1", Occupied: false}
	midRoom2 := &Room{Name: "Mid2", Occupied: false}

	paths := [][]*Room{
		{startRoom, midRoom1, endRoom},
		{startRoom, midRoom2, endRoom},
	}

	// Initialisation des fourmis
	ants := []Ant{
		{Name: "Ant 1", OccupyingRoom: startRoom},
		{Name: "Ant 2", OccupyingRoom: startRoom},
	}

	data := LeminData{
		AntList:    ants,
		StartRoom:  *startRoom,
		EndRoom:    *endRoom,
		OtherRooms: []Room{*midRoom1, *midRoom2},
	}

	pathFinder := &PathFinder{AllPaths: paths}

	// Appel de la fonction à tester avec une limite de 10 tours
	const maxTurns = 20
	MoveAnts(pathFinder, &data, maxTurns)

	// Vérification du résultat attendu après l'exécution
	// On doit s'assurer que les fourmis ont atteint la salle de fin
	for _, ant := range ants {
		if ant.OccupyingRoom != endRoom {
			t.Errorf("%s should be in the End room, but is in %s", ant.Name, ant.OccupyingRoom.Name)
		}
	}

	// Vérifier que la salle de fin est occupée par les fourmis
	if !endRoom.Occupied {
		t.Errorf("End room should be occupied, but it is not")
	}

	// Vérification que toutes les fourmis ont effectivement bougé
	if len(data.AntList) != 2 {
		t.Errorf("Expected 2 ants, but found %d", len(data.AntList))
	}
}

// Exemple de test unitaire simple
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

	// Création des connexions
	connAB := Connection{From: roomA, To: roomB, Distance: 1}
	connAC := Connection{From: roomA, To: roomC, Distance: 1}
	connBD := Connection{From: roomB, To: roomD, Distance: 1}
	connCD := Connection{From: roomC, To: roomD, Distance: 1}

	// Initialisation des connexions dans LeminData
	data := &LeminData{
		ConnectionList: []Connection{connAB, connAC, connBD, connCD},
		EndRoom:        *roomD,
	}

	// Initialisation du PathFinder et appel DFS
	pathFinder := &PathFinder{}
	visited := make(map[string]bool)
	var path []*Room

	DFS(data, roomA, visited, path, pathFinder)

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
	expectedPaths := 2 // Il y a deux chemins possibles : A -> B -> D et A -> C -> D
	if len(pathFinder.AllPaths) != expectedPaths {
		t.Errorf("Nombre de chemins incorrect : attendu %d, obtenu %d", expectedPaths, len(pathFinder.AllPaths))
	}

	// Vérification des chemins spécifiques trouvés
	expectedPath1 := []string{"0", "1", "3"}
	expectedPath2 := []string{"0", "2", "3"}

	for _, p := range pathFinder.AllPaths {
		var pathNames []string
		for _, r := range p {
			pathNames = append(pathNames, r.Name)
		}
		if !equalPaths(pathNames, expectedPath1) && !equalPaths(pathNames, expectedPath2) {
			t.Errorf("Chemin inattendu trouvé : %v", pathNames)
		}
	}
}

// Fonction d'aide pour comparer deux chemins
func equalPaths(path1, path2 []string) bool {
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
