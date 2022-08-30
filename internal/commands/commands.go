package commands

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

const (
	LostSector string = "lostsector"
	Nightfall  string = "nightfall"
	Coolness   string = "cool"
)

var commands []SpeakerCommand = []SpeakerCommand{
	NewLostSectorCommand(),
	NewNightfallCommand(),
	NewCoolnessCommand(),
}

var commandMappings map[string]interface{} = map[string]interface{}{}

func AddCommands(s *discordgo.Session) {
	appId := os.Getenv("DISCORD_APP_ID")
	if appId == "" {
		log.Warnf("no bot token found, exiting..")
		return
	}
	for _, c := range commands {
		s.ApplicationCommandCreate(appId, "", c.GetCommand())
		log.Infof("registered %v command", c.GetName())
		commandMappings[c.GetName()] = c
	}
}

func AddHandler(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		data := i.ApplicationCommandData()
		cm, ok := commandMappings[data.Name]
		if !ok {
			log.Warnf("command %v doesn't have a mapping, not processing", data.Name)
			return
		}

		respondAck(s, i.Interaction)
		log.Infof("processing cmd %v in %v channel from %v guild...", data.Name, i.ChannelID, i.GuildID)

		sc := cm.(SpeakerCommand)
		res, err := sc.Handler(s, &data)
		if err != nil {
			log.Warnf("error processing command %v: %v", data.Name, err)
			respondMessage(s, i.Interaction, fmt.Sprintf("[error] %v: %v", sc.GetName(), err))
		}
		respondMessage(s, i.Interaction, res)
	})
}

func respondAck(s *discordgo.Session, i *discordgo.Interaction) {
	s.InteractionRespond(i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponsePong,
		Data: &discordgo.InteractionResponseData{},
	})
}

func respondMessage(s *discordgo.Session, i *discordgo.Interaction, msg string) {
	s.InteractionRespond(i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	})
}
