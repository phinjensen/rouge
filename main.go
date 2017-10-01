package main

import (
    "fmt"

    bear "bearlibterminal"

    "rogue/maps"
    "rogue/entities"
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

const ROOT = 0
const MAP_LAYER = 1
const STATS_LAYER = 6

const MAP_WIDTH = .6
const MAP_HEIGHT = .8

var bear_width = 80
var bear_height = 25

func draw_map(levelmap maps.MapData) {
    bear.Layer(ROOT)
    bear.BkColor(0x172b56ff)
    bear.ClearArea(
        0,
        0,
        int(MAP_WIDTH * float64(bear_width)),
        int(MAP_HEIGHT * float64(bear_height)),
    )
    bear.Layer(MAP_LAYER)
    for y, row := range levelmap {
        for x, tile := range row {
            bear.Put(x, y, int(tile.Character))
        }
    }
    bear.Put(
        entities.Player.X,
        entities.Player.Y,
        int(entities.Player.Character),
    )
}

func draw_resource_bar(current, max, x, y int, full, empty uint32, message string) {
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
        bar_offset + bar_width/2 - len(bar_string)/2,
        y,
        bar_string,
    )
}

func draw_stats() {
    var cur, max, left_border int
    var full, empty uint32
    bear.Layer(STATS_LAYER)
    left_border = int(MAP_WIDTH * float64(bear_width)) + 1
    bear.Print(
        left_border,
        1,
        "player_name",
    )
    cur, max = entities.Player.Health, entities.Player.MaxHealth
    full = bear.ColorFromName("red")
    empty = bear.ColorFromName("darker red")
    draw_resource_bar(cur, max, left_border, 2, full, empty, "HP:     ")
    cur, max = entities.Player.Energy, entities.Player.MaxEnergy
    full = bear.ColorFromName("blue")
    empty = bear.ColorFromName("darker blue")
    draw_resource_bar(cur, max, left_border, 3, full, empty, "Energy: ")
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
        bear.Layer(ROOT)
        bear.BkColor(0x00000000)
        bear.Clear()
        if bear.Peek() == bear.TK_RESIZED {
            bear.Read()
            bear_width = bear.State(bear.TK_WIDTH)
            bear_height = bear.State(bear.TK_HEIGHT)
            fmt.Println(bear_width, bear_height)
        }
        draw_map(levelmap)
        draw_stats()
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
