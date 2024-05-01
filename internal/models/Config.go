package models

type Config struct {
	Core   CoreConfig `json:"core"`
	Emojis Emojis     `json:"emojis"`
	Color  Color      `json:"colors"`
}

type CoreConfig struct {
	Token   string `json:"token"`
	Prefix  string `json:"prefix"`
	GuildID string `json:"guildID"`
}

type Emojis struct {
	Success string `json:"success"`
}

type Color struct {
	Success    string `json:"success"`
	Danger     string `json:"danger"`
	Primary    string `json:"primary"`
	Secondary  string `json:"secondary"`
	Default    string `json:"default"`
	Info       string `json:"info"`
	Warning    string `json:"warning"`
	Magic      string `json:"magic"`
	Green      string `json:"green"`
	Fuchsia    string `json:"fuchsia"`
	Azoxo      string `json:"azoxo"`
	Developer  string `json:"developer"`
	Balance    string `json:"balance"`
	Brilliance string `json:"brilliance"`
	Nitro      string `json:"nitro"`
	Bravery    string `json:"bravery"`
}
