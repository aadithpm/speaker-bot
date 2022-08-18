package member_add

import (
	"fmt"

	"github.com/aadithpm/speaker-bot/internal/utils"
	"github.com/bwmarrin/discordgo"

	log "github.com/sirupsen/logrus"
)

// AddHandlers add handlers for GuildMemberAdd, fires when a new member joins the server
func AddHandlers(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
		log.Info("member joined")
		log.Infof("%v with ID %v joined the guild %v at %v", m.Member.User.Username, m.Member.User.ID, m.GuildID, m.JoinedAt)

		sendWelcomeMessage(s, m)
		setRoles(s, m)
	})
}

// Send welcome message to new member
func sendWelcomeMessage(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	c, err := s.GuildChannels(m.GuildID)
	if err != nil {
		log.Warnf("error getting guild channels for %v: %v", m.GuildID, err)
	}
	// Send a welcome message
	tc, err := utils.GetChannelByName(c, "welcome-wagon")
	if err != nil {
		log.Warnf("error getting channel name %v: %v", err)
	}

	msg := fmt.Sprintf("Hello <@%s>, welcome to the ARCH server! Please review <#786618669408976936> and register with <#785909633352466442>", m.User.ID)
	if !m.User.Bot {
		err = utils.SendMessageInChannel(s, tc, msg)
		if err != nil {
			log.Warnf("error sending welcome message: %v", err)
		}
	}
}

// Set roles for new member
func setRoles(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	err := s.GuildMemberRoleAdd(m.GuildID, m.Member.User.ID, "938261248829702195")
	if err != nil {
		log.Warnf("error adding role 938261248829702195 to member %v: %v", m.Member.User.Username, err)
	}
	err = s.GuildMemberRoleAdd(m.GuildID, m.Member.User.ID, "785952913301176360")
	if err != nil {
		log.Warnf("error adding role 785952913301176360 to member %v: %v", m.Member.User.Username, err)
	}
}
