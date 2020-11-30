package main

// implementing RLE algoritm
//https://en.wikipedia.org/wiki/Run-length_encoding

import (
	"fmt"
	"strconv"
	"strings"
)

// Message portugal, caralho!
type Message struct {
	Phrase        string
	EncodedPhrase string
}

func main() {
	msg := Message{
		Phrase: "AAAAAAAAAABBB",
	}

	fmt.Println(msg)
	fmt.Println(msg.encode())

	fmt.Println(msg)
	decodedMsg, err := msg.decode()
	if err != nil {
		println(err)
	}
	fmt.Println(decodedMsg)
}

func (m *Message) decode() (string, error) {
	var multiplier int32
	var decodedString strings.Builder

	for _, char := range m.EncodedPhrase {
		if char >= 48 && char <= 57 {
			realIntChar := char - 48
			multiplier = (multiplier * 10) + realIntChar
			continue
		}

		for i := multiplier; i >= 1; i-- {
			_, err := decodedString.WriteRune(char)
			if err != nil {
				return "", err
			}
		}
		multiplier = 0
	}

	m.Phrase = decodedString.String()
	return m.Phrase, nil
}

func (m *Message) encode() string {
	returnChar := ""
	prevChar := ""
	countChar := 0
	for i, char := range m.Phrase {
		if prevChar == "" {
			prevChar = string(char)
			countChar++
			continue
		}
		if prevChar == string(char) {
			countChar++
			if i == len(m.Phrase)-1 {
				returnChar += strconv.Itoa(countChar) + string(prevChar)
			}
			continue
		}
		returnChar += strconv.Itoa(countChar) + string(prevChar)

		prevChar = string(char)
		countChar = 1
		if i == len(m.Phrase)-1 {
			returnChar += strconv.Itoa(countChar) + string(prevChar)
		}
	}

	m.EncodedPhrase = returnChar
	return returnChar
}
