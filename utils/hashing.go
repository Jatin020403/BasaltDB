package utils

import (
	"fmt"

	murmur "github.com/twmb/murmur3"
)

func MurmurHash(key string) string {
	hash := fmt.Sprint(murmur.StringSum64(key))

	for len(key) < 5{
		key += "0"
	}
	key = key[:5]
	return hash + key
}
