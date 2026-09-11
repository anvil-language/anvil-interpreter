package interface_repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/anvil-language/anvil-interpreter/pkg/evaluator"
	"github.com/anvil-language/anvil-interpreter/pkg/lexer"
	"github.com/anvil-language/anvil-interpreter/pkg/object"
	"github.com/anvil-language/anvil-interpreter/pkg/parser"
)

const PROMPT = "anvil > "

func Start(in io.Reader, out io.Writer) {
	fmt.Println("Anvil Programming Language Engine (v0.1.0-alpha)")
	fmt.Println("Type Anvil code below to evaluate statements. Press Ctrl+C to exit.\n")

	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
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
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			fmt.Fprintln(out, evaluated.Inspect())
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		fmt.Fprintf(out, "Parser Error: %s\n", msg)
	}
}
