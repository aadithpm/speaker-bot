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

type LostSectorCommand struct {
	Name string
}

func NewLostSectorCommand() (l SpeakerCommand) {
	return LostSectorCommand{
		Name: LostSector,
	}
}

func (l LostSectorCommand) GetCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        LostSector,
		Type:        discordgo.ChatApplicationCommand,
		Description: "Lost Sector for the day",
	}
}

func (l LostSectorCommand) GetName() string {
	return l.Name
}

func (l LostSectorCommand) Handler(s *discordgo.Session, d *discordgo.ApplicationCommandInteractionData) (res string, emb *discordgo.MessageEmbed, err error) {
	log.Infof("got command %v from handler", d.Name)

	data := data.ReadRotationData("./data/lost_sectors.json")
	diff := utils.GetTimeDifferenceInDays(data.StartDate)

	if !data.RotationComplete && diff >= len(data.ContentRotation) {
		return "", nil, fmt.Errorf("today's Lost Sector doesn't have an entry")
	}

	lost_sector := data.ContentRotation[diff%len(data.ContentRotation)]
	gear := data.GearRotation[diff%len(data.GearRotation)]

	var shields []string
	for _, shield := range lost_sector.Shields {
		shields = append(shields, data.ElementList[shield])
	}

	var champions []string
	for _, champion := range lost_sector.Champions {
		champions = append(champions, data.ChampionList[champion])
	}

	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0x284030,
		Description: fmt.Sprintf("**%v**", data.LocationList[lost_sector.Location]),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Drops",
				Value:  data.GearList[gear],
				Inline: false,
			},
			{
				Name:   "Threat",
				Value:  data.ElementList[lost_sector.Threat],
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
		Title:     lost_sector.Name,
	}

	return "", embed, nil
}
