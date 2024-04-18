package auth

import "crypto/sha512"

func Sum(value string) string {
	sum := sha512.Sum512([]byte(value))
	return string(sum[:])
}
