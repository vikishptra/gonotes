package util

import "github.com/google/uuid"

func GenerateID(n int) string {
	ID := uuid.New().String()

	return ID
}

// https://github.com/matoous/go-nanoid
// https://github.com/segmentio/ksuid
