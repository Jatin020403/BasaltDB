package utils

import (
	"fmt"

	murmur "github.com/twmb/murmur3"
)

func MurmurHashInt(key string) uint64{
	return murmur.StringSum64(key)
}

func MurmurHash(key string) string {
	hash := fmt.Sprint(MurmurHashInt(key))

	for len(key) < 5{
		key += "0"
	}
	key = key[:5]
	return hash + key
}
