package commands

import "github.com/bwmarrin/discordgo"

type SpeakerCommand interface {
	GetName() string
	GetCommand() *discordgo.ApplicationCommand
	Handler(s *discordgo.Session, d *discordgo.ApplicationCommandInteractionData) (string, error)
}
