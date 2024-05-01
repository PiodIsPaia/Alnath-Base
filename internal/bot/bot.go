package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/piodispaia/alnath/internal/handler"
)

func Start(token string) {
	dg, err := createDiscordSession(token)
	if err != nil {
		log.Fatal("Error creating Discord session:", err)
	}

	err = setupHandlers(dg)
	if err != nil {
		log.Fatal("Error setting up handlers:", err)
	}

	defer dg.Close()

	log.Printf("Logged in as: %s", dg.State.User.Username)

	select {}
}

func createDiscordSession(token string) (*discordgo.Session, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	dg.Identify.Intents = discordgo.IntentsAll

	err = dg.Open()
	if err != nil {
		return nil, err
	}

	return dg, nil
}

func setupHandlers(s *discordgo.Session) error {
	err := handler.SlashCommand(s)
	if err != nil {
		return err
	}

	err = handler.Component(s)
	if err != nil {
		return err
	}

	err = handler.MessageCreate(s)
	if err != nil {
		return err
	}

	return nil
}
