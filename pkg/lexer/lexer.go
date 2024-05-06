package lexer

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

// Tokens represents a slice of any type.
type Tokens []any

// JSON_WHITESPATE represents the whitespace characters in JSON.
var JSON_WHITESPATE = []rune{' ', '\t', '\r', '\n'}

// JSON_SYNTAX represents the syntax characters in JSON.
var JSON_SYNTAX = []rune{JSON_COMMA, JSON_COLON, JSON_LEFTBRACKET, JSON_RIGHTBRACKET,
	JSON_LEFTBRACE, JSON_RIGHTBRACE}

const (
	JSON_QUOTE        = '"'     // JSON_QUOTE represents the double quote character in JSON.
	JSON_COMMA        = ','     // JSON_COMMA represents the comma character in JSON.
	JSON_COLON        = ':'     // JSON_COLON represents the colon character in JSON.
	JSON_LEFTBRACKET  = '['     // JSON_LEFTBRACKET represents the left square bracket character in JSON.
	JSON_RIGHTBRACKET = ']'     // JSON_RIGHTBRACKET represents the right square bracket character in JSON.
	JSON_LEFTBRACE    = '{'     // JSON_LEFTBRACE represents the left curly brace character in JSON.
	JSON_RIGHTBRACE   = '}'     // JSON_RIGHTBRACE represents the right curly brace character in JSON.
	JSON_TRUE         = "true"  // JSON_TRUE represents the true boolean value in JSON.
	JSON_FALSE        = "false" // JSON_FALSE represents the false boolean value in JSON.
	JSON_NULL         = "null"  // JSON_NULL represents the null value in JSON.
)

// lexString lexically analyzes a JSON string.
func lexString(payload string) (string, string, error) {
	var builder strings.Builder

	if payload[0] != JSON_QUOTE {
		return "", payload, fmt.Errorf("payload is not a string")
	}
	payload = payload[1:]
	for idx, c := range payload {
		if c == JSON_QUOTE {
			return builder.String(), payload[idx+1:], nil
		}
		builder.WriteRune(c)
	}
	return "", payload, fmt.Errorf("expected end-of-string quote")
}

// lexNumber lexically analyzes a JSON number.
func lexNumber(payload string) (int, string, error) {
	var builder strings.Builder

	for _, c := range payload {
		if unicode.IsDigit(c) {
			builder.WriteRune(c)
			continue
		}
		break
	}
	remaining := payload[builder.Len():]
	if builder.Len() == 0 {
		return 0, payload, fmt.Errorf("payload is not a number")
	}
	jsonInt, _ := strconv.Atoi(builder.String())

	return jsonInt, remaining, nil

}

// lexBool lexically analyzes a JSON boolean.
func lexBool(payload string) (bool, string, error) {
	if strings.HasPrefix(payload, JSON_TRUE) {
		return true, payload[len(JSON_TRUE):], nil
	} else if strings.HasPrefix(payload, JSON_FALSE) {
		return false, payload[len(JSON_FALSE):], nil
	}
	return false, payload, fmt.Errorf("payload is not a boolean")
}

// lexNull lexically analyzes a JSON null value.
func lexNull(payload string) (any, string, error) {
	if strings.HasPrefix(payload, JSON_NULL) {
		return nil, payload[len(JSON_NULL):], nil
	}
	return nil, payload, fmt.Errorf("payload is not null")
}

// Lex lexically analyzes a JSON payload and returns a slice of tokens.
func Lex(payload string) (Tokens, error) {
	tokens := Tokens{}

	for len(payload) > 0 {

		if jsonStr, remaining, err := lexString(payload); err == nil {
			tokens = append(tokens, jsonStr)
			payload = remaining
			continue
		}

		if jsonNum, remaining, err := lexNumber(payload); err == nil {
			tokens = append(tokens, jsonNum)
			payload = remaining
			continue
		}

		if jsonBool, remaining, err := lexBool(payload); err == nil {
			tokens = append(tokens, jsonBool)
			payload = remaining
			continue
		}

		if _, remaining, err := lexNull(payload); err == nil {
			tokens = append(tokens, nil)
			payload = remaining
			continue
		}

		if slices.Contains(JSON_WHITESPATE, rune(payload[0])) {
			payload = payload[1:]
			continue
		}

		if slices.Contains(JSON_SYNTAX, rune(payload[0])) {
			tokens = append(tokens, rune(payload[0]))
			payload = payload[1:]
			continue
		}

		return nil, fmt.Errorf("unexpected character: %d", payload[0])
	}

	return tokens, nil
}
