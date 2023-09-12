package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jar-b/jsondup"
)

func init() {
	// slightly better usage output
	flag.Usage = func() {
		fmt.Fprint(flag.CommandLine.Output(), "Detect duplicate keys in a JSON object.\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [filename]\n", os.Args[0])
	}
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatal("unexpected number of arguments")
	}

	f := flag.Arg(0)
	b, err := os.ReadFile(f)
	if err != nil {
		log.Fatalf("reading file: %s", err)
	}

	if err := jsondup.ValidateNoDuplicateKeys(string(b)); err != nil {
		log.Fatal(err)
	}
}
