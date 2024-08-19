package lemin

import (
	"bufio"
	"common"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
Reads the file named 'fname' and returns its content inside a newly created LeminData structure.

If an error is occured, an error message is returned, along with a nil pointer.
*/
func ParseLeminFile(fname string) (*LeminData, error) {
	fobj, errOpen := os.Open(fname)
	if errOpen != nil {
		return nil, fmt.Errorf("could not open file %s:\n%v", fname, errOpen)
	}

	data := new(LeminData)
	lineScanner := bufio.NewScanner(fobj)
	lineCount := 0
	startNext, endNext := false, false

	for lineScanner.Scan() {
		lineCount++
		line := lineScanner.Text()

		// Comment (or not) parsing bit
		if line[0] == '#' {
			if line[1] != '#' {
				if line[1] == ' ' {
					data = nil
					return nil, fmt.Errorf("illegal syntax at line %d", lineCount)
				}
				continue
			} else {
				switch line[2:] {
				case "start":
					if startNext {
						data = nil
						return nil, fmt.Errorf("duplicate 'start' indicator at line %d", lineCount)
					}
					startNext = true

				case "end":
					if endNext {
						data = nil
						return nil, fmt.Errorf("duplicate 'end' indicator at line %d", lineCount)
					}
					endNext = true

				default:
					return nil, fmt.Errorf("invalid text after '##' at line %d: %s", lineCount, line[2:])
				}
			}
		}

		// Parsing number of ants
		if lineCount == 1 {
			nbr, errAtoi := strconv.Atoi(line)
			if errAtoi != nil {
				data = nil
				return nil, fmt.Errorf("could not read ant amount on first line:\n%v", errAtoi)
			}
			data.AntAmount = uint(nbr)
		}

		// Room declaration parsing bit
		if len(strings.Split(line, " ")) > 1 {
			vals := strings.Split(line, " ")
			if len(vals) != 3 {
				data = nil
				return nil, fmt.Errorf("invalid room declaration format at line %d", lineCount)
			}

			var room Room
			if data.GetRoomFromName(vals[0]) != nil {
				data = nil
				return nil, fmt.Errorf("duplicate room name '%s' at line %d", vals[0], lineCount)
			}
			room.Name = vals[0]

			if !common.IsNumeric(vals[1]) || !common.IsNumeric(vals[2]) {
				data = nil
				return nil, fmt.Errorf("invalid room coordinates at line %d", lineCount)
			}

			x, errX := strconv.Atoi(vals[1])
			if errX != nil {
				data = nil
				return nil, fmt.Errorf("an error occured reading X coord at line %d:\n%v", lineCount, errX)
			}
			room.X = x

			y, errY := strconv.Atoi(vals[2])
			if errY != nil {
				data = nil
				return nil, fmt.Errorf("an error occured reading Y coord at line %d:\n%v", lineCount, errY)
			}
			room.Y = y

			room.AntNb = 0

			if startNext {
				if data.StartRoom != (Room{}) {
					data = nil
					return nil, fmt.Errorf("duplicate start room instantiation at line %d", lineCount)
				}
				data.StartRoom = room
				data.StartRoom.AntNb = data.AntAmount
				startNext = false

				for i := range data.AntAmount {
					data.AntList = append(data.AntList, Ant{
						Name:          "L" + strconv.Itoa(int(i+1)),
						OccupyingRoom: &data.StartRoom,
					})
				}
			} else if endNext {
				if data.EndRoom != (Room{}) {
					data = nil
					return nil, fmt.Errorf("duplicate end room instantiation at line %d", lineCount)
				}
				data.EndRoom = room
				endNext = false
			} else {
				data.OtherRooms = append(data.OtherRooms, room)
			}
		}

		// Path declaration parsing bit
		if len(strings.Split(line, "-")) > 1 {
			vals := strings.Split(line, "-")
			if len(vals) != 2 {
				data = nil
				return nil, fmt.Errorf("invalid path declaration at line %d", lineCount)
			}

			var path Path

			from := data.GetRoomFromName(vals[0])
			if from == nil {
				data = nil
				return nil, fmt.Errorf("unknown room '%s' at line %d", vals[0], lineCount)
			}
			path.From = from

			to := data.GetRoomFromName(vals[1])
			if to == nil {
				data = nil
				return nil, fmt.Errorf("unknown room '%s' at line %d", vals[1], lineCount)
			}
			path.To = to

			path.Distance = math.Sqrt(float64((to.X-from.X)*(to.X-from.X)) + float64((to.Y-from.Y)*(to.Y-from.Y)))

			data.Paths = append(data.Paths, path)
		}

		data.FileContent += line + "\n"
	}

	if lineScanner.Err() != nil {
		data = nil
		return nil, fmt.Errorf("could not scan line %d in file %s:\n%v", lineCount+1, fname, lineScanner.Err().Error())
	}

	return data, nil
}
