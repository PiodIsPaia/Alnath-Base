# For a Discord bot in Go, we'll use a Go image.
FROM golang:1.22.1

# Set the working directory inside the container.
WORKDIR /app

# Copy your bot's source code to the working directory.
COPY . .

# Install your bot's dependencies (if any).
RUN go mod download

# Build your bot's executable.
RUN go build -o bot ./cmd/bot/main.go

# Expose the port that your bot uses to communicate with Discord.
EXPOSE 80

# Limit memory usage to 512MB and start the bot when the container is started.
CMD ["sh", "-c", "ulimit -m 512000 && ./bot"]
