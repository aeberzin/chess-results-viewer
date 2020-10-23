package api

import (
	"log"
	"reflect"
)

type Pair struct {
	Table   string
	WName   string
	BName   string
	WRating string
	BRating string
	WRes    string
	BRes    string
	Result  string
}

type Pairs struct {
	constructor
}

func NewPairs(table [][]string) *Pairs {
	var pairs Pairs
	pairs.reflect = reflect.TypeOf(Pair{})
	pairs.setIndexes(table[0])
	pairs.fill(table)
	return &pairs
}

func (p *Pairs) setIndexes(row []string) {
	p.indexes = make(map[int]string)
	color := "W"
	for i, v := range row {
		switch text := v; text {
		case "Имя":
			p.indexes[i] = color + "Name"
		case "Результат":
			p.indexes[i] = "Result"
			color = "B"
		case "Bo.":
			p.indexes[i] = "Table"
		}
	}
	log.Println(p.indexes)
}
