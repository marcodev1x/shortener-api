package helpers

import (
	"crypto/sha256"
	"fmt"
)

const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func base62Encode(data []byte) string {
	var num uint64
	for _, b := range data {
		num = num*256 + uint64(b)
	}

	if num == 0 {
		return string(base62Chars[0])
	}

	result := ""
	for num > 0 {
		result = string(base62Chars[num%62]) + result
		num /= 62
	}

	return result
}

func GenerateHash(id int, seed string) string {
	var input string

	if id == 0 {
		if seed == "" {
			return ""
		}
		input = seed
	} else {
		input = fmt.Sprintf("%d:%s", id, seed)
	}

	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashBytes := hasher.Sum(nil)

	return base62Encode(hashBytes[:6])
}
