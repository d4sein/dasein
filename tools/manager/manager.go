package manager

import (
	"github.com/bwmarrin/discordgo"
	commands "github.com/d4sein/Dasein/commands/fun"
	"github.com/d4sein/Dasein/events"
	"github.com/d4sein/Dasein/tools/router"
)

// InitEvents ..
func InitEvents(client *discordgo.Session) {
	client.AddHandler(events.OnMessageCreate)
}

// InitCommands ..
func InitCommands() {
	// fun
	router.MyRouter.AddCommand(commands.Ping)
}
