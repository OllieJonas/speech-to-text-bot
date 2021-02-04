package main

import (
	"github.com/Necroforger/dgrouter/exrouter"
	"time"
)

const COMMAND_PREFIX = "!"

func getRouter() *exrouter.Route {

	router := exrouter.New()

	router.On("yansucks", yanSucks())
	router.On("stt", speechToText())

	return router
}

func speechToText() exrouter.HandlerFunc {
	return func(ctx *exrouter.Context) {
		session := ctx.Ses
		guildId := ctx.Msg.GuildID
		guild, err := session.State.Guild(guildId)

		if err != nil {
			_, _ = ctx.Reply("something is seriously seriously fucked up (cant find guild) | " + err.Error())
			return
		}

		voiceChannel := getVoiceChannelFor(guild, ctx.Msg)

		if voiceChannel == nil {
			_, _ = ctx.Reply("You aren't currently in a voice channel!")
			return
		}

		channelId := voiceChannel.ChannelID

		connection, err := joinVoiceChannel(session, guildId, channelId)

		if connection == nil || err != nil {
			_, _ = ctx.Reply("Error connecting to voice channel!")
			return
		}

		time.Sleep(500 * time.Millisecond)
		_ = connection.Disconnect()

		_, _ = ctx.Reply("You are currently in voice channel " + voiceChannel.ChannelID)
	}
}

func yanSucks() exrouter.HandlerFunc {
	return func(ctx *exrouter.Context) {
		_, _ = ctx.Reply("yan sucks")
	}
}
