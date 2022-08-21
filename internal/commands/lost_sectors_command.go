package commands

import (
	"github.com/aadithpm/speaker-bot/internal/data"
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

func (l LostSectorCommand) Handler(s *discordgo.Session, d *discordgo.ApplicationCommandInteractionData) (err error) {
	log.Infof("got command %v from handler", d.Name)
	data.ReadRotationData("./data/lost_sectors.json")
	return nil
}
