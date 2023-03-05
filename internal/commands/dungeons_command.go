package commands

import (
	"fmt"
	"time"

	"github.com/aadithpm/speaker-bot/internal/data"
	"github.com/aadithpm/speaker-bot/internal/utils"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

type DungeonCommand struct {
	Name string
}

func NewDungeonCommand() (c SpeakerCommand) {
	return DungeonCommand{
		Name: Dungeon,
	}
}

func (c DungeonCommand) GetCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        Dungeon,
		Type:        discordgo.ChatApplicationCommand,
		Description: "Featured Dungeon for the week",
	}
}

func (c DungeonCommand) GetName() string {
	return c.Name
}

func (c DungeonCommand) Handler(s *discordgo.Session, d *discordgo.ApplicationCommandInteractionData) (res string, emb *discordgo.MessageEmbed, err error) {
	log.Infof("got command %v from handler", d.Name)

	dungeons := data.ReadRotationData("./data/dungeons.json")
	current_week := utils.GetTimeDifferenceInWeeks(dungeons.StartDate)

	dungeon := dungeons.ContentRotation[current_week%len(dungeons.ContentRotation)]

	master := "✅"
	if !dungeon.MasterAvailable {
		master = "❌"
	}

	craftable := "✅"
	if !dungeon.Craftable {
		craftable = "❌"
	}

	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0x284030,
		Description: fmt.Sprintf("%v", dungeons.LocationList[dungeon.Location]),
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
		Title:     dungeon.Name,
	}

	return "", embed, nil
}
