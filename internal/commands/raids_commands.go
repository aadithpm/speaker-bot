package commands

import (
	"fmt"

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

func (c RaidCommand) Handler(s *discordgo.Session, d *discordgo.ApplicationCommandInteractionData) (res string, err error) {
	log.Infof("got command %v from handler", d.Name)

	raids := data.ReadRotationData("./data/raids.json")
	current_week := utils.GetTimeDifferenceInWeeks(raids.StartDate)
	raid := raids.ContentRotation[current_week%len(raids.ContentRotation)]

	str := "Featured Raid for this week is **%v** at %v"
	if raid.MasterAvailable {
		str += " with Master difficulty"
	}
	if raid.Craftable {
		str += " and craftable weapons"
	}
	str += "."

	msg := fmt.Sprintf(
		str,
		raid.Name,
		raids.LocationList[raid.Location],
	)

	return msg, nil
}
