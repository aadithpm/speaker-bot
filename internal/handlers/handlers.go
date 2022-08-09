package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func AddHandlers(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		log.Info(m.Message.Content)
		// TODO do something with this later
	})

	s.AddHandler(func(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
		log.Info("member joined")
		log.Infof("%v with ID %v joined the guild %v at %v", m.Member.User.Username, m.Member.User.ID, m.GuildID, m.JoinedAt)

		c, err := s.GuildChannels(m.GuildID)
		if err != nil {
			log.Warnf("error getting guild channels for %v: %v", m.GuildID, err)
		}
		// Send a welcome message
		// TODO let user configure the channel and the message
		tc, err := GetChannelByName(c, "bot-tests")
		if err != nil {
			log.Warnf("error getting channel name %v")
		}

		msg := fmt.Sprintf("Hello <@%s>", m.User.ID)
		if !m.User.Bot {
			SendMessageInChannel(s, tc, msg)
		}
	})
}

// Gets channel by name from list of channels, sets err if not found
func GetChannelByName(c []*discordgo.Channel, n string) (channel *discordgo.Channel, err error) {
	for _, channel := range c {
		log.Infof("channel %v %v %v", channel.Name, channel.ID, channel.Type)
		if channel.Name == n {
			return channel, nil
		}
	}
	return nil, fmt.Errorf("channel name %v not found", n)
}

// Sends message in channel, sets err if failed
func SendMessageInChannel(s *discordgo.Session, c *discordgo.Channel, m string) {
	id := c.ID

	msg, err := s.ChannelMessageSend(id, m)
	if err != nil {
		log.Warn("error sending msg {%v} in channel %v %v: %v", msg.Content, c.Name, id, err)
		return
	}
	log.Infof("sent msg %v in channel %v %v", msg.Content, c.Name, id)
}
