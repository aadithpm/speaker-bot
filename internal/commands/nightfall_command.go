package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/aadithpm/speaker-bot/internal/data"
	"github.com/aadithpm/speaker-bot/internal/utils"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

type NightfallCommand struct {
	Name string
}

func NewNightfallCommand() (l SpeakerCommand) {
	return NightfallCommand{
		Name: Nightfall,
	}
}

func (n NightfallCommand) GetCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        Nightfall,
		Type:        discordgo.ChatApplicationCommand,
		Description: "Nightfall for the week",
	}
}

func (n NightfallCommand) GetName() string {
	return n.Name
}

func (n NightfallCommand) Handler(s *discordgo.Session, d *discordgo.ApplicationCommandInteractionData) (res string, emb *discordgo.MessageEmbed, err error) {
	log.Infof("got command %v from handler", d.Name)

	data := data.ReadRotationData("./data/nightfalls.json")
	diff := utils.GetTimeDifferenceInWeeks(data.StartDate)

	if !data.RotationComplete && diff >= len(data.ContentRotation) {
		return "", nil, fmt.Errorf("this week's Nightfall doesn't have an entry")
	}

	nightfall := data.ContentRotation[diff%len(data.ContentRotation)]
	gear := data.GearRotation[diff%len(data.GearRotation)]

	var shields []string
	for _, shield := range nightfall.Shields {
		shields = append(shields, data.ElementList[shield])
	}

	var champions []string
	for _, champion := range nightfall.Champions {
		champions = append(champions, data.ChampionList[champion])
	}

	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0x284030,
		Description: fmt.Sprintf("**%v**", data.LocationList[nightfall.Location]),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Drops",
				Value:  fmt.Sprintf("%v, Exotics, Enhancement Prisms, Ascendant Shards", data.GearList[gear]),
				Inline: false,
			},
			{
				Name:   "Threat",
				Value:  data.ElementList[nightfall.Threat],
				Inline: true,
			},
			{
				Name:   "Shields",
				Value:  strings.Join(shields, ", "),
				Inline: true,
			},
			{
				Name:   "Champions",
				Value:  strings.Join(champions, ", "),
				Inline: true,
			},
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Title:     nightfall.Name,
	}

	return "", embed, nil
}
