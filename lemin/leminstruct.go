package lemin

import (
	"fmt"
	"lemin/common"
)

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
-3 == Error

-2 == EndRoom

-1 == StartRoom

Else == Room figuring in data.OtherRooms
*/
func (data *LeminData) GetRoomIndexFromName(name string) int {
	if name == "" {
		return -3
	}

	if *data.GetRoomFromName(name) == data.StartRoom {
		return -1
	}

	if *data.GetRoomFromName(name) == data.EndRoom {
		return -2
	}

	for i := range data.OtherRooms {
		if data.OtherRooms[i].Name == name {
			return i
		}
	}

	return -3
}

func Compare(r1, r2 *Room) bool {
	return r1.Name == r2.Name && r1.X == r2.X && r1.Y == r2.Y && r1.Occupied == r2.Occupied
}

func (data *LeminData) GetTunnel(from *Room, to *Room) *Tunnel {
	for _, connection := range data.TunnelList {
		if (connection.From.Name == from.Name && connection.To.Name == to.Name) || (connection.From.Name == to.Name && connection.To.Name == from.Name) {
			return &connection
		}
	}
	return nil
}

func (data *LeminData) GetNeighbors(room *Room) []*Room {
	neighbors := make([]*Room, 0)

	for _, tunnel := range data.TunnelList {
		if *tunnel.From == *room && *tunnel.To != data.StartRoom {
			neighbors = append(neighbors, tunnel.To)
		} else if *tunnel.To == *room && *tunnel.From != data.StartRoom {
			neighbors = append(neighbors, tunnel.From)
		}
	}

	return neighbors
}

/*
Returns whether or not the LeminData object is a valid data structure.

In the 'false' case, the string given as returned value gives the reason.

------------------------------------------------------------

Failing cases are :

- Amount of ants being 0;

- Any room beginning with 'L' specifically;

- Any room's name being the same as another's;

- Any path linking a room to itself;

- Any path containing a room whose name does not exist;

- Any path that is the opposite way of an already existing path.
*/
func (lem *LeminData) IsValidData() (bool, string) {
	if len(lem.AntList) == 0 {
		return false, "amount of ants cannot be zero"
	}

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

	for i, name := range names {
		if common.IndexOf(names[i+1:], name) != -1 {
			return false, fmt.Sprintf("duplicate room name '%s'", name)
		}
	}

	for _, path1 := range lem.TunnelList {
		if *path1.From == *path1.To {
			return false, "paths cannot link a room to itself"
		}

		if common.IndexOf(names, path1.From.Name) == -1 {
			return false, fmt.Sprintf("unknown room '%s' in path '%s-%s'", path1.From.Name, path1.From.Name, path1.To.Name)
		}

		if common.IndexOf(names, path1.To.Name) == -1 {
			return false, fmt.Sprintf("unknown room '%s' in path '%s-%s'", path1.To.Name, path1.From.Name, path1.To.Name)
		}
	}

	return true, ""
}
