package http

import (
	"net/http"
	"phillzbot/domain"

	"github.com/gorilla/mux"
)

type DiscordCommandHandler struct {
	DiscordService domain.DiscordService
}

func AddDiscordHandler(r *mux.Router, ds domain.DiscordService) {
	handler := &DiscordCommandHandler{
		DiscordService: ds,
	}

	r.HandleFunc("/discord/commands", handler.HandleCreate).Methods("POST").Name("CreateDiscordCommand")
	r.HandleFunc("/discord/commands", handler.HandleFetch).Methods("GET").Name("FetchAllDiscordCommands")
	r.HandleFunc("/discord/commands/:id", handler.HandleFetchById).Methods("GET").Name("FetchDiscordCommandById")
	r.HandleFunc("/discord/commands/:name", handler.HandleFetchByName).Methods("GET").Name("FetchDiscordCommandByName")
	r.HandleFunc("/discord/commands/:id", handler.HandleUpdate).Methods("POST").Name("UpdateDiscordCommand")
	r.HandleFunc("/discord/commands/:id", handler.HandleDelete).Methods("POST").Name("DeleteDiscordCommand")

}

func (tch *DiscordCommandHandler) HandleCreate(w http.ResponseWriter, r *http.Request)      {}
func (tch *DiscordCommandHandler) HandleFetch(w http.ResponseWriter, r *http.Request)       {}
func (tch *DiscordCommandHandler) HandleFetchById(w http.ResponseWriter, r *http.Request)   {}
func (tch *DiscordCommandHandler) HandleFetchByName(w http.ResponseWriter, r *http.Request) {}
func (tch *DiscordCommandHandler) HandleUpdate(w http.ResponseWriter, r *http.Request)      {}
func (tch *DiscordCommandHandler) HandleDelete(w http.ResponseWriter, r *http.Request)      {}
