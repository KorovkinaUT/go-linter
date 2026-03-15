package a

import (
	"log"
	"log/slog"
)

func sensitiveViolation() {
	log.Printf("password: %d", 5) // want "log message should not contain sensitive data" "log message should not contain special symbols or emoji"
	slog.Info("password=value")   // want "log message should not contain sensitive data" "log message should not contain special symbols or emoji"
	slog.Info("api key = value")  // want "log message should not contain sensitive data" "log message should not contain special symbols or emoji"
	slog.Info("token: value")     // want "log message should not contain sensitive data" "log message should not contain special symbols or emoji"
	slog.Info("secret=value")     // want "log message should not contain sensitive data" "log message should not contain special symbols or emoji"
}
