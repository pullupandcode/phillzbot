package command

import (
	"fmt"
	"phillzbot/domain"

	"github.com/Adeithe/go-twitch/irc"
)

type TwitchIRCCommandHandler struct {
	TwitchCommandService domain.TwitchCommandService
}

func HandleTwitchCommand(msg irc.ChatMessage, chat *irc.Conn) {
	fmt.Print("handled homie!")
	chat.Say(msg.Channel, "handled!")
}
