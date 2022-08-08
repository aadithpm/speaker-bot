package handlers

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func AddHandlers(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		log.Info(m.Message.Content)
	})

	s.AddHandler(func(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
		log.Info("Member joined")
		log.Infof("%s with ID %s joined the guild %s at %s", m.Member.User.Username, m.Member.User.ID, m.GuildID, m.JoinedAt)

		channels, err := s.GuildChannels(m.GuildID)
		if err != nil {
			log.Warnf("Error getting guild channels for %s: %s", m.GuildID, err)
		}
		log.Infof("%s", channels)

	})
}
