package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/d4sein/Dasein/pkg/router"
)

func init() {
	Router.AddCommand(router.Command{
		Name:       "kick",
		Permission: discordgo.PermissionKickMembers,
		Run:        func(s *discordgo.Session, m *discordgo.MessageCreate, a router.Arguments) {},
	})
}
