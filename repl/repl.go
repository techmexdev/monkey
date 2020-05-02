package repl

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"monkey/lexer"
	"monkey/token"
)

const prompt = ">> "

// Start reapetedly scans in, and prints tokens in its text.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(prompt)
		if !scanner.Scan() {
			return
		}

		txt := scanner.Text()
		l := lexer.New(txt)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			log.Printf("%+#v\n", tok)
		}
	}
}
