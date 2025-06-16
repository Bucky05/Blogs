package services

import (
	"github.com/google/uuid"
)

// GenerateUUID generates a new UUID string.
func GenerateUUID() string {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		panic("Failed to generate UUID: " + err.Error())
	}

	return newUUID.String()
}
