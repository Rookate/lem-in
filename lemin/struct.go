package lemin

type Room struct {
	AntNb    uint
	Name     string
	X        int
	Y        int
	Occupied bool
}

// Structure où on va stocker la connection entre deux rooms + la distance entre elles
type Tunnel struct {
	From     *Room
	To       *Room
	Distance float64
}

type Ant struct {
	Name          string
	OccupyingRoom *Room
}

type LeminData struct {
	FileContent string // Content of the read file saved per project requirements
	AntList     []Ant
	StartRoom   Room   // Room where all the ants start into
	EndRoom     Room   // Room where ants must finish
	OtherRooms  []Room // Array of checkpoint rooms
	TunnelList  []Tunnel
}

type PathFinder struct {
	AllPaths         [][]*Room
	AllDistancePaths []float64
}
