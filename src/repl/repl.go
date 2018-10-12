package repl

import (
	"bufio"
	"fmt"
	"io"
	"lexer"
	"token"
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

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
