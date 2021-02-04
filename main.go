package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if len(os.Args) != 2 {
		panic("API Key not Specified! :(")
	}

	apiKey := os.Args[1]
	bot := Bot{ApiKey: apiKey}

	bot.init()
	bot.run()

	handleTerminatingSignals(&bot)
}

func handleTerminatingSignals(bot *Bot) {
	terminatingChan := make(chan os.Signal, 1)
	signal.Notify(terminatingChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<- terminatingChan

	bot.close()
}
