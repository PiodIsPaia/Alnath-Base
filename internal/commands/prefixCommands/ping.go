package prefixcommands

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/piodispaia/alnath/internal/models"
	"github.com/piodispaia/alnath/internal/settings"
)

func PingCommand() models.MessageCreate {
	return models.MessageCreate{
		Trigger: "ping",
		Aliases: []string{"ws"},
		Run: func(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
			message, err := s.ChannelMessageSend(m.ChannelID, "Pong 🏓")
			if err != nil {
				fmt.Println("Error sending message:", err)
				return
			}

			time.Sleep(2 * time.Second)

			config, err := settings.GetConfig()
			if err != nil {
				fmt.Println("Error getting config:", err)
				return
			}

			color, err := settings.HexToDecimalColor(config.Color.Success)
			if err != nil {
				fmt.Println("Error converting color:", err)
				return
			}

			emoji, err := settings.GetEmoji(s, config.Emojis.Success, config.Core.GuildID)
			if err != nil {
				fmt.Println("Error getting emoji:", err)
				return
			}

			content := ""
			_, err = s.ChannelMessageEditComplex(&discordgo.MessageEdit{
				ID:      message.ID,
				Channel: message.ChannelID,
				Content: &content,
				Embed: &discordgo.MessageEmbed{
					Description: fmt.Sprintf("<:%s:%s> Client Ping: ``%d``ms", emoji.Name, emoji.ID, s.HeartbeatLatency().Milliseconds()),
					Color:       color,
				},
			})
			if err != nil {
				fmt.Println("Error editing message:", err)
				return
			}
		},
	}
}
