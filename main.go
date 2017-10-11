package main

import (
	"math"

	bear "github.com/phinjensen/rouge/bearlibterminal"

	"github.com/phinjensen/rouge/entities"
	"github.com/phinjensen/rouge/maps"
	"github.com/phinjensen/rouge/ui"
)

var movement_keys = []int{
	bear.TK_H, bear.TK_J, bear.TK_K, bear.TK_L,
	bear.TK_Y, bear.TK_U, bear.TK_B, bear.TK_N,
	bear.TK_LEFT, bear.TK_DOWN, bear.TK_UP, bear.TK_RIGHT,
}

func contains(a int, array []int) bool {
	for _, b := range array {
		if a == b {
			return true
		}
	}
	return false
}

func movedirection(key int) string {
	var direction string
	switch key {
	case bear.TK_H, bear.TK_LEFT:
		direction = "w"
	case bear.TK_J, bear.TK_DOWN:
		direction = "s"
	case bear.TK_K, bear.TK_UP:
		direction = "n"
	case bear.TK_L, bear.TK_RIGHT:
		direction = "e"
	case bear.TK_Y:
		direction = "nw"
	case bear.TK_U:
		direction = "ne"
	case bear.TK_B:
		direction = "sw"
	case bear.TK_N:
		direction = "se"
	}
	return direction
}

const VIEW_DISTANCE = 6

func get_visible(levelmap maps.MapData) {
	// MapData.ParseMapFile() adds an empty array at the end (newline at end of file?)
	// which makes this loop break when using map_height and map_width
	for y := 0; y < len(levelmap); y++ {
		for x := 0; x < len(levelmap[y]); x++ {
			levelmap[y][x].Visible = false
		}
	}
	for i := 0.0; i < 360.0; i++ {
		var x = math.Cos(i * (math.Pi / 180.0))
		var y = math.Sin(i * (math.Pi / 180.0))
		var ox = float64(entities.Player.X) + 0.5
		var oy = float64(entities.Player.Y) + 0.5
		for j := 0; j < VIEW_DISTANCE; j++ {
			levelmap[int(oy)][int(ox)].Visible = true
			levelmap[int(oy)][int(ox)].Seen = true
			if levelmap[int(oy)][int(ox)].Walkable == false {
				break
			}
			ox += x
			oy += y
		}
	}
}

func main() {
	bear.Open()
	defer bear.Close()
	bear.Set(
		"window: title='some roguelike thing', resizeable=true; font: mononoki-Regular.ttf, size=12; input: filter={keyboard}",
	)
	bear.Color(bear.ColorFromName("white"))

	levelmap, err := maps.ParseMapFile()
	if err != nil {
		panic(err)
	}

	for {
		bear.Layer(ui.ROOT)
		bear.BkColor(0x00000000)
		bear.Clear()
		if bear.Peek() == bear.TK_RESIZED {
			bear.Read()
			ui.UpdateBearSize()
		}
		get_visible(levelmap)
		ui.DrawMap(levelmap)
		ui.DrawStats()
		if bear.HasInput() {
			key := bear.Read()
			if key == bear.TK_ESCAPE {
				break
			}
			if contains(key, movement_keys) {
				entities.Player.Move(movedirection(key), levelmap)
			}
		}
		bear.Refresh()
	}
}
