package common

import (
	"github.com/google/uuid"
	"strings"
)

func Uuid() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}
