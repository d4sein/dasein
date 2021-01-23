package parser

import (
	"errors"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/d4sein/Dasein/config"
	"github.com/d4sein/Dasein/tools/router"
)

// ParseCommand ..
func ParseCommand(s *discordgo.Session, m *discordgo.MessageCreate) (router.Command, error) {
	ctx := strings.Split(m.Content, " ")

	name := strings.TrimPrefix(ctx[0], config.Config.Prefix)

	if cmd, ok := router.MyRouter.Commands[name]; ok {
		return cmd, nil
	}

	return router.Command{}, errors.New("notFound")
}
