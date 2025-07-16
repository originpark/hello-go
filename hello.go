package hello_go

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is empty")
	}
	message := fmt.Sprintf("hello twice, %v, weclome!", name)
	return message, nil
}

func RandomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Once you're looking for, %v.",
		"Great to see you how are you?",
	}

	return formats[rand.Intn(len(formats))]
}
