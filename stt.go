package main

import "github.com/bwmarrin/discordgo"

func getVoiceChannelFor(guild *discordgo.Guild, message *discordgo.Message) *discordgo.VoiceState {
	for _, vc := range guild.VoiceStates {
		if vc.UserID == message.Author.ID {
			return vc
		}
	}
	return nil
}

func joinVoiceChannel(session *discordgo.Session, guildID, channelID string) (*discordgo.VoiceConnection, error) {
	connection, err := session.ChannelVoiceJoin(guildID, channelID, true, false)

	if err != nil {
		return nil, err
	}
	return connection, nil
}
