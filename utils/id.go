package utils

import (
	"strings"

	"github.com/google/uuid"
)

func MakeID() *string {
	return Pointer(strings.Split(uuid.New().String(), "-")[0])
}
