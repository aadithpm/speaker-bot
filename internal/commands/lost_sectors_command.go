package commands

import (
	"fmt"

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

func (l LostSectorCommand) Handler(s *discordgo.Session, d *discordgo.ApplicationCommandInteractionData) (res string, err error) {
	log.Infof("got command %v from handler", d.Name)

	data := data.ReadRotationData("./data/lost_sectors.json")
	diff := utils.GetTimeDifferenceInDays(data.StartDate)
	lost_sector := data.Rotation[diff]

	msg := fmt.Sprintf("Lost Sector for today is **%v** in %v, dropping **%v**. *Warning: Lost Sectors might be out of date until a week into the season*", lost_sector.Name, data.LocationList[lost_sector.Location], data.GearList[lost_sector.Gear])

	return msg, nil
}
