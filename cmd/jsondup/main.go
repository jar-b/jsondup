package main

import (
	"log"

	"github.com/jar-b/jsondup"
)

func main() {
	s := `{"a":"foo", "a":"bar"}`
	if err := jsondup.ValidateNoDuplicateKeys(s); err != nil {
		log.Fatal(err)
	}
}
