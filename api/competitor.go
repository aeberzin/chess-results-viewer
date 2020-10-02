package api

import (
	"reflect"

	"github.com/PuerkitoBio/goquery"
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
	Indexes     indexes
	Competitors []Competitor
}

func NewCompetitors(node *goquery.Selection) *Competitors {
	var competitors Competitors
	competitors.Indexes = make(map[int]string)
	space := string([]byte{194, 160})
	node.Find("td").Each(func(i int, s *goquery.Selection) {
		switch text := s.Text(); text {
		case "Имя":
			competitors.Indexes[i] = "Name"
		case "Ст.ном":
			competitors.Indexes[i] = "Finish"
		case "Ст.ном.":
			competitors.Indexes[i] = "Start"
		case "Очки":
			competitors.Indexes[i] = "Result"
		case space + "Доп1" + space:
			competitors.Indexes[i] = "Rate1"
		case space + "Доп2" + space:
			competitors.Indexes[i] = "Rate2"
		case space + "Доп3" + space:
			competitors.Indexes[i] = "Rate3"
		}
	})
	return &competitors
}

func (competitors *Competitors) AddCompetitor(node *goquery.Selection) {
	var competitor Competitor
	node.Find("td").Each(func(i int, s *goquery.Selection) {
		if name, ok := competitors.Indexes[i]; ok {
			reflect.ValueOf(&competitor).Elem().FieldByName(name).SetString(s.Text())
		}
	})
	competitors.Competitors = append(competitors.Competitors, competitor)
}
