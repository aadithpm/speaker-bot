package message_create

import (
	"fmt"
	"regexp"

	"github.com/aadithpm/speaker-bot/internal/data"
	"github.com/aadithpm/speaker-bot/internal/utils"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// #oxides-bot-playground
const BotPlaygroundChannel = "832796289342242817"

// ItzPandi#9468
const PandiUserId = "313141452991627266"

// Neone#0376
const NeoneUserId = "121042733199523840"

// Pikachu face GIF
const PikachuGif = "https://tenor.com/view/pikachu-shocked-face-stunned-pokemon-shocked-not-shocked-omg-gif-24112152"

// Fortnite GIF
const FortniteGif = "https://tenor.com/view/we-like-fortnite-we-like-fortnite-speed-up-gif-26419282"

// AddHandlers add handlers for MessageCreate, fires when a new message is sent on the server
func AddHandlers(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		log.Info(m.Message.Content)

		// Add new handlers here
		// test in #oxides-bot-playground - 832796289342242817

		alertFortniteToChannel(s, m)
		alertPoHToChannel(s, m)
	})
}

// Spam Pandi when someone says forkknife for the memes
func alertFortniteToChannel(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := fmt.Sprintf(`<@%v> %v`, PandiUserId, FortniteGif)
	c, err := s.GuildChannels(m.GuildID)
	if err != nil {
		log.Warnf("error getting channels: %v", err)
	}
	tc, err := utils.GetChannelById(c, m.ChannelID)
	if err != nil {
		log.Warnf("error getting channel to msg: %v", err)
		return
	}

	r, _ := regexp.Compile(`(?i)fortnite`)
	res := r.MatchString(m.Content)

	if res && m.Content != msg {
		err = utils.SendMessageInChannel(s, tc, msg)
		if err != nil {
			log.Warnf("error sending fortnite message: %v", err)
		}
	}
}

// Spam Neone when someone mentions Pit of Heresy and it's the dungeon for the week (why do I do this to myself)
func alertPoHToChannel(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := fmt.Sprintf(`<@%v> Did you know PoH is this week? %v`, NeoneUserId, PikachuGif)

	r, _ := regexp.Compile(`PoH|(?i)pit of heresy`)
	res := r.MatchString(m.Content)

	if res && m.Content != msg {
		dungeons := data.ReadRotationData("./data/dungeons.json")
		current_week := utils.GetTimeDifferenceInWeeks(dungeons.StartDate)
		dungeon := dungeons.ContentRotation[current_week%len(dungeons.ContentRotation)]

		if dungeon.Name == "Pit of Heresy" {
			c, err := s.GuildChannels(m.GuildID)
			if err != nil {
				log.Warnf("error getting channels: %v", err)
			}
			tc, err := utils.GetChannelById(c, m.ChannelID)
			if err != nil {
				log.Warnf("error getting destiny-talk channel: %v", err)
				return
			}

			err = utils.SendMessageInChannel(s, tc, msg)
			if err != nil {
				log.Warnf("error sending poh message: %v", err)
			}
		}
	}
}
