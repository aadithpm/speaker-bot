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

	if !data.RotationComplete && diff >= len(data.ContentRotation) {
		return "", fmt.Errorf("today's Lost Sector doesn't have an entry")
	}

	lost_sector := data.ContentRotation[diff%len(data.ContentRotation)]
	gear := data.GearRotation[diff%len(data.GearRotation)]

	msg := fmt.Sprintf("**WARNING: Nightfall data will be inaccurate from the release of Lightfall until the end of March.** Lost Sector for today is **%v** in %v, dropping **%v**.", lost_sector.Name, data.LocationList[lost_sector.Location], data.GearList[gear])

	return msg, nil
}
