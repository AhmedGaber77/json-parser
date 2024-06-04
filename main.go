package main

import (
	"fmt"
	"log"

	jsonparser "github.com/AhmedGaber77/json-parser/cmd/json"
)

func main() {
	jsonstr := `{"name": "Ahmed", "age": null}`
	jsonParsed := map[string]interface{}{}
	err := jsonparser.Unmarshal([]byte(jsonstr), &jsonParsed)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", jsonParsed)
}
