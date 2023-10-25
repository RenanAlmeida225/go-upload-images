package uuid

import "github.com/google/uuid"

func GenerateConfirmationToken() string {
	return uuid.New().String()
}
