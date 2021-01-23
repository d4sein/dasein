package router

import "github.com/bwmarrin/discordgo"

var (
	// MyRouter ..
	MyRouter = Router{
		Commands: make(map[string]Command),
	}
)

// Router ..
type Router struct {
	Commands map[string]Command
}

// Command ..
type Command struct {
	Name string
	Run  callback
}

type callback func(s *discordgo.Session, m *discordgo.MessageCreate)

// AddCommand ..Adds a new command to the router
func (r Router) AddCommand(c Command) {
	r.Commands[c.Name] = c
}
