package events

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/d4sein/Dasein/tools/parser"
)

// OnMessageCreate ..
func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	cmd, err := parser.ParseCommand(s, m)
	if err != nil {
		log.Println(err.Error())
		return
	}

	cmd.Run(s, m)
}
