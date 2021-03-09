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
		Competitors *Competitors
		Round       string
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

func (s *Socket) SendRemoveTimer() {
	s.server.BroadcastToRoom("all", "RemoveTimer", 1)
}

func (s *Socket) SendSetTimer(Text string, Timer string) {
	s.server.BroadcastToRoom("all", "SetTimer", Text, Timer)
}

func (s *Socket) SendInfo() {
	s.server.BroadcastToRoom("all", "SetInfo", s.tournament.info)
}
