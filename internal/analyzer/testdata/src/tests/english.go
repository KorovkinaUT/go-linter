package a

import (
	"log"
	"log/slog"
)

func englishViolation() {
	slog.Info("user аuthenticated успешна") // want "log message should contain only english letters"
	log.Print("пользователь создан")        // want "log message should contain only english letters"
}
