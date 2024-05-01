package slashcommands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/piodispaia/alnath/internal/models"
)

func PingCommand() models.SlashCommand {
	return models.SlashCommand{
		Data: discordgo.ApplicationCommand{
			Name:        "ping",
			Description: "Reply as pong",
			Type:        discordgo.ChatApplicationCommand,
		},
		Run: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   discordgo.MessageFlagsEphemeral,
					Content: "Pong",
				},
			}); err != nil {
				fmt.Println(err.Error())
				return
			}
		},
	}
}
