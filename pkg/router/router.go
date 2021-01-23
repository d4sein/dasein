package router

import (
	"errors"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/d4sein/Dasein/config"
)

const (
	// ErrCommandNotFound ...
	ErrCommandNotFound = "command not found"
)

// Router ...
type Router struct {
	Commands map[string]Command
}

// Command ..
type Command struct {
	Name string
	Run  callback
}

type callback func(s *discordgo.Session, m *discordgo.MessageCreate)

func New() Router {
	return Router{
		Commands: make(map[string]Command),
	}
}

// AddCommand ..Adds a new command to the router
func (r Router) AddCommand(c Command) {
	r.Commands[c.Name] = c
}

func (r Router) parseCommand(s *discordgo.Session, m *discordgo.MessageCreate) (Command, error) {
	ctx := strings.Split(m.Content, " ")

	name := strings.TrimPrefix(ctx[0], config.Config.Prefix)

	if cmd, ok := r.Commands[name]; ok {
		return cmd, nil
	}

	return Command{}, errors.New(ErrCommandNotFound)
}

// OnMessageCreate ...
func (r Router) OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	cmd, err := r.parseCommand(s, m)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}

	cmd.Run(s, m)
}
