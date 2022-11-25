package utils

import (
	"fmt"
	"math"
	"time"

	"github.com/aadithpm/speaker-bot/internal/data"
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
		if channel.Name == n {
			log.Infof("found channel %v %v %v", channel.Name, channel.ID, channel.Type)
			return channel, nil
		}
	}
	return nil, fmt.Errorf("channel name %v not found", n)
}

// GetChannelById Gets channel by ID from list of channels, sets err if not found
func GetChannelById(c []*discordgo.Channel, i string) (channel *discordgo.Channel, err error) {
	for _, channel := range c {
		if channel.ID == i {
			log.Infof("found channel %v %v %v", channel.Name, channel.ID, channel.Type)
			return channel, nil
		}
	}
	return nil, fmt.Errorf("channel ID %v not found", i)
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

// GetTimeDifferenceInWeeks Gets the time difference in weeks between time arg and current time
func GetTimeDifferenceInWeeks(t time.Time) (diff int) {
	return GetTimeDifferenceInDays(t) / 7
}

// GetCurrentSeasonWeek Gets the current week number of the season
func GetCurrentSeasonWeek() (week int) {
	return GetTimeDifferenceInWeeks(data.ReadSeasonData("./data/season.json").StartDate)
}
