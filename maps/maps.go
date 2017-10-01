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

type MapData [][]Tile

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
		var row_tiles []Tile
		for _, tile := range strings.Split(row, "") {
			row_tiles = append(row_tiles, tiles[tile])
		}
		rows = append(rows, row_tiles)
	}
	return rows, nil
}
