package settings

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/piodispaia/alnath/internal/models"
)

func GetConfig() (models.Config, error) {
	dir, err := os.Getwd()
	if err != nil {
		return models.Config{}, err
	}

	configFilePath := filepath.Join(dir, "internal/settings/settings.json")

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		return models.Config{}, fmt.Errorf("settings.json file not found at %s", configFilePath)
	}

	file, err := os.Open(configFilePath)
	if err != nil {
		return models.Config{}, err
	}
	defer file.Close()

	var config models.Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return models.Config{}, err
	}

	return config, nil
}

func HexToDecimalColor(hexColor string) (int, error) {
	if len(hexColor) > 0 && hexColor[0] == '#' {
		hexColor = hexColor[1:]
	}

	color, err := strconv.ParseInt(hexColor, 16, 32)
	if err != nil {
		return 0, fmt.Errorf("failed to convert color: %v", err)
	}

	return int(color), nil
}

func GetEmoji(s *discordgo.Session, id string, guildID string) (*discordgo.Emoji, error) {
	emoji, err := s.State.Emoji(guildID, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get emoji: %v", err)
	}

	return emoji, nil
}
