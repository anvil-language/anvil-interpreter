package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/anvil-language/anvil-interpreter/pkg/evaluator"
	"github.com/anvil-language/anvil-interpreter/pkg/lexer"
	"github.com/anvil-language/anvil-interpreter/pkg/object"
	"github.com/anvil-language/anvil-interpreter/pkg/parser"
)

func main() {
	fmt.Println("Anvil Programming Language Engine (v0.1.0-alpha)")
	fmt.Println("Type Anvil code below to evaluate statements. Press Ctrl+C to exit.\n")

	env := object.NewEnvironment()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("anvil > ")
		if !scanner.Scan() {
			return
		}

		line := scanner.Text()
		if line == "" {
			continue
		}

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			for _, msg := range p.Errors() {
				fmt.Printf("Parser Error: %s\n", msg)
			}
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			fmt.Println(evaluated.Inspect())
		}
	}
}
