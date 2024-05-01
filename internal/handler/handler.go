package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/piodispaia/alnath/internal/commands/prefixcommands"
	slashcommands "github.com/piodispaia/alnath/internal/commands/slashCommands"
	"github.com/piodispaia/alnath/internal/models"
	"github.com/piodispaia/alnath/internal/settings"
)

func MessageCreate(s *discordgo.Session) error {
	config, err := settings.GetConfig()
	if err != nil {
		return err
	}

	commandHandler := models.NewCommandHandler(config.Core.Prefix)

	commandHandler.RegisterCommand(prefixcommands.PingCommand())

	s.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		commandHandler.HandleMessage(s, m)
	})

	return nil
}

func SlashCommand(s *discordgo.Session) error {
	slashCommandRegistry := models.NewSlashCommandRegistry(s)

	slashCommandRegistry.RegisterCommand(slashcommands.PingCommand())
	slashCommandRegistry.RegisterCommand(slashcommands.CatCommand())

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		slashCommandRegistry.HandleInteraction(s, i)
	})

	return nil
}

func Component(s *discordgo.Session) error {
	componentRegistry := models.NewComponentRegistry()

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		componentRegistry.HandleInteraction(s, i)
	})

	return nil
}
