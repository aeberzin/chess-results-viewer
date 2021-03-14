package api

import (
	"encoding/json"
	"net/http"

	socketio "github.com/googollee/go-socket.io"

	"github.com/gorilla/mux"
)

type Status int

const (
	NotStarted Status = iota
	InProgress
	Finished
	StartList
	Info
)

type API struct {
	cors string
	*mux.Router
	socket     *Socket
	Tournament *Tournament
	Status     Status
}

func NewAPI(subrouter *mux.Router, server *socketio.Server) *API {
	tournament := NewTournament("551049")
	socket := NewSocket(server, tournament)
	api := &API{
		"",
		subrouter,
		socket,
		tournament,
		NotStarted,
	}
	api.HandleFunc("/info", api.handleGetInfo()).Methods(http.MethodGet)
	api.HandleFunc("/info", api.handleSetInfo()).Methods(http.MethodPost)
	api.HandleFunc("/pairs", api.handleGetPairs()).Methods(http.MethodGet)
	api.HandleFunc("/competitors", api.handleGetCompetitors()).Methods(http.MethodGet)
	api.HandleFunc("/round", api.handleSetRound()).Methods(http.MethodPost)
	api.HandleFunc("/tournament", api.handleSetTournament()).Methods(http.MethodPost)
	api.HandleFunc("/result", api.handleSetResults()).Methods(http.MethodPost)
	api.HandleFunc("/timer", api.handleSetTimer()).Methods(http.MethodPost)
	api.HandleFunc("/timer", api.handleDeleteTimer()).Methods(http.MethodGet)
	return api
}

func (api *API) handleGetPairs() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		pairs := api.Tournament.GetPairs()
		res, _ := json.Marshal(pairs.Items)
		resp.Write(res)
	}
}

func (api *API) handleGetCompetitors() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		competitors := api.Tournament.GetCompetitors()
		res, _ := json.Marshal(competitors.Items)
		resp.Write(res)
	}
}

func (api *API) handleGetInfo() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {

		result := struct {
			Round      string
			Tournament string
			Status     Status
			Info       string
			Data       string
		}{
			api.Tournament.round,
			api.Tournament.id,
			api.Status,
			api.Tournament.info,
			api.Tournament.data,
		}
		res, _ := json.Marshal(result)
		resp.Write(res)
	}
}

func (api *API) handleSetInfo() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		var post struct {
			Info string
		}
		_ = json.NewDecoder(req.Body).Decode(&post)

		api.Tournament.SetInfo(post.Info)
		api.Status = Info
		api.socket.SendInfo()
		resp.WriteHeader(http.StatusOK)
	}
}

func (api *API) handleSetData() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		var post struct {
			Info string
		}
		_ = json.NewDecoder(req.Body).Decode(&post)

		api.Tournament.SetInfo(post.Info)
		api.Status = Info
		api.socket.SendInfo()
		resp.WriteHeader(http.StatusOK)
	}
}

func (api *API) handleSetRound() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		var post struct {
			Round string
		}
		_ = json.NewDecoder(req.Body).Decode(&post)

		api.Tournament.SetRound(post.Round)
		api.Status = InProgress
		api.socket.SendNewRound()
		resp.WriteHeader(http.StatusOK)
	}
}

func (api *API) handleSetTournament() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		var post struct {
			Tournament string
		}
		_ = json.NewDecoder(req.Body).Decode(&post)
		api.Tournament.SetID(post.Tournament)
		api.Status = StartList
		api.socket.SendStartList()
		resp.WriteHeader(http.StatusOK)
	}
}

func (api *API) handleSetResults() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		var post struct {
			Round string
		}
		_ = json.NewDecoder(req.Body).Decode(&post)
		api.Tournament.SetRound(post.Round)
		api.Status = Finished
		api.socket.SendResults()
		resp.WriteHeader(http.StatusOK)
	}
}

func (api *API) handleDeleteTimer() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		api.socket.SendRemoveTimer()
		resp.WriteHeader(http.StatusOK)
	}
}

func (api *API) handleSetTimer() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		var post struct {
			Time string
			Text string
		}
		_ = json.NewDecoder(req.Body).Decode(&post)
		api.socket.SendSetTimer(post.Text, post.Time)
		resp.WriteHeader(http.StatusOK)
	}
}
