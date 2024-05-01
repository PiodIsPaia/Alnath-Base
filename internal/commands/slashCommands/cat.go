package slashcommands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/piodispaia/alnath/internal/api"
	"github.com/piodispaia/alnath/internal/models"
	"github.com/piodispaia/alnath/internal/settings"
)

func CatCommand() models.SlashCommand {
	return models.SlashCommand{
		Data: discordgo.ApplicationCommand{
			Name:        "cat",
			Description: "Reply with a cat image",
			Type:        discordgo.ChatApplicationCommand,
		},
		Run: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			cat, err := api.CatImage()
			if err != nil {
				fmt.Println("Error fetching cat image:", err)
				return
			}

			config, err := settings.GetConfig()
			if err != nil {
				fmt.Println("Error getting bot configuration:", err)
				return
			}

			color, err := settings.HexToDecimalColor(config.Color.Azoxo)
			if err != nil {
				fmt.Println("Error converting color:", err)
				return
			}

			embed := &discordgo.MessageEmbed{
				Image: &discordgo.MessageEmbedImage{
					URL: cat,
				},
				Color: color,
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{embed},
				},
			})
		},
	}
}
