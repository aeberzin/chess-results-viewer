package api

import (
	"strconv"

	"github.com/aeberzin/chess-results-viewer/parser"
)

type Tournament struct {
	id    string
	round string
}

func NewTournament(id string) *Tournament {
	return &Tournament{id, "2"}
}

func (t *Tournament) SetID(id string) {
	t.id = id
}

func (t *Tournament) SetRound(round string) {
	t.round = round
}

func (t *Tournament) GetPairs() *Pairs {
	return t.getPairs(t.round)
}

func (t *Tournament) getPairs(round string) *Pairs {
	rows := parser.GetTable(t.id, parser.Pairs, round)
	pairs := NewPairs(rows)
	return pairs
}

func (t *Tournament) GetCompetitors() *Competitors {
	round, _ := strconv.Atoi(t.round)
	if round == 1 {
		var competitors Competitors
		return &competitors
	}
	return t.getCompetitors(strconv.Itoa(round - 1))
}

func (t *Tournament) getCompetitors(round string) *Competitors {
	rows := parser.GetTable(t.id, parser.Competitors, round)
	competitors := NewCompetitors(rows)
	return competitors
}

func (t *Tournament) GetResults() *Competitors {
	return t.getCompetitors(t.round)
}

func (t *Tournament) GetStartList() *Players {
	rows := parser.GetTable(t.id, parser.StartList)
	competitors := NewPlayers(rows)
	return competitors
}
