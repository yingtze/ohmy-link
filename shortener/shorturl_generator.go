package shortener

import (
	"crypto/sha256"
)

func sha2560f(input stirng) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}
