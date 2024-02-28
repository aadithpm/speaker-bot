package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aadithpm/speaker-bot/internal/utils"
	"github.com/bwmarrin/discordgo"

	log "github.com/sirupsen/logrus"
)

type WorkersAiHttpResponse []struct {
	Inputs struct {
		Messages []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"messages"`
	} `json:"inputs"`
	Response struct {
		Response string `json:"response"`
	} `json:"response"`
}

type AiCommand struct {
	Name string
}

func NewAiCommand() (l SpeakerCommand) {
	return AiCommand{
		Name: Ai,
	}
}

func (a AiCommand) GetCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        Ai,
		Type:        discordgo.ChatApplicationCommand,
		Description: "Give a prompt to the AI",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "prompt",
				Description: "Prompt for the AI",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
		},
	}
}

func (a AiCommand) GetName() string {
	return a.Name
}

func (a AiCommand) Handler(s *discordgo.Session, d *discordgo.ApplicationCommandInteractionData) (res string, emb *discordgo.MessageEmbed, err error) {
	m := utils.GetOptionsMap(d.Options)
	p := m["prompt"].StringValue()

	jsonStr := []byte(fmt.Sprintf(`{"prompt": "%s"}`, p))
	log.Infof("making POST request to worker..")
	resp, err := http.Post("https://malahayati-discord-ai.aadith-pm.workers.dev", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Errorf("error making request to workers: %v", err)
		return "", nil, err
	}
	defer resp.Body.Close()

	var workersResponse WorkersAiHttpResponse
	body, err := ioutil.ReadAll(resp.Body)
	bodyString := string(body)
	log.Info(bodyString)
	if err != nil {
		log.Errorf("error reading body: %v", err)
		return "", nil, err
	}
	err = json.Unmarshal(body, &workersResponse)
	if err != nil {
		log.Errorf("error unmarshaling: %v", err)
		return "", nil, err
	}
	if len(workersResponse[0].Response.Response) > 0 {
		return workersResponse[0].Response.Response, nil, nil
	}
	return "", nil, fmt.Errorf("workers response was empty")
}
