package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/d4sein/Dasein/pkg/router"
)

func init() {
	Router.AddCommand(router.Command{
		Name: "ping",
		Args: make(map[string][]string),
		Run: func(s *discordgo.Session, m *discordgo.MessageCreate, a router.Arguments) {
			log.Println(a)
			s.ChannelMessageSend(m.ChannelID, "p-pong :flushed:")
		},
		Permission: discordgo.PermissionSendMessages,
	})
}
