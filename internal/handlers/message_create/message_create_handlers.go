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
		alertNewLfgToRaidLfg(s, m)
		alertAdaToDestinyTalk(s, m)
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
			tc, err := utils.GetChannelByName(c, "destiny-talk")
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

// Alert on all new lfgs in lfg-management to raid-lfg
func alertNewLfgToRaidLfg(s *discordgo.Session, m *discordgo.MessageCreate) {
	// IDs:
	// lfg-management: 785978702483292240
	// raid-lfg: 785972816570613791
	// hype-emoji: <:hype:798225963422580747>
	if m.ChannelID == "785978702483292240" {
		r, _ := regexp.Compile(`LFG Post: \*\*(?P<lfgId>\d+)\*\* created`)
		res := r.FindStringSubmatch(m.Content) // golang StringSubmatch groups are so awkward
		if len(res) > 1 {
			log.Infof("attempting to post lfg id to raid-lfg...")
			joinId := res[1]
			c, err := s.GuildChannels(m.GuildID)
			if err != nil {
				log.Warnf("error getting channels: %v", err)
			}
			tc, err := utils.GetChannelByName(c, "raid-lfg")
			if err != nil {
				log.Warnf("error getting raid-lfg channel: %v", err)
				return
			}
			err = utils.SendMessageInChannel(s, tc, fmt.Sprintf(`<:hype:798225963422580747> @everyone **LFG Alert:** Please use !lfg %v to look up the LFG or go to <#940341727007498240>.  <:hype:798225963422580747>`, joinId))
			if err != nil {
				log.Warnf("error sending message for new LFG: %v", err)
			}
		}
	}
}
