package api

import (
	"reflect"
)

type Player struct {
	Number string
	Name   string
	Fide   string
	Rating string
}

type Players struct {
	constructor
}

func NewPlayers(table [][]string) *Players {
	var players Players
	players.config = map[string]string{
		"Имя":      "Name",
		"Ном.":     "Number",
		"код FIDE": "Fide",
		"Рейт.":    "Rating",
	}
	players.reflect = reflect.TypeOf(Player{})
	players.setIndexes(table[0])
	players.fill(table)
	return &players
}
