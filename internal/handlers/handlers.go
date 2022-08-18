package handlers

import (
	"github.com/aadithpm/speaker-bot/internal/handlers/member_add"
	"github.com/aadithpm/speaker-bot/internal/handlers/message_create"
	"github.com/bwmarrin/discordgo"
)

func AddHandlers(s *discordgo.Session) {
	message_create.AddHandlers(s)
	member_add.AddHandlers(s)
}
