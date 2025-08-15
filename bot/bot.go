package bot

import (
        "fmt"
        "friday-bot/config"
        "friday-bot/images"
        "friday-bot/logger"
        "strings"

        tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Bot represents the Telegram bot
type Bot struct {
        api           *tgbotapi.BotAPI
        config        *config.Config
        imageManager  *images.Manager
        updateChannel tgbotapi.UpdatesChannel
        stopChannel   chan struct{}
}

// New creates a new Bot instance
func New(cfg *config.Config) (*Bot, error) {
        if err := cfg.Validate(); err != nil {
                return nil, fmt.Errorf("invalid configuration: %w", err)
        }

        api, err := tgbotapi.NewBotAPI(cfg.BotToken)
        if err != nil {
                return nil, fmt.Errorf("failed to create bot API: %w", err)
        }

        logger.Info("Authorized on account %s", api.Self.UserName)

        imageManager := images.New(cfg.ImagesDir)

        return &Bot{
                api:          api,
                config:       cfg,
                imageManager: imageManager,
                stopChannel:  make(chan struct{}),
        }, nil
}

// Start begins listening for updates
func (b *Bot) Start() error {
        u := tgbotapi.NewUpdate(0)
