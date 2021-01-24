package router

import (
	"errors"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const (
	// ErrCommandNotFound defines the error when a command is not found
	ErrCommandNotFound = "command not found"
)

// Router defines the structure of the Router
type Router struct {
	Commands map[string]Command
	prefix   string
}

// Command defines the structure of a command
type Command struct {
	Name string
	Run  callback
}

type callback func(s *discordgo.Session, m *discordgo.MessageCreate)

// New returns a new Router
func New() Router {
	return Router{
		Commands: make(map[string]Command),
		prefix:   "!",
	}
}

// AddCommand Adds a new command to the router
func (r Router) AddCommand(c Command) {
	if _, ok := r.Commands[c.Name]; ok {
		log.Fatalf("command '%s' has been declared already\n", c.Name)
	}
	r.Commands[c.Name] = c
}

func (r Router) parseCommand(s *discordgo.Session, m *discordgo.MessageCreate) (Command, error) {
	ctx := strings.Split(m.Content, " ")

	name := strings.TrimPrefix(ctx[0], r.prefix)

	if cmd, ok := r.Commands[name]; ok {
		return cmd, nil
	}

	return Command{}, errors.New(ErrCommandNotFound)
}

// OnMessageCreateHandler handles the message create event
func (r Router) OnMessageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Message.ChannelID != "781513534470094868" {
		return
	}

	cmd, err := r.parseCommand(s, m)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}

	cmd.Run(s, m)
}

func (r Router) SetPrefix(prefix string) Router {
	r.prefix = prefix
	return r
}
