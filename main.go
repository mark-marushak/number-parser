package main

import (
	"flag"
	"fmt"
	"github.com/mark-marushak/number-parser/parser"
	"github.com/mark-marushak/number-parser/repository"
	"os"
)

/**
"1;4-7;9"
*/

type StandardLogger struct{}

func (s StandardLogger) Info(args ...interface{}) {
	_, err := fmt.Fprintln(os.Stdout, args)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error while trying to log info: %v", args))
	}
}

func (s StandardLogger) Error(args ...interface{}) {
	_, err := fmt.Fprintln(os.Stderr, args)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error while trying to log error: %v", args))
	}
}

func main() {
	var input *string
	input = flag.String("input", "1;4-7;9", "enter string with numbers by pattern '1;4-7;9' or '1;2;3;5;6'")

	service := parser.NewParser(repository.NewParser(), &StandardLogger{})
	fmt.Println(service.Parse(*input))
}
