package a

import (
	"log"
	"log/slog"
)

func specialSymbolsViolation() {
	log.Print("warning: ")               // want "log message should not contain special symbols or emoji"
	slog.Error("connection failed!!!")   // want "log message should not contain special symbols or emoji"
	slog.Warn("something went wrong...") // want "log message should not contain special symbols or emoji"
	slog.Info("server started 🚀")        // want "log message should not contain special symbols or emoji"
}
