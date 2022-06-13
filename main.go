package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func Perform(args Arguments, writer io.Writer) error {
	operation, foundOp := args["operation"]
	if !foundOp || operation == defaultArgValue {
		return fmt.Errorf("-operation flag has to be specified")
	}
	filename, foundFn := args["fileName"]
	if !foundFn || filename == defaultArgValue {
		return fmt.Errorf("-fileName flag has to be specified")
	}

	switch operation {
	case "add":
		item, foundItem := args["item"]
		if !foundItem || item == defaultArgValue {
			return fmt.Errorf("-item flag has to be specified")
		}

		return Add(item, filename, writer)
	case "list":
		return List(filename, writer)
	case "findById":
		id, foundId := args["id"]
		if !foundId || id == defaultArgValue {
			return fmt.Errorf("-id flag has to be specified")
		}

		return findById(id, filename, writer)
	case "remove":
		id, foundId := args["id"]
		if !foundId || id == defaultArgValue {
			return fmt.Errorf("-id flag has to be specified")
		}

		return Remove(id, filename, writer)
	default:
		return fmt.Errorf("Operation %s not allowed!", args["operation"])
	}
}

func parseArgs() Arguments {
	flag.Parse()

	args := Arguments{
		"operation": argOperation,
		"fileName":  argFilename,
		"item":      argItem,
		"id":        argId,
	}

	return args
}

const (
	defaultArgValue = ""
)

var argOperation string
var argFilename string
var argItem string
var argId string

func init() {
	flag.StringVar(&argOperation, "operation", defaultArgValue, "operation for processing ('add', 'list', 'findById', 'remove')")
	flag.StringVar(&argFilename, "fileName", defaultArgValue, "data filename")
	flag.StringVar(&argItem, "item", defaultArgValue, "item info (used in 'add' operation)")
	flag.StringVar(&argId, "id", defaultArgValue, "element id (used in 'remove'/'findById' operation)")
}

func main() {
	// go run main -operation "add" -fileName "users.json" -item "{\"id\": \"1\", \"email\": \"email@test.com\", \"age\": 23}"

	// go run main -operation "list" -fileName "users.json"

	// go run main -operation "remove" -fileName "users.json -id 1"

	// go run main -operation "findById" -fileName "users.json -id 1"

	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}
