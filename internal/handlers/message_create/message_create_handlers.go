package message_create

import (
	"fmt"
	"regexp"

	"github.com/aadithpm/speaker-bot/internal/utils"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// #oxides-bot-playground
const BotPlaygroundChannel = "832796289342242817"

// ItzPandi#9468
const PandiUserId = "313141452991627266"

// Fortnite GIF
const FortniteGif = "https://tenor.com/view/we-like-fortnite-we-like-fortnite-speed-up-gif-26419282"

// Rhulk GIF
const RhulkGif = "https://tenor.com/view/ratio-rhulk-destiny2-gif-25108424"

// AddHandlers add handlers for MessageCreate, fires when a new message is sent on the server
func AddHandlers(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		log.Info(m.Message.Content)

		// Add new handlers here
		// test in #oxides-bot-playground - 832796289342242817

		alertFortniteToChannel(s, m)
		alertRatio(s, m)
	})
}

// Spam Pandi when someone says forkknife for the memes
func alertFortniteToChannel(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := fmt.Sprintf(`<@%v> %v`, PandiUserId, FortniteGif)
	r, _ := regexp.Compile(`(?i)fortnite`)
	res := r.MatchString(m.Content)

	if res && m.Content != msg && m.Type != discordgo.MessageTypeChatInputCommand {
		c, err := s.GuildChannels(m.GuildID)
		if err != nil {
			log.Warnf("error getting channels: %v", err)
		}
		tc, err := utils.GetChannelById(c, m.ChannelID)
		if err != nil {
			log.Warnf("error getting channel to msg: %v", err)
			return
		}

		err = utils.SendMessageInChannel(s, tc, msg)
		if err != nil {
			log.Warnf("error sending fortnite message: %v", err)
		}
	}
}

func alertRatio(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := fmt.Sprintf(`%v`, RhulkGif)
	r, _ := regexp.Compile(`(?i)ratio`)
	res := r.MatchString(m.Content)

	if res && m.Content != msg && m.Type != discordgo.MessageTypeChatInputCommand {
		c, err := s.GuildChannels(m.GuildID)
		if err != nil {
			log.Warnf("error getting channels: %v", err)
		}
		tc, err := utils.GetChannelById(c, m.ChannelID)
		if err != nil {
			log.Warnf("error getting channel to msg: %v", err)
			return
		}

		err = utils.SendMessageInChannel(s, tc, msg)
		if err != nil {
			log.Warnf("error sending ratio message: %v", err)
		}
	}
}
