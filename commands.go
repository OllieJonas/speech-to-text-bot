package main

import "github.com/Necroforger/dgrouter/exrouter"

const COMMAND_PREFIX = "!"

func getRouter() *exrouter.Route {

	router := exrouter.New()

	router.On("yansucks", yanSucks())

	return router
}

func yanSucks() exrouter.HandlerFunc {
	return func(ctx *exrouter.Context) {
		_, _ = ctx.Reply("yan sucks")
	}
}