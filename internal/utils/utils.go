package utils

import (
	"fmt"
	"math"
	"time"

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

// GetTimeDifferenceInDays Gets time difference in days between time arg and current time
func GetTimeDifferenceInDays(t time.Time) (diff int) {
	currentTime := time.Now()
	dur := currentTime.Sub(t)
	return int(math.Floor(dur.Hours() / 24))
}

// GetTimeDifferenceInDays Gets time difference in days between args
func GetTimeDifferenceInDaysFrom(t time.Time, s time.Time) (diff int) {
	dur := s.Sub(t)
	return int(math.Floor(dur.Hours() / 24))
}
