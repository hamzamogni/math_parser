package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"parser_lexer/evaluator"
	"parser_lexer/lexer"
	"parser_lexer/parser"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Printf("Type your comnmands\n")
		StartREPL(os.Stdin, os.Stdout)
	} else if len(os.Args) == 2 {
        out := os.Stdout
		l := lexer.New(os.Args[1])
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			return
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		} else {
			io.WriteString(out, "nil\n")
		}
	} else {
		fmt.Printf("Usage: ./interpreter or ./interpreter 'expression'\n")
	}

}

func StartREPL(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(">>> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		} else {
			io.WriteString(out, "nil\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
