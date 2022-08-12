package handlers

import (
	"fmt"
	"regexp"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func AddHandlers(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		log.Info(m.Message.Content)

		// Message handlers by channel
		AlertNewLfgToRaidLfg(s, m)
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
		tc, err := GetChannelByName(c, "welcome-wagon")
		if err != nil {
			log.Warnf("error getting channel name %v")
		}

		msg := fmt.Sprintf("Hello <@%s>, welcome to the ARCH server! Please review <#786618669408976936> and register with <#785909633352466442>", m.User.ID)
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

// Alert on all new lfgs in lfg-management to raid-lfg
func AlertNewLfgToRaidLfg(s *discordgo.Session, m *discordgo.MessageCreate) {
	// IDs:
	// lfg-list: 940341727007498240
	// lfg-management: 785978702483292240
	// raid-lfg: 785972816570613791
	// hype-emoji: <:hype:798225963422580747>
	if m.ChannelID == "785978702483292240" {
		r, _ := regexp.Compile(`LFG Post: \*\*(?P<lfgId>\d+)\*\* created`)
		res := r.FindStringSubmatch(m.Content) // golang StringSubmatch groups are so awkward
		if len(res) > 1 {
			log.Infof("attempting to post lfg id to raid-lfg...")
			joinId := res[1]
			c, err := s.Channel("785972816570613791")
			if err != nil {
				log.Warnf("error getting raid-lfg channel")
				return
			}
			SendMessageInChannel(s, c, fmt.Sprintf(`<:hype:798225963422580747> @everyone **LFG Alert:** Please use !lfg %v to look up the LFG or go to <#940341727007498240>.  <:hype:798225963422580747>`, joinId))
		}
		log.Infof("%v", res)
	}
}
