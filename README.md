## Go JSON Parser

This repository contains a JSON parser written in Go, created as part of a learning exercise to understand parsing techniques.

### What I Learned

Building this JSON parser provided valuable insights into the fundamental concepts of parsing, including:

- **Lexical Analysis**: Understanding how to break down a sequence of characters into meaningful tokens.
- **Syntactic Analysis**: Learning to analyze a list of tokens to match a formal grammar.
- **Error Handling**: Developing robust error handling to provide useful feedback on invalid JSON structures.

### Challenges Faced

Throughout the development process, several challenges were encountered and overcome:

- **Tokenization**: Designing an efficient lexer to correctly identify and categorize different parts of JSON data.
- **Grammar Matching**: Implementing a parser that accurately interprets the structure of JSON according to its formal definition.
- **Data Types Handling**: Ensuring correct parsing and representation of various JSON data types, including strings, numbers, booleans, null values, objects, and arrays.
- **Nested Structures**: Handling complex nested JSON objects and arrays required careful consideration of recursive parsing techniques.
- **Validation and Testing**: Creating extensive test cases to validate the parser's accuracy and robustness against both valid and invalid JSON inputs.

### Conclusion

This project was an excellent opportunity to delve into the world of parsing and compilers. The knowledge gained from building a JSON parser from scratch is foundational and applicable to more advanced programming challenges, such as developing custom data formats and building compilers.


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
