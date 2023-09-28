package backend_utils

import (
	"errors"
	"math/rand"
	"time"
)

func RandomListChoice[T any](m []T) (*T, error) {
	if len(m) == 0 {
		return nil, errors.New("empty list")
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(m))

	return &m[randomIndex], nil
}

func StringExistInArray(arr []string, element string) bool {
	for _, value := range arr {
		if value == element {
			return true
		}
	}
	return false
}
