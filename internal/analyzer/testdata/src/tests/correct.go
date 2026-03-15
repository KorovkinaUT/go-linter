package a

import (
	"log"
	"log/slog"
)

func correctStdLog() {
	log.Print("server started")
	log.Printf("request %d completed", 5)
}

func correctSlog() {
	slog.Info("user authenticated successfully")
	slog.Debug("api request completed")

	logger := slog.Default()
	logger.Warn("something went wrong")
}
