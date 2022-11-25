package commands

import (
	"fmt"

	"github.com/aadithpm/speaker-bot/internal/data"
	"github.com/aadithpm/speaker-bot/internal/utils"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

type DungeonCommand struct {
	Name string
}

func NewDungeonCommand() (c SpeakerCommand) {
	return DungeonCommand{
		Name: Dungeon,
	}
}

func (c DungeonCommand) GetCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        Dungeon,
		Type:        discordgo.ChatApplicationCommand,
		Description: "Featured Dungeon for the week",
	}
}

func (c DungeonCommand) GetName() string {
	return c.Name
}

func (c DungeonCommand) Handler(s *discordgo.Session, d *discordgo.ApplicationCommandInteractionData) (res string, err error) {
	log.Infof("got command %v from handler", d.Name)

	data := data.ReadFeaturedContentData("./data/dungeons.json")
	current_week := utils.GetTimeDifferenceInWeeks(data.StartDate)
	content_count := len(data.ContentRotation)
	adjusted_week := current_week % content_count

	dungeon := data.ContentRotation[adjusted_week]

	str := "Featured Dungeon for this week is **%v** at %v."
	if dungeon.MasterAvailable {
		str += " Master difficulty is available!"
	}

	msg := fmt.Sprintf(
		str,
		dungeon.Name,
		dungeon.Location,
	)

	return msg, nil
}
