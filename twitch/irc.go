package twitch

import (
	"fmt"
	"os"
	"phillzbot/domain"
	"time"

	gt "github.com/Adeithe/go-twitch"
	"github.com/Adeithe/go-twitch/irc"
)

type TwitchIRC struct {
	See *irc.Client
	Say *irc.Conn
}

// type TwitchIRC interface {
// 	InitChat()
// 	Close()
// }

var (
	instance *TwitchIRC
)

func NewIRC() (*TwitchIRC, error) {
	if instance != nil {
		return instance, nil
	}

	username := os.Getenv("TMI_USER")
	if username == "" {
		return nil, fmt.Errorf("error: username not found")
	}

	oauth := os.Getenv("TMI_OAUTH")
	if oauth == "" {
		return nil, fmt.Errorf("error: oauth token not found")
	}
	see := gt.IRC()
	say := &irc.Conn{}

	err := say.SetLogin(username, oauth)
	if err != nil {
		return nil, err
	}

	instance = &TwitchIRC{
		See: see,
		Say: say,
	}

	return instance, nil
}

func onShardReconnect(shardID int) {
	fmt.Printf("Shard #%d reconnected\n", shardID)
}

func onShardLatencyUpdate(shardID int, latency time.Duration) {
	fmt.Printf("Shard #%d has %dms ping\n", shardID, latency.Milliseconds())
}

// func onShardMessage(shardID int, msg irc.ChatMessage) {
// 	if strings.HasPrefix(msg.Text, "!") {
// 		fmt.Printf("%s has a command!", msg.Text)
// 	}
// 	fmt.Printf("#%s %s: %s\n\n", msg.Channel, msg.Sender.DisplayName, msg.Text)
// }

func (i *TwitchIRC) InitChat(callback domain.TwitchHandleCallback) {
	// setup

	sc := make(chan os.Signal, 1)

	if err := i.Say.Connect(); err != nil {
		panic("failed to start writer")
	}

	fmt.Printf("we can talk back as %s", i.Say.Username)
	i.Say.Say("its_jay_phillz", fmt.Sprintf("%s is connected to chat! \n", i.Say.Username))

	i.See.OnShardMessage(callback)
	i.See.OnShardLatencyUpdate(onShardLatencyUpdate)
	i.See.OnShardReconnect(onShardReconnect)

	if err := i.See.Join("its_jay_phillz"); err != nil {
		panic(err)
	}

	fmt.Println("Connected to domain.TwitchIRC! \n")
	<-sc
}

func (i *TwitchIRC) Close() {
	i.See.Close()
	i.Say.Close()
}
