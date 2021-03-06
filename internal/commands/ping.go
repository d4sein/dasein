package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/d4sein/Dasein/pkg/router"
)

func init() {
	Router.AddCommand(router.Command{
		Name: "ping",
		Run: func(s *discordgo.Session, m *discordgo.MessageCreate, a router.Arguments) {
			s.ChannelMessageSend(m.ChannelID, "p-pong :flushed:")
		},
		Permission: discordgo.PermissionSendMessages,
	})
}
