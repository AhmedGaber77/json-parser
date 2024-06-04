package jsonparser

import (
	"fmt"

	"github.com/AhmedGaber77/json-parser/pkg/lexer"
	"github.com/AhmedGaber77/json-parser/pkg/parser"
)

// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
// The data parameter is the JSON-encoded data to be parsed.
// The v parameter is a pointer to a map[string]interface{} where the parsed JSON will be stored.
// Unmarshal returns an error if there was an issue parsing the JSON.
func Unmarshal(data []byte, v *map[string]interface{}) error {
	tokens, err := lexer.Lex(string(data))
	if err != nil {
		return err
	}
	jsonParsed, _, err := parser.Parse(tokens, true)
	if err != nil {
		return err
	}
	var ok bool
	*v, ok = jsonParsed.(map[string]interface{})
	if !ok {
		return fmt.Errorf("error parsing JSON: %v", jsonParsed)
	}
	return nil
}
