package bot

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// CLI runs the Twitter bot and returns its exit status.
func CLI(args []string) int {
	var app appEnv
	log.Println()
	err := app.fromArgs(args)
	if err != nil {
		return 2
	}
	if err = app.run(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		return 1
	}
	return 0
}

type appEnv struct {
	at string
	as string
	ck string
	cs string
}

func (app *appEnv) fromArgs(args []string) error {
	fl := flag.NewFlagSet("virtualcoffee-twitter-bot", flag.ContinueOnError)
	fl.StringVar(&app.at, "ac", "abcdefg", "Twitter Access Token")
	fl.StringVar(&app.as, "act", "abcdefg", "Twitter Access Token Secret")
	fl.StringVar(&app.ck, "ck", "abcdefg", "Twitter Consumer Key")
	fl.StringVar(&app.cs, "cs", "abcdefg", "Twitter Consumer Secret")

	if err := fl.Parse(args); err != nil {
		return err
	}
	return nil
}

func (app *appEnv) run() error {
	return nil
}
