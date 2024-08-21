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

	antCount := &LeminData{AntAmount: 2}
	pathFinder := &PathFinder{AllPaths: paths}

	// Appel de la fonction à tester
	MoveAnts(pathFinder, antCount, ants)

	// Vérification du résultat attendu
	if ants[0].OccupyingRoom != endRoom {
		t.Errorf("Ant 1 should be in the End room, but is in %s", ants[0].OccupyingRoom.Name)
	}

	if ants[1].OccupyingRoom != endRoom {
		t.Errorf("Ant 2 should be in the End room, but is in %s", ants[1].OccupyingRoom.Name)
	}

	if !endRoom.Occupied {
		t.Errorf("End room should be occupied, but it is not")
	}
}

func TestCreateAnts(t *testing.T) {
	// Setup des données de test
	startRoom := &Room{Name: "StartRoom", X: 0, Y: 0}
	antCount := &LeminData{AntAmount: 3} // Créons 3 fourmis

	// Appel de la fonction à tester
	ants := CreateAnts(antCount, startRoom)

	// Vérifications
	if len(ants) != 3 {
		t.Fatalf("Expected 3 ants, but got %d", len(ants))
	}

	for i, ant := range ants {
		expectedName := fmt.Sprintf("Ant %d", i+1)
		if ant.Name != expectedName {
			t.Errorf("Ant %d: Expected name %s, but got %s", i, expectedName, ant.Name)
		}
		if ant.OccupyingRoom != startRoom {
			t.Errorf("Ant %d: Expected occupying room %v, but got %v", i, startRoom, ant.OccupyingRoom)
		}
	}
}

func TestDFS(t *testing.T) {
	// Création des salles
	roomA := Room{Name: "0"}
	roomB := Room{Name: "1"}
	roomC := Room{Name: "2"}
	roomD := Room{Name: "3"}

	// Création des chemins (Paths)
	paths := []Path{
		{From: &roomA, To: &roomB},
		{From: &roomA, To: &roomD},
		{From: &roomB, To: &roomC},
		{From: &roomD, To: &roomC},
	}

	// Map pour suivre les salles visitées
	PathFinder := &PathFinder{}
	visited := make(map[string]bool)
	var path []*Room

	// Appel de DFS
	// Création du graphe et ajout des salles et chemins
	graph := NewGraph()
	graph.AddRoom(&roomA)
	graph.AddRoom(&roomB)
	graph.AddRoom(&roomC)
	graph.AddRoom(&roomD)
	for _, path := range paths {
		graph.AddEdge(path.From, path.To)
	}

	DFS(&roomA, &roomD, visited, path, PathFinder, graph)

	// Imprimer les chemins trouvés
	fmt.Println("Chemins trouvés :")
	for i, p := range PathFinder.AllPaths {
		var pathNames []string
		for _, r := range p {
			pathNames = append(pathNames, r.Name)
		}
		fmt.Printf("Chemin %d: %v\n", i+1, pathNames)
	}

	// Vérification du nombre de chemins trouvés
	expectedPaths := 2 // Il y a deux chemins possibles : A -> B -> C et A -> D
	if len(PathFinder.AllPaths) != expectedPaths {
		t.Errorf("Nombre de chemins incorrect : attendu %d, obtenu %d", expectedPaths, len(PathFinder.AllPaths))
	}

	// Vérification des chemins spécifiques trouvés
	expectedPath1 := []string{"0", "1", "2", "3"}
	expectedPath2 := []string{"0", "3"}

	for _, p := range PathFinder.AllPaths {
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
