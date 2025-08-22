# Friday Bot - Docker Deployment

This guide explains how to run Friday Bot using Docker containers.

## Prerequisites

- Docker installed on your system
- Bot credentials (BOT_TOKEN and CHAT_ID)

## Quick Start with Docker

### 1. Download all project files to your local machine

### 2. Build the Docker image
```bash
docker build -t friday-bot .
```

### 2. Run the container
```bash
docker run -d \
  --name friday-bot \
  --restart unless-stopped \
  -e BOT_TOKEN="your_telegram_bot_token" \
  -e CHAT_ID="your_chat_id" \
  -v $(pwd)/sample_images:/home/friday/sample_images:ro \
  friday-bot
```

## Using Docker Compose (Recommended)

### 1. Set up environment variables
```bash
cp .env.example .env
# Edit .env with your bot credentials
```

### 2. Start the service
```bash
docker-compose up -d
```

### 3. View logs
```bash
docker-compose logs -f friday-bot
```

### 4. Stop the service
```bash
docker-compose down
```

## Configuration

### Environment Variables
- `BOT_TOKEN` - Your Telegram bot token (required)
- `CHAT_ID` - Target chat ID for posts (required)
- `POST_HOUR` - Hour to post (0-23, default: 9)
