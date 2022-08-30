package commands

import (
	"fmt"

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

func (n NightfallCommand) Handler(s *discordgo.Session, d *discordgo.ApplicationCommandInteractionData) (res string, err error) {
	log.Infof("got command %v from handler", d.Name)

	data := data.ReadRotationData("./data/nightfalls.json")
	diff := utils.GetTimeDifferenceInDays(data.StartDate) / 7
	nightfall := data.Rotation[diff]

	msg := fmt.Sprintf("Nightfall for this week is **%v** in %v, dropping **%v**.", nightfall.Name, data.LocationList[nightfall.Location], data.GearList[nightfall.Gear])

	return msg, nil
}
