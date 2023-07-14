package utils

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

func GenerateUniqueString() string {

	return uuid.NewString()
}

func GenerateRandomUserName(FirstName string) string {

	suffix := make([]byte, 4)

	numbers := "1234567890"
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	for i := range suffix {
		suffix[i] = numbers[rng.Intn(10)]
	}

	userName := (FirstName + string(suffix))

	return strings.ToLower(userName)
}
