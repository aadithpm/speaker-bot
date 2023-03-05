package commands

import (
	"fmt"
	"time"

	"github.com/aadithpm/speaker-bot/internal/data"
	"github.com/aadithpm/speaker-bot/internal/utils"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

type LostSectorListCommand struct {
	Name string
}

func NewLostSectorListCommand() (l SpeakerCommand) {
	return LostSectorListCommand{
		Name: LostSectorList,
	}
}

func (l LostSectorListCommand) GetCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        LostSectorList,
		Type:        discordgo.ChatApplicationCommand,
		Description: "Lost Sector list for the week",
	}
}

func (l LostSectorListCommand) GetName() string {
	return l.Name
}

func (l LostSectorListCommand) Handler(s *discordgo.Session, d *discordgo.ApplicationCommandInteractionData) (res string, emb *discordgo.MessageEmbed, err error) {
	log.Infof("got command %v from handler", d.Name)

	dates := []time.Time{
		time.Now(),
		time.Now().AddDate(0, 0, 1),
		time.Now().AddDate(0, 0, 2),
		time.Now().AddDate(0, 0, 3),
		time.Now().AddDate(0, 0, 4),
		time.Now().AddDate(0, 0, 5),
	}

	var resp string

	for _, date := range dates {
		data := data.ReadRotationData("./data/lost_sectors.json")
		diff := utils.GetTimeDifferenceInDaysFrom(data.StartDate, date)

		if !data.RotationComplete && diff >= len(data.ContentRotation) {
			return "", nil, fmt.Errorf("lost sector list is not complete")
		}

		lost_sector := data.ContentRotation[diff%len(data.ContentRotation)]
		gear := data.GearRotation[diff%len(data.GearRotation)]

		msg := fmt.Sprintf("__%v:__ **%v** in %v, dropping **%v**.\n", date.Weekday(), lost_sector.Name, data.LocationList[lost_sector.Location], data.GearList[gear])
		resp = resp + msg
	}
	return resp, nil, nil
}
