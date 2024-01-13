package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkeyInterpreter/src/lexer"
	"monkeyInterpreter/src/token"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const PROMPT = " $ "

func Start(in io.Reader, out io.Writer, userName string) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(userName + PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()

		// for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		// 	fmt.Printf("%+v\n", tok)
		// }
		if line == "clear" {
			clearTerminal()
			continue
		}

		lexedTokens := lexer.New(line)

		tok := lexedTokens.NextToken()

		for tok.Type != token.EOF {
			fmt.Printf("%+v\n", beautifyTokensOutput(tok))
			tok = lexedTokens.NextToken()
		}
	}
}

func beautifyTokensOutput(tok token.Token) string {
	const PADDING_RIGHT = 8

	quotedValue := strconv.Quote(tok.Value)

	return fmt.Sprintf("Type: %s, Value: %s, Column: %s",
		padRightSmallWord(string(tok.Type), PADDING_RIGHT),
		padRightSmallWord(quotedValue, PADDING_RIGHT),
		padRightSmallWord(strconv.Itoa(tok.Column), PADDING_RIGHT))
}

func padRightSmallWord(s string, padding int) string {
	if len(s) > padding {
		return s
	}

	padTotal := padding - len(s)
	pad := ""

	for i := 0; i < padTotal; i++ {
		pad += " "
	}

	return s + pad

}

func clearTerminal() {
	var cmd *exec.Cmd
	shell := os.Getenv("SHELL")
	if strings.Contains(shell, "bash") {
		cmd = exec.Command("clear")
	} else if strings.Contains(strings.ToLower(shell), "cmd.exe") {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		// Default to clear if the shell type cannot be determined
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
