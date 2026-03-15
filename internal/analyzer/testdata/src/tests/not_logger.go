package a

import (
	"errors"
	"fmt"
)

func notLogs() {
	fmt.Println("User authenticated successfully!")
	fmt.Printf("password=%d", 5)
	errors.New("Все сломалось")
}
