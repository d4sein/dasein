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
	// ErrNoPermission defines the error when a user does not have sufficient permission
	ErrNoPermission = "you can't run this command"
)

// Router defines the structure of the Router
type Router struct {
	Commands map[string]Command
	prefix   string
}

// Command defines the structure of a command
type Command struct {
	Name       string
	Args       Arguments
	Run        Callback
	Permission int
}

// Arguments ..
type Arguments map[string][]string

// Callback ..
type Callback func(s *discordgo.Session, m *discordgo.MessageCreate, a Arguments)

// New returns a new Router
func New() *Router {
	return &Router{
		Commands: make(map[string]Command),
		prefix:   ";",
	}
}

// SetPrefix sets the command prefix
func (r *Router) SetPrefix(prefix string) {
	r.prefix = prefix
}

// AddCommand Adds a new command to the router
func (r *Router) AddCommand(c Command) {
	if _, ok := r.Commands[c.Name]; ok {
		log.Fatalf("command '%s' has been declared already\n", c.Name)
	}
	r.Commands[c.Name] = c
}

func (r *Router) parseCommand(s *discordgo.Session, m *discordgo.MessageCreate) (cmd Command, err error) {
	argPrefix := "--"

	ctx := strings.Split(m.Content, " ")

	name := strings.TrimPrefix(ctx[0], r.prefix)

	ctx = reverse(ctx[1:])

	cmd, ok := r.Commands[name]
	if !ok {
		return cmd, errors.New(ErrCommandNotFound)
	}

	// p, err := s.State.UserChannelPermissions(m.Author.ID, m.ChannelID)
	// if err != nil || (p&cmd.Permission != cmd.Permission) {
	// 	return cmd, errors.New(ErrNoPermission)
	// }

	tempArgs := []string{}

	for _, s := range ctx {
		if strings.HasPrefix(s, argPrefix) {
			cmd.Args[strings.TrimPrefix(s, argPrefix)] = tempArgs

			tempArgs = []string{}
			continue
		}

		tempArgs = append(tempArgs, s)
	}

	return cmd, nil
}

// OnMessageCreateHandler handles the message create event
func (r *Router) OnMessageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID ||
		!strings.HasPrefix(m.Content, r.prefix) {
		return
	}

	if m.Message.ChannelID != "781513534470094868" &&
		m.Message.ChannelID != "619389904866115606" {
		return
	}

	cmd, err := r.parseCommand(s, m)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}

	cmd.Run(s, m, cmd.Args)
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}
