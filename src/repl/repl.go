// Copyright (c) 2018 meritozh
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package repl

import (
	"bufio"
	"fmt"
	"io"
	"lexer"
	"parser"
)

// PROMPT prompt symbol
const PROMPT = ">> "

// Start input stream reading function
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
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

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
