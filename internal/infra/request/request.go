package request

import (
	"time"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	return uuid.New().String()
}

func GetCurrentTime() string {
	return time.Now().Format(time.RFC3339)
}
