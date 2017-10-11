package ui

import (
	"fmt"

	bear "github.com/phinjensen/rouge/bearlibterminal"

	"github.com/phinjensen/rouge/entities"
	"github.com/phinjensen/rouge/maps"
)

const ROOT = 0
const MAP_LAYER = 1
const STATS_LAYER = 6

const MAP_WIDTH = .6
const MAP_HEIGHT = .8

var bear_width = 80
var bear_height = 25

func UpdateBearSize() {
	bear_width = bear.State(bear.TK_WIDTH)
	bear_height = bear.State(bear.TK_HEIGHT)
}

func DrawMap(levelmap maps.MapData) {
	var original_color = uint32(bear.State(bear.TK_COLOR))
	var original_bk_color = uint32(bear.State(bear.TK_BKCOLOR))

	bear.Layer(ROOT)
	bear.BkColor(0x172b56ff)
	var view_width = int(MAP_WIDTH * float64(bear_width))
	var view_height = int(MAP_HEIGHT * float64(bear_height))
	bear.ClearArea(0, 0, view_width, view_height)

	bear.Layer(MAP_LAYER)
	for y := 0; y < view_height; y++ {
		map_y := entities.Player.Y - view_height/2 + y
		if map_y < 0 || map_y >= len(levelmap) {
			continue
		}
		for x := 0; x < view_width; x++ {
			map_x := entities.Player.X - view_width/2 + x
			if map_x < 0 || map_x >= len(levelmap[map_y]) {
				continue
			}
			tile := levelmap[map_y][map_x]
			if tile.Seen {
				bear.Color(tile.Color())
				bear.Put(x, y, int(tile.Character))
			}
		}
	}
	bear.Color(0xffd8af5f)
	bear.Put(
		view_width/2,
		view_height/2,
		int(entities.Player.Character),
	)
	bear.Color(original_color)
	bear.BkColor(original_bk_color)
}

func DrawResourceBar(current, max, x, y int, full, empty uint32, message string) {
	fraction := float64(current) / float64(max)
	bar_string := fmt.Sprintf(
		"%d/%d (%.f%%)",
		current,
		max,
		fraction*100,
	)
	bar_offset := x + len(message)
	bar_width := bear_width - bar_offset - 1
	bear.Layer(ROOT)
	bear.BkColor(empty)
	bear.ClearArea(bar_offset, y, bar_width, 1)
	bear.BkColor(full)
	bear.ClearArea(
		bar_offset,
		y,
		int(fraction*float64(bar_width)),
		1,
	)
	bear.Layer(STATS_LAYER)
	bear.Print(x, y, message)
	bear.Print(
		bar_offset+bar_width/2-len(bar_string)/2,
		y,
		bar_string,
	)
}

func DrawStats() {
	var cur, max, left_border int
	var full, empty uint32
	bear.Layer(STATS_LAYER)
	left_border = int(MAP_WIDTH*float64(bear_width)) + 1
	bear.Print(
		left_border,
		1,
		"player_name",
	)
	cur, max = entities.Player.Health, entities.Player.MaxHealth
	full = bear.ColorFromName("red")
	empty = bear.ColorFromName("darker red")
	DrawResourceBar(cur, max, left_border, 2, full, empty, "HP:     ")
	cur, max = entities.Player.Energy, entities.Player.MaxEnergy
	full = bear.ColorFromName("blue")
	empty = bear.ColorFromName("darker blue")
	DrawResourceBar(cur, max, left_border, 3, full, empty, "Energy: ")
}
