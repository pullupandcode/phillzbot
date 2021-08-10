package command

import (
	"context"
	"fmt"
	"phillzbot/domain"
	"strings"

	"github.com/Adeithe/go-twitch/irc"
)

type TwitchIRCCommandHandler struct {
	TwitchCommandService domain.TwitchCommandService
}

func HandleTwitchCommand(msg irc.ChatMessage, tcs domain.TwitchCommandService, chat *irc.Conn) {
	response, _ := tcs.FormatCommandMessage(msg)
	chat.Say(msg.Channel, response)
}

func getTwitchCommandMap(tcs domain.TwitchCommandService) map[string]string {
	commands := make(map[string]string)
	commandData, _ := tcs.Fetch(context.Background())

	for _, cmd := range commandData {
		commands[cmd.Name] = cmd.Value
	}

	return commands
}

func formatTwitchResponse(msg irc.ChatMessage, tcs domain.TwitchCommandService) string {
	// commands := make(map[string]string)
	variables := make(map[string]string)
	var data string

	commands := getTwitchCommandMap(tcs)
	phrase := strings.Fields(msg.Text)

	if len(phrase) > 1 {
		variables["$viewer"] = phrase[1]
	}

	command := strings.Replace(phrase[0], "!", "", 1)
	variables["$viewer"] = ""

	for k, v := range variables {
		if strings.Contains(commands[command], k) {
			fmt.Printf("%s, %s, %s \n", commands[command], k, v)
			data = strings.Replace(commands[command], k, v, 1)
		}
	}

	fmt.Print(data)

	if data != "" {
		return data
	} else {
		return msg.Text
	}
}
