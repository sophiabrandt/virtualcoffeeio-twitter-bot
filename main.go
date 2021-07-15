package main

import (
	"os"

	"github.com/sophiabrandt/virtualcoffeeio-twitter-bot/bot"
)

func main() {
	os.Exit(bot.CLI(os.Args[1:]))
}
