package grettings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Greet(name string) (string, error) {
	if name == "" || name == " " {
		//return "", fmt.Errorf("empty name")
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Greets(names []string) (map[string]string, error) {
	//messages := make(map[string]string)
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Greet(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	formats := [] string {
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
