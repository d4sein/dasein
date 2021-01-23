package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/d4sein/Dasein/config"
	"github.com/d4sein/Dasein/internal/commands"
	"github.com/d4sein/Dasein/pkg/router"
)

var Router = router.New()

func main() {
	client, err := discordgo.New("Bot " + config.Config.Token)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Open()
	if err != nil {
		log.Fatal(err)
	}

	Router.AddCommand(commands.Ping)
	client.AddHandler(Router.OnMessageCreate)

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	client.Close()
}
