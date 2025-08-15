package scheduler

import (
	"friday-bot/bot"
	"friday-bot/config"
	"friday-bot/logger"
	"time"
)

// Scheduler handles the Friday posting schedule
type Scheduler struct {
	bot         *bot.Bot
	config      *config.Config
	ticker      *time.Ticker
	stopChannel chan struct{}
	lastPosted  time.Time
}

// New creates a new Scheduler instance
func New(bot *bot.Bot, cfg *config.Config) *Scheduler {
	return &Scheduler{
		bot:         bot,
		config:      cfg,
		stopChannel: make(chan struct{}),
	}
}

// Start begins the scheduling loop
func (s *Scheduler) Start() {
	// Create ticker for the check interval
	s.ticker = time.NewTicker(time.Duration(s.config.CheckInterval) * time.Minute)
	defer s.ticker.Stop()

	logger.Info("Scheduler started. Checking every %d minutes for Friday posts", s.config.CheckInterval)

	// Check immediately on startup
	s.checkAndPost()

	for {
		select {
		case <-s.ticker.C:
			s.checkAndPost()
		case <-s.stopChannel:
			logger.Info("Scheduler stopped")
			return
		}
	}
}

// Stop gracefully stops the scheduler
func (s *Scheduler) Stop() {
	close(s.stopChannel)
	if s.ticker != nil {
		s.ticker.Stop()
	}
}

// checkAndPost checks if it's time to post and posts if needed
