package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/aadithpm/speaker-bot/internal/handlers"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var Session *discordgo.Session
var Token string

func main() {
	if err := godotenv.Load(); err != nil {
		log.Warn("error loading .env")
	}

	Token = os.Getenv("DISCORD_BOT_TOKEN")
	if Token == "" {
		log.Fatal("no bot token found, exiting..")
		os.Exit(1)
	}

	Session, _ = discordgo.New(Token)
	Session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)

	handlers.AddHandlers(Session)

	err := Session.Open()
	if err != nil {
		log.Fatal("error opening connection to Discord: %v\n", err)
		os.Exit(1)
	}

	usd := discordgo.UpdateStatusData{
		Activities: []*discordgo.Activity{
			{
				Name: "Ghaul",
				Type: 0,
			},
		},
	}
	err = Session.UpdateStatusComplex(usd)
	if err != nil {
		log.Warn("error updating bot status")
	}

	log.Info("bot is running")
	log.Info(Session.Identify.Intents)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	Session.Close()
}
