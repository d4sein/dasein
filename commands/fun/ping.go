package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/d4sein/Dasein/tools/router"
)

var (
	// Ping ..
	Ping = router.Command{
		Name: "ping",
		Run:  callback,
	}
)

func callback(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "p-pong :flushed:")
}
