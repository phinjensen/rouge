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
	"w": Tile{'#', false},
	"f": Tile{'.', true},
}

// Map data

type MapData [][]Tile

func ParseMapFile() (MapData, error) {
	data, err := ioutil.ReadFile("data/maps/f1.map")
	if err != nil {
		panic(err)
	}
	var mapstring = string(data)
	var rows MapData
	for _, row := range strings.Split(mapstring, "\n") {
		var row_tiles []Tile
		for _, tile := range strings.Split(row, ",") {
			row_tiles = append(row_tiles, tiles[tile])
		}
		rows = append(rows, row_tiles)
	}
	return rows, nil
}
