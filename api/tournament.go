package api

type Tournament struct {
	id    string
	round string
}

func NewTournament(id string) *Tournament {
	return &Tournament{id, "1"}
}

func (tournament *Tournament) SetID(id string) {
	tournament.id = id
}

func (tournament *Tournament) SetRound(round string) {
	tournament.round = round
}
