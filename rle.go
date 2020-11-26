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
	multiplier int32
	message    string
}

func main() {
	msg := Message{
		Phrase: "AAAAAAAAAABBB",
	}

	fmt.Println(msg)
	fmt.Println(msg.encode())

	fmt.Println(msg)
	fmt.Println(msg.decode())
}

func (m *Message) decode() string {
	var multiplier int32
	decodeCollection := make([]toDecode, 0)

	for _, char := range m.EncodedPhrase {
		if char >= 48 && char <= 57 {
			realIntChar := char - 48
			multiplier = (multiplier * 10) + realIntChar
			continue
		}
		decodeCollection = append(decodeCollection, toDecode{multiplier, string(char)})
		multiplier = 0
	}

	m.Phrase = "#"
	for _, decodeInstruction := range decodeCollection {
		for i := int32(0); i < decodeInstruction.multiplier; i++ {
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
