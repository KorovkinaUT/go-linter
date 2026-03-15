package a

import (
	"log"
	"log/slog"
)

func multipleViolations() {
	slog.Info("Password: value!") // want "log message should start with a lowercase letter" "log message should not contain special symbols or emoji" "log message should not contain sensitive data"
	log.Print("Начинаем...")      // want "log message should start with a lowercase letter" "log message should contain only english letters" "log message should not contain special symbols or emoji"
}
