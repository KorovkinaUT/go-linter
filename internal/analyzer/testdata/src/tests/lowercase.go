package a

import (
	"log"
	"log/slog"
)

func lowercaseViolation() {
	slog.Info("User authenticated successfully") // want "log message should start with a lowercase letter"
	log.Print("Auth success")                    // want "log message should start with a lowercase letter"
}
