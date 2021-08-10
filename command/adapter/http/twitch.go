package http

import (
	"context"
	"encoding/json"
	"net/http"
	"phillzbot/domain"
	"phillzbot/utils"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	r.HandleFunc("/twitch/commands/{id:[a-zA-Z0-9]+}", handler.HandleFetchById).Methods("GET").Name("FetchTwitchCommandById")
	r.HandleFunc("/twitch/commands/:id", handler.HandleUpdate).Methods("POST").Name("UpdateTwitchCommand")
	r.HandleFunc("/twitch/commands/:id", handler.HandleDelete).Methods("POST").Name("DeleteTwitchCommand")
}

func (tch *TwitchCommandHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var commandItem *domain.TwitchCommand
	_ = json.NewDecoder(r.Body).Decode(&commandItem)

	commandItem.ID = primitive.NewObjectID()
	err := tch.TwitchCommandService.Create(context.Background(), commandItem)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	utils.RespondWithJSON(w, http.StatusCreated, commandItem)
}
func (tch *TwitchCommandHandler) HandleFetch(w http.ResponseWriter, r *http.Request) {
	result, err := tch.TwitchCommandService.Fetch(context.Background())
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	utils.RespondWithJSON(w, http.StatusOK, result)
}
func (tch *TwitchCommandHandler) HandleFetchById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commandId := vars["id"]

	command, err := tch.TwitchCommandService.FetchById(context.Background(), commandId)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	utils.RespondWithJSON(w, http.StatusOK, command)
}
func (tch *TwitchCommandHandler) HandleFetchByName(w http.ResponseWriter, r *http.Request) {}
func (tch *TwitchCommandHandler) HandleUpdate(w http.ResponseWriter, r *http.Request)      {}
func (tch *TwitchCommandHandler) HandleDelete(w http.ResponseWriter, r *http.Request)      {}
