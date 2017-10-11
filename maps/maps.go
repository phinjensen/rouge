package maps

import (
	"io/ioutil"
	"strings"
)

// Tiles

type Tile struct {
	Character rune
	Walkable  bool
}

var tiles map[string]Tile = map[string]Tile{
	" ": Tile{' ', false},
	"#": Tile{'#', false},
	".": Tile{'.', true},
}

// Map data

type GameTile struct {
	Tile
	Visible bool
	Seen    bool
}

func (t GameTile) Color() uint32 {
	var color uint32 = 0xffffffff
	if t.Visible {
		color = 0xffffffff
	} else {
		color = 0x88ffffff
	}
	return color
}

type MapData [][]GameTile

func (m MapData) Dimensions() (int, int) {
	if len(m) > 0 {
		return len(m), len(m[0])
	}
	return 0, 0
}

func ParseMapFile() (MapData, error) {
	data, err := ioutil.ReadFile("data/maps/f1.map")
	if err != nil {
		panic(err)
	}
	var mapstring = string(data)
	var rows MapData
	for _, row := range strings.Split(mapstring, "\n") {
		var row_tiles []GameTile
		for _, tile := range strings.Split(row, "") {
			row_tiles = append(row_tiles, GameTile{tiles[tile], false, false})
		}
		rows = append(rows, row_tiles)
	}
	return rows, nil
}
