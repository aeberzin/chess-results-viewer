package api

import socketio "github.com/googollee/go-socket.io"

type Socket struct {
	server     *socketio.Server
	tournament *Tournament
}

func NewSocket(server *socketio.Server, tournament *Tournament) *Socket {
	return &Socket{
		server,
		tournament,
	}
}

func (s *Socket) SendNewRound() {
	pairs := s.tournament.GetPairs()
	competitors := s.tournament.GetCompetitors()
	broadcastInfo := struct {
		Pairs       *Pairs
		Competitors *Competitors
		Round       string
	}{
		pairs,
		competitors,
		s.tournament.round,
	}
	s.server.BroadcastToRoom("all", "SetRound", broadcastInfo)
}

func (s *Socket) SendResults() {
	competitors := s.tournament.GetResults()
	broadcastInfo := struct {
		competitors *Competitors
		round       string
	}{
		competitors,
		s.tournament.round,
	}
	s.server.BroadcastToRoom("all", "SetResults", broadcastInfo)
}

func (s *Socket) SendStartList() {
	players := s.tournament.GetStartList()
	s.server.BroadcastToRoom("all", "SetStartList", players)
}
