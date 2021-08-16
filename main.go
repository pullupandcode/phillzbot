package main

import (
	"fmt"
	"net/http"
	"os"

	commandHandler "phillzbot/command/adapter/http"
	commandRepo "phillzbot/command/repo/mongo"
	commandService "phillzbot/command/service"
	_ "phillzbot/config"
	"phillzbot/infrastructure/repository"
	"phillzbot/twitch"

	"github.com/Adeithe/go-twitch/irc"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func LoadEnvVars(keys []string) map[string]string {
	var envVals = make(map[string]string)
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	for _, key := range keys {
		fmt.Println(key)
		envVals[key] = os.Getenv(key)
	}

	return envVals
}

func main() {
	var log log.Logger
	var err error
	_ = godotenv.Load()

	r := mux.NewRouter().PathPrefix("/api/v1").Subrouter()

	dbConn := &repository.MongoDB{}
	dbConn, err = repository.GetConnection()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d \n", dbConn.Client.NumberSessionsInProgress())

	twitchRepo := commandRepo.NewMongoCommandRepo(dbConn.Database)
	twitchCommandService := commandService.NewTwitchCommandService(twitchRepo)
	commandHandler.AddTwitchHandler(r, twitchCommandService)

	twitchChat, err := twitch.NewIRC()
	if err != nil {
		log.Fatal(err)
	}
	go (func() {
		twitchChat.InitChat(func(shardID int, msg irc.ChatMessage) {
			err := twitchChat.Say.Say("BNiCe0rELsE", "jay phillz told me its a party in here LUL ")
			if err != nil {
				fmt.Println(err)
			}
		})
	})()

	// discord
	// commandRepo := mongo.NewMongoCommandRepo(dbConn.Database)
	// twitchCommandService := commandService.NewTwitchCommandService(commandRepo)
	// commandHandler.AddTwitchHandler(r, twitchCommandService)

	// loggedRouter := handlers.LoggingHandler(log.Writer(), r)

	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tmpl, _ := route.GetPathTemplate()
		m, _ := route.GetMethods()
		if len(m) != 0 {
			fmt.Printf("%30s:     %-10s    %s\n", route.GetName(), m[0], tmpl)
		}
		return nil
	})

	fmt.Printf("PORT: %s \n\n", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r))

}
