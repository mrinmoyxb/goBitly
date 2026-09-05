package utils

import (
	"math/rand/v2"
	"strings"
)

func ShortURLGenerator(length int64) string {
	const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	var builder strings.Builder
	builder.Grow(int(length))

	for range length {
		randomIndex := rand.IntN(len(characters))
		builder.WriteByte(characters[randomIndex])
	}

	return builder.String()
}

//  CONSTRAINT "urls_user_constraint"

//         FOREIGN KEY (user_id)

//         REFERENCES users(id)

//         ON DELETE CASCADE
