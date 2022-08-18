package utils

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// SendMessageInChannel Sends message m in channel c, sets err if failed
func SendMessageInChannel(s *discordgo.Session, c *discordgo.Channel, m string) (err error) {
	id := c.ID

	msg, err := s.ChannelMessageSend(id, m)
	if err != nil {
		log.Warn("error sending msg {%v} in channel %v %v: %v", msg.Content, c.Name, id, err)
		return err
	}
	log.Infof("sent msg %v in channel %v %v", msg.Content, c.Name, id)
	return nil
}

// GetChannelByName Gets channel by name from list of channels, sets err if not found
func GetChannelByName(c []*discordgo.Channel, n string) (channel *discordgo.Channel, err error) {
	for _, channel := range c {
		log.Infof("channel %v %v %v", channel.Name, channel.ID, channel.Type)
		if channel.Name == n {
			return channel, nil
		}
	}
	return nil, fmt.Errorf("channel name %v not found", n)
}
