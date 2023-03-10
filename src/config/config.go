package config

import "os"

type Config struct {
	OpenAIApiKey string
	DiscordToken string
	GuildID      string
}

func CreateConfig() *Config {
	var config Config

	config.DiscordToken = os.Getenv("DISCORD_TOKEN")
	config.OpenAIApiKey = os.Getenv("OPEN_AI_API_KEY")
	config.GuildID = os.Getenv("GUILD_ID")

	return &config
}
