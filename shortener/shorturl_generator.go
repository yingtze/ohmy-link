package shortener

import (
	"crypto/sha256"
	"fmt"
	"github.com/itchyny/base58-go"
	"os"
)

func sha2560f(input stirng) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {

}
