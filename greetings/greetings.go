package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
  if name == "" {
    return "", errors.New("Empty Name, btet5awwath?")
  }

  message := fmt.Sprintf(GetRandomFormat(), name)

  return message, nil;
}

func Hellos(names []string) (map[string]string, error) {
  messages := make(map[string]string)

  for _, name := range names {
    message, err := Hello(name)

    if err != nil {
      return nil, err
    }

    messages[name] = message
  }

  return messages, nil;
}

func GetRandomFormat() string  {
  formats := []string {
    "Greetings, %v",
    "Sho m3allem %v?",
    "%v Ahlan wa Sahlan",
 }

  return formats[rand.Intn(len(formats))]
}
