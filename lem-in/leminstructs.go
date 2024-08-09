package lemin

import (
	"common"
	"fmt"
)

type Room struct {
	Name string
	X    int
	Y    int
}

type Path struct {
	From Room
	To   Room
}

type LeminData struct {
	FileContent string // Content of the read file saved per project requirements
	AntAmount   uint   // Amount of ants present in the farm
	StartRoom   Room   // Room where all the ants start into
	EndRoom     Room   // Room where ants must finish
	OtherRooms  []Room // Array of checkpoint rooms
	Paths       []Path // Array of unidirectional paths
}

func (lem *LeminData) GetRoomFromName(s string) *Room {
	if len(s) == 0 {
		return nil
	}

	if lem.StartRoom.Name == s {
		return &lem.StartRoom
	}

	if lem.EndRoom.Name == s {
		return &lem.EndRoom
	}

	for _, room := range lem.OtherRooms {
		if room.Name == s {
			return &room
		}
	}

	return nil
}

/*
Returns whether or not the LeminData object is a valid data structure.

In the 'false' case, the string given as returned value gives the reason.

------------------------------------------------------------

Failing cases are :

- Amount of ants being 0;

- Any room beginning with 'L' specifically;

- Any path linking a room to itself;

- Any path containing a room whose name does not exist;

- Any path that is the opposite way of an already existing path.
*/
func (lem *LeminData) IsValidData() (bool, string) {
	if lem.AntAmount == 0 {
		return false, "amount of ants cannot be zero"
	}

	// Is this really an error?
	// if len(lem.OtherRooms) == 0 {
	// 	return false, "there must be other rooms in the farm"
	// }

	var names []string

	if lem.StartRoom.Name[0] == 'L' {
		return false, "start room cannot begin with 'L'"
	}
	names = append(names, lem.StartRoom.Name)

	if lem.EndRoom.Name[0] == 'L' {
		return false, "end room cannot begin with 'L'"
	}
	names = append(names, lem.EndRoom.Name)

	for _, room := range lem.OtherRooms {
		if room.Name[0] == 'L' {
			return false, fmt.Sprintf("room '%s' cannot begin with 'L'", room.Name)
		}
		names = append(names, room.Name)
	}

	for _, path1 := range lem.Paths {
		if path1.From.Name == path1.To.Name {
			return false, "paths cannot link a room to itself"
		}

		if common.IndexOf(names, path1.From.Name) == -1 {
			return false, fmt.Sprintf("unknown room '%s' in path '%s-%s'", path1.From.Name, path1.From.Name, path1.To.Name)
		}

		if common.IndexOf(names, path1.To.Name) == -1 {
			return false, fmt.Sprintf("unknown room '%s' in path '%s-%s'", path1.To.Name, path1.From.Name, path1.To.Name)
		}

		// Searching for eventual bidirectional paths
		for _, path2 := range lem.Paths {
			if path1.From.Name == path2.To.Name && path1.To.Name == path2.From.Name {
				return false, fmt.Sprintf("path '%s-%s' is the opposite of already existing path '%s-%s'", path2.From.Name, path2.To.Name, path1.From.Name, path1.To.Name)
			}
		}
	}

	return true, ""
}

// Not necessary + not as easy as I thought
// func (lem *LeminData) DrawFarm() (bool, error) {
// 	{
// 		ok, err := lem.IsValidData()
// 		if !ok {
// 			return ok, fmt.Errorf("invalid data structure:\n%s", err)
// 		}
// 	}

// 	// TODO : Draw the best farm ever

// 	return true, nil
// }
