package utils

import (
	"math/rand/v2"
	"strings"
)

const characters string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
const base62Base = int64(len(characters))

// Phase 1 approach
func ShortURLGenerator(length int64) string {
	var builder strings.Builder
	builder.Grow(int(length))

	for range length {
		randomIndex := rand.IntN(len(characters))
		builder.WriteByte(characters[randomIndex])
	}

	return builder.String()
}

// Phase 2 approach
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func padLeft(s string, minLength int) string {
	if len(s) >= minLength {
		return s
	}

	padding := strings.Repeat(string(characters[0]), minLength-len(s))
	return padding + s
}

func EncodeBase62(id int64) string {
	if id == 0 {
		return string(characters[0])
	}

	var builder strings.Builder
	for id > 0 {
		remainder := id % base62Base
		builder.WriteByte(characters[remainder])
		id = id / base62Base
	}

	encoded := reverseString(builder.String())
	return padLeft(encoded, 5)
}
