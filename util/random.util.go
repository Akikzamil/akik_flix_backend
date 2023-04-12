package util

import "crypto/rand"

const chars = "1234567890"

func GetSixDigitRandomNumber() (string, error) {
	buffer := make([]byte, 6)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	charsLength := len(chars)
	for i := 0; i < 6; i++ {
		buffer[i] = chars[int(buffer[i])%charsLength]
	}

	return string(buffer), nil
}