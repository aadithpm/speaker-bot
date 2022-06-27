package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type Config struct {
	Token string `json:"token"`
}

var Session *discordgo.Session
var Token string

func init() {
	var token = os.Getenv("DISCORD_BOT_TOKEN")

	if token == "" {
		token = ReadTokenFromJson()
	}

	Token = token
}

func init() {
	Session, _ = discordgo.New(Token)
}

func ReadTokenFromJson() string {
	/**
	* If token isn't set as an environment variable (local testing)
	* Get it from a config file
	**/
	jsonFile, err := os.Open("config.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	bytes, _ := ioutil.ReadAll(jsonFile)
	var config Config

	json.Unmarshal(bytes, &config)

	return config.Token
}

func main() {
	err := Session.Open()
	if err != nil {
		log.Printf("error opening connection to Discord: %s\n", err)
		os.Exit(1)
	}

	log.Printf("Bot is running")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	Session.Close()
}
