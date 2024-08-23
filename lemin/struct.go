package lemin

type Room struct {
	AntNb    uint
	Name     string
	X        int
	Y        int
	Occupied bool
}

type Ant struct {
	Name          string
	Position      int
	Path          []*Room
	OccupyingRoom *Room
}

type LeminData struct {
	FileContent    string // Content of the read file saved per project requirements
	AntList        []Ant
	StartRoom      Room   // Room where all the ants start into
	EndRoom        Room   // Room where ants must finish
	OtherRooms     []Room // Array of checkpoint rooms
	ConnectionList []Connection
}

type PathFinder struct {
	AllPaths         [][]*Room
	AllDistancePaths []float64
}

type Graph struct {
	Rooms map[string]*Room
	Edges map[string][]Connection
}

// Structure où on va stocker la connection entre deux rooms + la distance entre elles
type Connection struct {
	From     *Room
	To       *Room
	Distance float64
}
