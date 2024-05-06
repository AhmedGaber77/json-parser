# JSON Parser

This is a JSON parser written in Go.

## Structure

The project is structured as follows:

- `cmd/json/json-parser.go`: This is where the [`Unmarshal`](cmd/json/json-parser.go) function is located.
- `pkg/lexer/lexer.go`: This file contains the lexer logic.
- `pkg/lexer/lexer_test.go`: This file contains the tests for the lexer, including [`TestLex`](pkg/lexer/lexer_test.go).
- `pkg/parser/parser.go`: This file contains the parser logic.
- `pkg/parser/parser_test.go`: This file contains the tests for the parser, including [`TestParse`](pkg/parser/parser_test.go).

## Usage

To use the JSON parser, you need to call the `Unmarshal` function from the `json-parser.go` file in the `cmd/json` directory.

## Testing

To run the tests for the lexer and the parser, navigate to the `pkg/lexer` and `pkg/parser` directories respectively and run `go test`.
