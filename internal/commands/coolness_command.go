package commands

import (
	"github.com/bwmarrin/discordgo"
)

type CoolnessCommand struct {
	Name string
}

func NewCoolnessCommand() (l SpeakerCommand) {
	return CoolnessCommand{
		Name: Coolness,
	}
}

func (c CoolnessCommand) GetCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        Coolness,
		Type:        discordgo.ChatApplicationCommand,
		Description: "Clan coolness chart",
	}
}

func (c CoolnessCommand) GetName() string {
	return c.Name
}

func (c CoolnessCommand) Handler(s *discordgo.Session, d *discordgo.ApplicationCommandInteractionData) (res string, err error) {

	msg := "https://docs.google.com/spreadsheets/d/1MX9pq6kedcJ_1tRGGeWIsp4GJgq2AnJLn8DcpgX9w48"

	return msg, nil
}
