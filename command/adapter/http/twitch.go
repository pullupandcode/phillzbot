package http

import (
	"net/http"
	"phillzbot/domain"

	"github.com/gorilla/mux"
)

type TwitchCommandHandler struct {
	TwitchCommandService domain.TwitchCommandService
}

func AddTwitchHandler(r *mux.Router, tcs domain.TwitchCommandService) {
	handler := &TwitchCommandHandler{
		TwitchCommandService: tcs,
	}

	r.HandleFunc("/twitch/commands", handler.HandleCreate).Methods("POST").Name("CreateTwitchCommand")
	r.HandleFunc("/twitch/commands", handler.HandleFetch).Methods("GET").Name("FetchAllTwitchCommands")
	r.HandleFunc("/twitch/commands/:id", handler.HandleFetchById).Methods("GET").Name("FetchTwitchCommandById")
	r.HandleFunc("/twitch/commands/:name", handler.HandleFetchByName).Methods("GET").Name("FetchTwitchCommandByName")
	r.HandleFunc("/twitch/commands/:id", handler.HandleUpdate).Methods("POST").Name("UpdateTwitchCommand")
	r.HandleFunc("/twitch/commands/:id", handler.HandleDelete).Methods("POST").Name("DeleteTwitchCommand")
}

func (tch *TwitchCommandHandler) HandleCreate(w http.ResponseWriter, r *http.Request)      {}
func (tch *TwitchCommandHandler) HandleFetch(w http.ResponseWriter, r *http.Request)       {}
func (tch *TwitchCommandHandler) HandleFetchById(w http.ResponseWriter, r *http.Request)   {}
func (tch *TwitchCommandHandler) HandleFetchByName(w http.ResponseWriter, r *http.Request) {}
func (tch *TwitchCommandHandler) HandleUpdate(w http.ResponseWriter, r *http.Request)      {}
func (tch *TwitchCommandHandler) HandleDelete(w http.ResponseWriter, r *http.Request)      {}
