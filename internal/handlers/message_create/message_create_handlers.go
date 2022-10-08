package message_create

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/aadithpm/speaker-bot/internal/utils"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// AddHandlers add handlers for MessageCreate, fires when a new message is sent on the server
func AddHandlers(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		log.Info(m.Message.Content)

		// Add new handlers here
		alertAdaToDestinyTalk(s, m)
		alertFortniteToChannel(s, m)
	})
}

// Alert on specific mods to destiny-talk
func alertAdaToDestinyTalk(s *discordgo.Session, m *discordgo.MessageCreate) {
	// IDs:
	// daily-reset: 785932523028873277
	// destiny-talk: 785889673691791451
	// hype-emoji: <:hype:798225963422580747>
	if m.ChannelID == "785932523028873277" && len(m.Embeds) > 0 {
		for _, e := range m.Embeds {
			var mods []string
			if e.Title == "Ada-1, Armor Synthesis" {
				log.Infof("found Ada-1 embed, attempting to check for mods...")
				for _, f := range e.Fields {
					if f.Name == "Daily Mods" {
						mods = strings.Split(f.Value, "\n")
					}
				}
			}

			var fMods []string
			for _, m := range mods {
				r, _ := regexp.Compile(`([eE]lemental [wW]ell|[cC]harged [wW]ith [lL]ight|[wW]armind [cC]ell)\s*`)
				res := r.MatchString(m)
				if res {
					m = strings.TrimSpace(m)
					fMods = append(fMods, m)
				}
			}

			if len(fMods) == 0 {
				log.Infof("Ada-1 is not selling any interesting mods today")
				return
			}

			for i, f := range fMods {
				fMods[i] = fmt.Sprintf("**%v**", f)
			}
			mMods := strings.Join(fMods[:], ", ")

			c, err := s.GuildChannels(m.GuildID)
			if err != nil {
				log.Warnf("error getting channels: %v", err)
			}
			tc, err := utils.GetChannelById(c, "785889673691791451")
			if err != nil {
				log.Warnf("error getting destiny-talk channel: %v", err)
				return
			}

			err = utils.SendMessageInChannel(s, tc, fmt.Sprintf(`Ada-1 is selling %v until tomorrow's reset.`, mMods))
			if err != nil {
				log.Warnf("error sending message for Ada-1 reset: %v", err)
			}
		}
	}
}

// Spam Pandi when someone says forkknife for the memes
func alertFortniteToChannel(s *discordgo.Session, m *discordgo.MessageCreate) {
	// 313141452991627266
	msg := `<@313141452991627266> https://tenor.com/view/we-like-fortnite-we-like-fortnite-speed-up-gif-26419282`
	c, err := s.GuildChannels(m.GuildID)
	if err != nil {
		log.Warnf("error getting channels: %v", err)
	}
	tc, err := utils.GetChannelById(c, m.ChannelID)
	if err != nil {
		log.Warnf("error getting destiny-talk channel: %v", err)
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
