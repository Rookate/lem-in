package lemin

import "math"

/*
	La structure Graph utilise exclusivement les données déjà présentes dans un objet LeminData.
*/

// Crée un nouveau graphe
func NewGraph(data *LeminData) *Graph {
	if data == nil {
		return nil
	}

	return &Graph{
		Rooms: make(map[string]*Room),
		Edges: make(map[string][]Connection),
	}
}

// Ajoute une salle au graphe
func (g *Graph) AddRoom(room *Room) {
	g.Rooms[room.Name] = room
}

// Ajoute une connection entre from et to
func (g *Graph) AddEdge(from *Room, to *Room) {
	distance := math.Sqrt(float64((to.X-from.X)*(to.X-from.X)) + float64((to.Y-from.Y)*(to.Y-from.Y)))

	g.Edges[from.Name] = append(g.Edges[from.Name], Connection{From: from, To: to, Distance: distance})
	g.Edges[to.Name] = append(g.Edges[to.Name], Connection{To: from, From: to, Distance: distance})
}

// Récupère les voisins d'une salle
func (g *Graph) GetNeighbors(room *Room) []Connection {
	return g.Edges[room.Name]
	// return un tableau de connection dedans se trouve To qui nous permet de savoir les salles qui sont connectées.
}
