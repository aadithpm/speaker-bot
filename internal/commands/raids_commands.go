package commands

import (
	"fmt"
	"time"

	"github.com/aadithpm/speaker-bot/internal/data"
	"github.com/aadithpm/speaker-bot/internal/utils"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

type RaidCommand struct {
	Name string
}

func NewRaidCommand() (c SpeakerCommand) {
	return RaidCommand{
		Name: Raid,
	}
}

func (c RaidCommand) GetCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        Raid,
		Type:        discordgo.ChatApplicationCommand,
		Description: "Featured Raid for the week",
	}
}

func (c RaidCommand) GetName() string {
	return c.Name
}

func (c RaidCommand) Handler(s *discordgo.Session, d *discordgo.ApplicationCommandInteractionData) (res string, emb *discordgo.MessageEmbed, err error) {
	log.Infof("got command %v from handler", d.Name)

	raids := data.ReadRotationData("./data/raids.json")
	current_week := utils.GetTimeDifferenceInWeeks(raids.StartDate)

	raid := raids.ContentRotation[current_week%len(raids.ContentRotation)]

	master := "✅"
	if !raid.MasterAvailable {
		master = "❌"
	}

	craftable := "✅"
	if !raid.Craftable {
		craftable = "❌"
	}

	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0x284030,
		Description: fmt.Sprintf("%v", raids.LocationList[raid.Location]),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Master",
				Value:  master,
				Inline: true,
			},
			{
				Name:   "Craftable",
				Value:  craftable,
				Inline: true,
			},
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Title:     raid.Name,
	}

	return "", embed, nil
}
