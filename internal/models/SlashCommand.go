package models

import (
	"fmt"
	"sync"

	"github.com/bwmarrin/discordgo"
)

type SlashCommandRegistry struct {
	sync.RWMutex
	commands         map[string]SlashCommand
	globalCommandIDs map[string]string
	session          *discordgo.Session
}

type SlashCommand struct {
	Data discordgo.ApplicationCommand
	Run  func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func NewSlashCommandRegistry(s *discordgo.Session) *SlashCommandRegistry {
	return &SlashCommandRegistry{
		commands:         make(map[string]SlashCommand),
		globalCommandIDs: make(map[string]string),
		session:          s,
	}
}

func (scr *SlashCommandRegistry) RegisterCommand(command SlashCommand) {
	scr.Lock()
	defer scr.Unlock()
	scr.commands[command.Data.Name] = command

	err := scr.registerGlobalCommand(command.Data)
	if err != nil {
		fmt.Println("Error registering global command:", err)
	}
}

func (scr *SlashCommandRegistry) GetCommand(commandName string) (SlashCommand, bool) {
	scr.RLock()
	defer scr.RUnlock()
	command, ok := scr.commands[commandName]
	return command, ok
}

func (scr *SlashCommandRegistry) HandleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		scr.RLock()
		defer scr.RUnlock()
		if command, ok := scr.commands[i.ApplicationCommandData().Name]; ok {
			command.Run(s, i)
		}
	}
}

func (scr *SlashCommandRegistry) registerGlobalCommand(commandData discordgo.ApplicationCommand) error {
	var commandID string
	var err error

	if globalID, exists := scr.globalCommandIDs[commandData.Name]; exists {
		commandID = globalID
		_, err = scr.session.ApplicationCommandEdit(scr.session.State.User.ID, "", commandID, &commandData)
		if err != nil {
			return fmt.Errorf("failed to update global slash command: %v", err)
		}
	} else {
		createdCommand, err := scr.session.ApplicationCommandCreate(scr.session.State.User.ID, "", &commandData)
		if err != nil {
			return fmt.Errorf("failed to register global slash command: %v", err)
		}
		commandID = createdCommand.ID
	}

	scr.globalCommandIDs[commandData.Name] = commandID

	return nil
}
