package main

import (
	"github.com/Necroforger/dgrouter/exrouter"
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

		makeContext(guildId)

		if isTextToSpeechAlreadyActiveIn(guildId) {
			markForDisconnect(guildId)
			return
		}

		guild, err := session.State.Guild(guildId)

		if err != nil {
			_, _ = ctx.Reply("something is seriously seriously fucked up (cant find guild) | " + err.Error())
			return
		}

		voiceChannel := getVoiceChannelFor(guild, ctx.Msg.Author)

		if voiceChannel == nil {
			_, _ = ctx.Reply("You aren't currently in a voice channel!")
			return
		}

		voiceChannelId := voiceChannel.ChannelID

		connection, err := connect(session, guildId, voiceChannelId, ctx.Msg.ChannelID)

		if err != nil {
			_, _ = ctx.Reply("Error connecting to voice channel! | " + err.Error())
			return
		}

		if connection == nil {
			_, _ = ctx.Reply("Error connecting to voice channel! | Connection is null")
		}

		go listen(session, connection, guildId)

		_, _ = ctx.Reply("You are currently in voice channel " + voiceChannel.ChannelID)
	}
}

func yanSucks() exrouter.HandlerFunc {
	return func(ctx *exrouter.Context) {
		_, _ = ctx.Reply("yan sucks")
	}
}
