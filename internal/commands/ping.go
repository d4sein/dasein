package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/d4sein/Dasein/pkg/helper"
	"github.com/d4sein/Dasein/pkg/router"
	"github.com/d4sein/Dasein/pkg/typeguard"
)

func init() {
	Router.AddCommand(router.Command{
		Name: "ping",
		Args: router.Arguments{
			"n": {
				To: typeguard.WantArrInt(),
			},
		},
		Run: func(s *discordgo.Session, m *discordgo.MessageCreate, a router.Arguments) {
			x, _ := a["n"].Output.ToArrInt()

			s.ChannelMessageSend(m.ChannelID,
				fmt.Sprintf("you ping pong me %v times", helper.Sum(x)*10))
		},
		Permission: discordgo.PermissionSendMessages,
	})
}
