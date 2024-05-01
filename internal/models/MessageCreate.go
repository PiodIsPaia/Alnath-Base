package models

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type MessageCreate struct {
	Trigger string
	Aliases []string
	Run     func(s *discordgo.Session, m *discordgo.MessageCreate, args []string)
}

type CommandHandler struct {
	prefix   string
	commands map[string]MessageCreate
}

func NewCommandHandler(prefix string) *CommandHandler {
	return &CommandHandler{
		prefix:   prefix,
		commands: make(map[string]MessageCreate),
	}
}

func (ch *CommandHandler) RegisterCommand(command MessageCreate) {
	ch.commands[command.Trigger] = command
	for _, alias := range command.Aliases {
		ch.commands[alias] = command
	}
}

func (ch *CommandHandler) HandleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	content := m.Content
	if !strings.HasPrefix(content, ch.prefix) {
		return
	}

	parts := strings.Fields(content[len(ch.prefix):])
	if len(parts) == 0 {
		return
	}

	cmd := parts[0]
	args := parts[1:]

	command, ok := ch.commands[cmd]
	if !ok {
		return
	}

	go command.Run(s, m, args)
}
