package twitch

import (
	"fmt"
	"os"

	tc "github.com/gempir/go-twitch-irc"
)

type Chat struct{}

func (t *Chat) Init() {

	client := tc.NewClient(os.Getenv("TMI_USER"), os.Getenv("TMI_ACCESS_TOKEN"))

	client.OnNewMessage(func(channel string, user tc.User, message tc.Message) {
		fmt.Println(message.Text)
	})

	client.Join(os.Getenv("TMI_CHANNEL"))

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
