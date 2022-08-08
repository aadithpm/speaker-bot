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

type Config struct {
	Token string `json:"token"`
}

var Session *discordgo.Session
var Token string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Warn("No .env file found")
	}

	Token = os.Getenv("DISCORD_BOT_TOKEN")
}

func init() {
	Session, _ = discordgo.New(Token)
	Session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)
}

func main() {
	err := Session.Open()
	if err != nil {
		log.Fatal("error opening connection to Discord: %s\n", err)
		os.Exit(1)
	}

	handlers.AddHandlers(Session)

	log.Info("Bot is running")
	log.Info(Session.Identify.Intents)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	Session.Close()
}
