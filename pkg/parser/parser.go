package parser

import (
	"fmt"

	"github.com/AhmedGaber77/json-parser/pkg/lexer"
)

// parseArray parses a JSON array from the given tokens and returns the parsed array,
// the remaining tokens after parsing, and any error encountered during parsing.
func parseArray(tokens lexer.Tokens) ([]any, lexer.Tokens, error) {
	jsonArr := []any{}

	if len(tokens) == 0 || tokens[0] == lexer.JSON_RIGHTBRACKET {
		return jsonArr, tokens[1:], nil
	}

	for len(tokens) > 0 {
		var json any
		var remaining lexer.Tokens
		var err error

		json, remaining, err = Parse(tokens, false)
		if err != nil {
			return nil, remaining, err
		}
		jsonArr = append(jsonArr, json)

		if len(remaining) == 0 {
			return nil, tokens, fmt.Errorf("expected value after comma in array")
		}

		if remaining[0] == lexer.JSON_RIGHTBRACKET {
			return jsonArr, remaining[1:], nil
		}

		if remaining[0] != lexer.JSON_COMMA {
			return nil, tokens, fmt.Errorf("expected comma after object in array")
		}

		tokens = remaining[1:]
	}

	return nil, tokens, fmt.Errorf("expected end-of-array bracket")
}

// parseObject parses a JSON object from the given tokens and returns the parsed object,
// the remaining tokens after parsing, and any error encountered during parsing.
func parseObject(tokens lexer.Tokens) (map[string]any, lexer.Tokens, error) {
	jsonObj := map[string]any{}

	if len(tokens) == 0 || tokens[0] == lexer.JSON_RIGHTBRACE {
		return jsonObj, tokens[1:], nil
	}

	for len(tokens) > 0 {
		key, ok := tokens[0].(string)
		if !ok {
			return nil, tokens, fmt.Errorf("expected string key, got: %v", tokens[0])
		}
		tokens = tokens[1:]

		if len(tokens) == 0 || tokens[0] != lexer.JSON_COLON {
			return nil, tokens, fmt.Errorf("expected colon after key in object, got: %v", tokens[0])
		}

		value, remaining, err := Parse(tokens[1:], false)
		if err != nil {
			return nil, tokens, err
		}
		jsonObj[key] = value

		if len(remaining) == 0 {
			return nil, tokens, fmt.Errorf("expected value after colon in object")
		}

		if remaining[0] == lexer.JSON_RIGHTBRACE {
			return jsonObj, remaining[1:], nil
		}

		if remaining[0] != lexer.JSON_COMMA {
			return nil, tokens, fmt.Errorf("expected comma after object in object")
		}

		tokens = remaining[1:]
	}

	return nil, tokens, fmt.Errorf("expected end-of-object bracket")
}

// Parse parses a sequence of tokens and returns the corresponding JSON value,
// along with the remaining tokens and any error encountered during parsing.
// The `isRoot` parameter indicates whether the parsed value is the root of the JSON document.
// If `isRoot` is true, the parsed value must be an object.
// If `isRoot` is false, the parsed value can be any valid JSON value.
// The function returns the parsed value, the remaining tokens after parsing,
// and an error if any occurred during parsing.
func Parse(tokens lexer.Tokens, isRoot bool) (any, lexer.Tokens, error) {
	if len(tokens) == 0 {
		return nil, tokens, fmt.Errorf("empty tokens")
	}

	token := tokens[0]

	if isRoot && token != lexer.JSON_LEFTBRACE {
		return nil, tokens, fmt.Errorf("root must be an object")
	}

	switch token {
	case lexer.JSON_LEFTBRACKET:
		return parseArray(tokens[1:])
	case lexer.JSON_LEFTBRACE:
		return parseObject(tokens[1:])
	default:
		return token, tokens[1:], nil
	}
}
