package entities

import (
	"github.com/phinjensen/rouge/maps"
)

func isValidMove(e Entity, levelmap maps.MapData) bool {
	return e.Y < len(levelmap) &&
		e.X < len(levelmap[e.Y]) &&
		levelmap[e.Y][e.X].Walkable
}

type Entity struct {
	X         int
	Y         int
	Character rune
	Health    int
	MaxHealth int
	Energy    int
	MaxEnergy int
}

func (e *Entity) Move(direction string, levelmap maps.MapData) {
	x, y := e.X, e.Y
	switch direction {
	case "w":
		e.X--
	case "s":
		e.Y++
	case "n":
		e.Y--
	case "e":
		e.X++
	case "nw":
		e.X--
		e.Y--
	case "ne":
		e.X++
		e.Y--
	case "sw":
		e.X--
		e.Y++
	case "se":
		e.X++
		e.Y++
	}
	if !isValidMove(*e, levelmap) {
		e.X = x
		e.Y = y
	}
}

var Player = Entity{4, 3, '@', 30, 40, 2, 17}
