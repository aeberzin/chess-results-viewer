package api

import (
	"reflect"
)

type Competitor struct {
	Start  string
	Finish string
	Name   string
	Result string
	Rate1  string
	Rate2  string
	Rate3  string
}

type Competitors struct {
	constructor
}

func NewCompetitors(table [][]string) *Competitors {
	var competitors Competitors
	competitors.config = map[string]string{
		"Имя":              "Name",
		"Ст.ном":           "Finish",
		"Ст.ном.":          "Start",
		"Очки":             "Result",
		withSpaces("Доп1"): "Rate1",
		withSpaces("Доп2"): "Rate2",
		withSpaces("Доп3"): "Rate3",
	}
	competitors.reflect = reflect.TypeOf(Competitor{})
	competitors.setIndexes(table[0])
	competitors.fill(table)
	return &competitors
}

func withSpaces(s string) string {
	space := string([]byte{194, 160})
	return space + s + space
}
