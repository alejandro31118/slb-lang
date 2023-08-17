package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/alejandro31118/slb-lang/lexer"
	"github.com/alejandro31118/slb-lang/token"
)

const INIT_PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Println(INIT_PROMPT)

		if !scanner.Scan() {
			return
		}

		line := scanner.Text()
		lex := lexer.New(line)

		for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
