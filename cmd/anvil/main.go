package main

import (
	"fmt"
	"os"

	"github.com/anvil-language/anvil-interpreter/pkg/evaluator"
	interface_repl "github.com/anvil-language/anvil-interpreter/pkg/interface"
	"github.com/anvil-language/anvil-interpreter/pkg/lexer"
	"github.com/anvil-language/anvil-interpreter/pkg/object"
	"github.com/anvil-language/anvil-interpreter/pkg/parser"
)

func main() {
	// If a file argument is passed: run the file script directly
	if len(os.Args) > 1 {
		filename := os.Args[1]
		bytes, err := os.ReadFile(filename)
		if err != nil {
			fmt.Printf("Error opening file %s: %s\n", filename, err
			os.Exit(1)
		}

		env := object.NewEnvironment()
		l := lexer.New(string(bytes))
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			for _, msg := range p.Errors() {
				fmt.Printf("Parser Error: %s\n", msg)
			}
			os.Exit(1)
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			fmt.Println(evaluated.Inspect())
		}
		return
	}

	// Default fallback: Launch REPL
	interface_repl.Start(os.Stdin, os.Stdout)
}
