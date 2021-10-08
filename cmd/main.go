package main

import (
	"fmt"
	"os"

	"github.com/mmeinzer/glox/scan"
)

func main() {
	args := os.Args[1:]
	var err error
	if len(args) > 1 {
		fmt.Println("Usage: jlox [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		err = runFile(args[0])
	} else {
		err = runPrompt()
	}

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func runFile(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = run(string(bytes))
	return err
}

func runPrompt() error {
	panic("not implemented")
}

func run(source string) error {
	scanner := scan.NewScanner(source)
	tokens := scanner.ScanTokens()

	for _, t := range tokens {
		fmt.Println(t)
	}

	return nil
}