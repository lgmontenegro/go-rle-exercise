package main

import (
	"fmt"
	"strconv"
)

// Message portugal, caralho!
type Message struct {
	Phrase        string
	EncodedPhrase string
}

type toDecode struct {
	multiplier int
	message    string
}

func main() {
	msg := Message{
		Phrase: "AAAAAAAAAA",
	}

	fmt.Println(msg)
	fmt.Println(msg.encode())

	fmt.Println(msg)
	fmt.Println(msg.decode())
}

func (m *Message) decode() string {
	var multiplier string
	decodeCollection := make([]toDecode, 0)

	for _, char := range m.EncodedPhrase {
		if char >= 48 && char <= 57 {
			multiplier = multiplier + string(char)
			continue
		}

		x, _ := strconv.Atoi(multiplier)
		decodeCollection = append(decodeCollection, toDecode{x, string(char)})
		multiplier = "0"
	}

	m.Phrase = "#"
	for _, decodeInstruction := range decodeCollection {
		for i := int(0); i < decodeInstruction.multiplier; i++ {
			m.Phrase += decodeInstruction.message
		}
	}

	return m.Phrase
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
