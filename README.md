# Alnath Base

Alnath is a Discord bot developed using DiscordGo and GORM.

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/PiodIsPaia/Alnath-Base.git
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Set up the configuration

- Navigate to `internal/settings/settings.json`.
- Add your Discord bot token to the `core.token` field.

```json
{
    "core": {
        "token": "YOUR_DISCORD_BOT_TOKEN",
        "guildID": "YOUR_SERVER_ID_HERE",
        "prefix": "YOUR_PREFIX_HERE"
    }
}

```

### 4. Build and run the bot

```bash
go build -o alnath cmd/bot/main.go
./alnath
```

---
