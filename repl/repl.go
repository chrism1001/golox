package repl

import (
	"bufio"
	"fmt"
	"golox/lexer"
	"io"
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
		for _, tt := range l.GetTokens() {
			fmt.Printf("%+v\n", tt)
		}
	}
}
