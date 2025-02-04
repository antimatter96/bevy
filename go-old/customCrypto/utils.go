package customcrypto

import "math/rand"

var letterRunes = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// GenerateRandomString is used to generate a string of given length
// String will follow the pattern [0-9a-zA-Z]+
func GenerateRandomString(length int) (string, error) {
	arr := make([]byte, length)
	for i := range arr {
		arr[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(arr), nil
}
