package lemin

import "math"

// Crée un nouveau graphe
func NewGraph() *Graph {
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

	g.Edges[from.Name] = append(g.Edges[from.Name], Connection{To: to, Distance: distance})
	g.Edges[to.Name] = append(g.Edges[to.Name], Connection{To: from, Distance: distance})
}

// Récupère les voisins d'une salle
func (g *Graph) GetNeighbors(room *Room) []Connection {
	return g.Edges[room.Name]
	// return un tableau de connection dedans se trouve To qui nous permet de savoir les salles qui sont connectées.
}
