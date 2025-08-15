package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Config holds all configuration for the bot
type Config struct {
	BotToken      string `json:"bot_token"`
	ChatID        int64  `json:"chat_id"`
	PostHour      int    `json:"post_hour"`
	PostMinute    int    `json:"post_minute"`
	ImagesDir     string `json:"images_dir"`
	CheckInterval int    `json:"check_interval_minutes"`
	TimeZone      string `json:"timezone"`
}

// Load configuration from environment variables and config file
func Load() (*Config, error) {
	config := &Config{
		PostHour:      9,  // Default to 9 AM
		PostMinute:    0,  // Default to 0 minutes
		ImagesDir:     "sample_images",
		CheckInterval: 30, // Check every 30 minutes
		TimeZone:      "UTC",
	}

	// Load from config file if exists
	if err := loadFromFile(config); err != nil {
		// Config file is optional, just log the error
		fmt.Printf("Config file not found or invalid, using defaults: %v\n", err)
	}
