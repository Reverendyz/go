package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name given")
	}
	message := fmt.Sprintf("Hi, %v ", name)
	return message, nil
}
