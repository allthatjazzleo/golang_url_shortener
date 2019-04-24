package random

import (
	"math/rand"
	"time"
)

// GenerateRandom to 8 characters
func GenerateRandom(n int) string {
	rand.Seed(time.Now().UnixNano())
	randRune := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = randRune[rand.Intn(len(randRune))]
	}
	return string(b)
}
