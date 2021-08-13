package utils

import (
	"crypto/sha1"
)

func Hash(s string) []byte {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return bs
}
