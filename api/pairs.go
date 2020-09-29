package api

import (
	"reflect"

	"github.com/PuerkitoBio/goquery"
)

type indexes map[int]string

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
	Indexes indexes
	Pairs   []Pair
}

func NewPairs(node *goquery.Selection) *Pairs {
	var pairs Pairs
	pairs.Indexes = make(map[int]string)
	color := "W"
	node.Find("td").Each(func(i int, s *goquery.Selection) {
		switch text := s.Text(); text {
		case "Имя":
			pairs.Indexes[i] = color + "Name"
		case "Результат":
			pairs.Indexes[i] = "Result"
			color = "B"
		case "Bo.":
			pairs.Indexes[i] = "Table"
		}
	})
	return &pairs
}

func (pairs *Pairs) AddPair(node *goquery.Selection) {
	var pair Pair
	node.Find("td").Each(func(i int, s *goquery.Selection) {
		if name, ok := pairs.Indexes[i]; ok {
			reflect.ValueOf(&pair).Elem().FieldByName(name).SetString(s.Text())
		}
	})
	pairs.Pairs = append(pairs.Pairs, pair)
}
