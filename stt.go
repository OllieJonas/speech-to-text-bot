package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type STTContext struct {
	IsConnected    bool
	CloseChan      chan bool
	VoiceChannelID string
	TextChannelID  string
	Connection     *discordgo.VoiceConnection
	Message        *discordgo.MessageEmbed
}

var contextMap = make(map[string]*STTContext)

// instantiates context
func makeContext(guildId string) {
	if contextMap[guildId] == nil {
		fmt.Println("Making context for guild " + guildId)
		contextMap[guildId] = &STTContext{
			IsConnected: false,
			CloseChan:   make(chan bool, 1),
			Connection:  nil,
			Message:     createTextBoard(),
		}
	}
}

// gets the voice channel for a given user
func getVoiceChannelFor(guild *discordgo.Guild, author *discordgo.User) *discordgo.VoiceState {
	for _, vc := range guild.VoiceStates {
		if vc.UserID == author.ID {
			return vc
		}
	}
	return nil
}

// checks whether text to speech is already enabled in a guild
func isTextToSpeechAlreadyActiveIn(guildId string) bool {
	return contextMap[guildId].IsConnected
}

// connects the bot to the voice channel
func connect(session *discordgo.Session, guildID, voiceChannelID, textChannelID string) (*discordgo.VoiceConnection, error) {
	connection, err := session.ChannelVoiceJoin(guildID, voiceChannelID, true, false)

	if err != nil {
		// weird ass solution from here: https://github.com/jonas747/yagpdb/issues/284
		if _, ok := session.VoiceConnections[guildID]; ok {
			connection = session.VoiceConnections[guildID]
		} else {
			return nil, err
		}
	}

	contextMap[guildID].VoiceChannelID = voiceChannelID
	contextMap[guildID].TextChannelID = textChannelID
	contextMap[guildID].Connection = connection
	contextMap[guildID].IsConnected = true

	return connection, nil
}

// listens to input from bot
func listen(session *discordgo.Session, voice *discordgo.VoiceConnection, guildId string) {
	stream := voice.OpusRecv
	closeChan := contextMap[guildId].CloseChan

	for {
		select {
		case input := <-stream:
			fmt.Printf("%b\n", input.Opus)
		case <-closeChan:
			disconnect(session, guildId)
		}
	}
}

// sends a boolean on the close channel for listen to disconnect
func markForDisconnect(guildId string) {
	contextMap[guildId].CloseChan <- false
}

// random stuff that needs to happen when the bot disconnects
func disconnect(session *discordgo.Session, guildId string) {
	context := contextMap[guildId]
	context.IsConnected = false

	_, _ = session.ChannelMessageSend(context.TextChannelID, "Goodbye!")
	_ = context.Connection.Disconnect()
}

// creates a text board that will be modified with the input text
func createTextBoard() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       "Text-to-Speech",
		Description: "",
		Footer:      nil,
		Image:       nil,
		Thumbnail:   nil,
		Video:       nil,
		Provider:    nil,
		Author:      nil,
		Fields:      make([]*discordgo.MessageEmbedField, 1, 1),
	}
}
