package command

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"phillzbot/domain"
	"phillzbot/twitch"
	"phillzbot/utils"
	"strconv"
	"strings"
	"time"

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

		cmdValue = strings.Replace(cmdValue, "$viewer.last_streamed", lastStreamed.Data[0].GameTopic, 1)
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

	if strings.Contains(cmdResponse, "$followage") {
		layout := "2006-01-02T15:04:05Z"

		var followResponse domain.HelixUserFollowsResponse

		followageRequestParams := fmt.Sprintf("/users/follows?from_id=%s&to_id=%s", strconv.FormatInt(msg.Sender.ID, 10), "193648255")
		fmt.Println(followageRequestParams)
		resp, err3 := s.api.Helix.Request("GET", followageRequestParams, nil)

		fmt.Println(err3)

		json.Unmarshal(resp.Body, &followResponse)
		// now := time.Now()
		if len(followResponse.Data) != 0 {
			t, err := time.Parse(layout, followResponse.Data[0].FollowedAt)
			if err != nil {
				fmt.Println(err)
			}

			// duration := time.Since(t)
			y, m, d, h, mi, s := utils.DateDiff(t, time.Now())
			cmdString := fmt.Sprintf("%d years, %d months, %d days, %d hours, %d minutes and %d seconds... and counting", y, m, d, h, mi, s)

			cmdValue = strings.Replace(cmdValue, "$followage", cmdString, 1)
		} else {
			cmdValue = strings.Replace(cmdValue, "$followage", "...oops, I couldn't tell. please try again later", 1)
		}

	}

	if strings.Contains(cmdResponse, "$commands") {
		commands := s.GetTwitchCommandMap()
		keys := ""

		for k := range commands {
			keys = keys + fmt.Sprintf("!%s ", k)
		}
		cmdValue = strings.Replace(cmdValue, "$commands", keys, 1)
	}

	return cmdValue, nil
}
