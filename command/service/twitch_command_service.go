package command

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"phillzbot/domain"
	"phillzbot/twitch"
	"strings"

	"github.com/Adeithe/go-twitch/api/helix"
	"github.com/Adeithe/go-twitch/irc"
	log "github.com/sirupsen/logrus"
)

type twitchCommandService struct {
	repo   domain.TwitchCommandRepo
	api    *domain.TwitchAPI
	logger *log.Logger
}

var varMap map[string]string

func NewTwitchCommandService(tcr domain.TwitchCommandRepo) domain.TwitchCommandService {
	httpclient, _ := twitch.NewAPIClient()
	logger := log.New()
	logger.SetLevel(log.DebugLevel)

	return &twitchCommandService{
		repo:   tcr,
		api:    httpclient,
		logger: logger,
	}
}

func (s *twitchCommandService) Fetch(ctx context.Context) (data []domain.TwitchCommand, err error) {
	result, err := s.repo.Fetch(ctx)
	fmt.Println(result)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (s *twitchCommandService) FetchById(ctx context.Context, id string) (data domain.TwitchCommand, err error) {
	result, err := s.repo.FetchById(ctx, id)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (s *twitchCommandService) FetchByName(ctx context.Context, name string) (data domain.TwitchCommand, err error) {
	return data, nil
}
func (s *twitchCommandService) Update(ctx context.Context, tc *domain.TwitchCommand) error {
	return nil
}
func (s *twitchCommandService) Create(ctx context.Context, tc *domain.TwitchCommand) error {
	err := s.repo.Create(ctx, tc)
	if err != nil {
		return err
	}
	return nil
}
func (s *twitchCommandService) Delete(ctx context.Context, id string) error {
	return nil
}

func (s *twitchCommandService) GetCommandVariables() map[string]string {
	if len(varMap) > 0 {
		return varMap
	}

	var variables domain.TwitchCommandVariableList

	variablesFile, err := os.Open("variables.json")
	if err != nil {
		log.Print(err)
	}

	rawData, _ := ioutil.ReadAll(variablesFile)
	json.Unmarshal(rawData, &variables)

	dataMap := map[string]string{}
	for _, vari := range variables.Variables {
		dataMap[vari.Name] = vari.Description
	}

	varMap = dataMap
	return varMap
}

func (s *twitchCommandService) GetTwitchCommandMap() map[string]string {
	commands := make(map[string]string)
	commandData, _ := s.Fetch(context.Background())

	for _, cmd := range commandData {
		commands[cmd.Name] = cmd.Value
	}

	return commands
}

func (s *twitchCommandService) FormatCommandMessage(msg irc.ChatMessage) (string, error) {
	var cmdValue string
	cmdMap := s.GetTwitchCommandMap()
	cmdArgs := strings.Fields(msg.Text)
	cmdName := strings.Replace(cmdArgs[0], "!", "", 1)

	cmdResponse := cmdMap[cmdName]
	cmdValue = cmdResponse

	s.logger.Infof("before: %s\n", cmdValue)

	if strings.Contains(cmdResponse, "$viewer.last_streamed") {

		var lastStreamed domain.HelixChannelResponse
		reqUsers := []string{}

		reqUser := strings.Replace(cmdArgs[1], "@", "", 1)
		reqUsers = append(reqUsers, reqUser)
		users, err := s.api.Helix.GetUsers(helix.UserOpts{
			Logins: reqUsers,
		})

		if err != nil {
			log.Print(err)
		}

		lastStreamRequestParams := fmt.Sprintf("/channels?broadcaster_id=%s", users.Data[0].ID)
		resp, _ := s.api.Helix.Request("GET", lastStreamRequestParams, nil)

		json.Unmarshal(resp.Body, &lastStreamed)

		s.logger.Info(lastStreamed)
		cmdValue = strings.Replace(cmdValue, "$viewer.last_streamed", lastStreamed.Data[0].GameTopic, 1)
		s.logger.Debugf("$viewer: %s\n", cmdValue)
	}

	if strings.Contains(cmdResponse, "$shoutname") {
		cmdValue = strings.Replace(cmdValue, "$shoutname", strings.Replace(cmdArgs[1], "@", "", 1), 1)
	}

	if strings.Contains(cmdValue, "$viewer") {
		cmdValue = strings.Replace(cmdValue, "$viewer", cmdArgs[1], 1)
	}

	if strings.Contains(cmdResponse, "$sender") {
		cmdValue = strings.Replace(cmdResponse, "$sender", msg.Sender.Username, 1)
	}

	if strings.Contains(cmdResponse, "$commands") {
		commands := s.GetTwitchCommandMap()
		keys := ""
		for k := range commands {
			keys = keys + fmt.Sprintf("!%s ", k)
		}
		cmdValue = strings.Replace(cmdResponse, "$commands", keys, 1)
	}

	return cmdValue, nil
}
