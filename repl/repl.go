package repl

import (
	"bufio"
	"fmt"
	"io"
	"rainbow/lexer"
	"rainbow/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("Type \"%s\", Literal \"%s\"\n", tok.Type, tok.Literal)
		}
	}
}
