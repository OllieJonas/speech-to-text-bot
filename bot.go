package main

import (
	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	ApiKey  string
	Session *discordgo.Session
}

func (b *Bot) init() {
	discordBot, err := discordgo.New("Bot " + b.ApiKey)

	if err != nil {
		panic("Error creating Discord session! " + err.Error())
	}

	discordBot.AddHandler(func(_ *discordgo.Session, m *discordgo.MessageCreate) {
		_ = getRouter().FindAndExecute(discordBot, COMMAND_PREFIX, discordBot.State.User.ID, m.Message)
	})

	discordBot.AddHandler(ready)

	b.Session = discordBot

	err = discordBot.Open()

	if err != nil {
		panic("Error opening Connection! " + err.Error())
	}
}

// Called from AddHandler(ready)
func ready(session *discordgo.Session, event *discordgo.Ready) {
	_ = session.UpdateGameStatus(1, "!stt")
}

func (b *Bot) run() {

}

func (b *Bot) close() {
	err := b.Session.Close()

	if err != nil {
		panic("Bot Error on Close (man that sucks lmao) " + err.Error())
	}
}
