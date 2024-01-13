package parser

import (
	"Monkey-D-Lang/src/ast"
	"Monkey-D-Lang/src/lexer"
	"Monkey-D-Lang/src/token"
	"fmt"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
	}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	// fmt.Println("nextToken")
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
	// fmt.Println("curToken", p.curToken)
	// fmt.Println("peekToken", p.peekToken)
}

func (parser *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for parser.curToken.Type != token.EOF {
		currentStatement := parser.parseStatement()
		if currentStatement != nil {
			program.Statements = append(program.Statements, currentStatement)
		}
		parser.nextToken()
	}

	return program
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// like above, but continues to next token
func (parser *Parser) peekIs(tokenToCheck token.TokenType) bool {

	if parser.peekTokenIs(tokenToCheck) {
		parser.nextToken()
		return true
	}
	return false

}

func (parser *Parser) peekIsNot(t token.TokenType) bool {
	return !parser.peekIs(t)
}

func (parser *Parser) getRepeatingSequence(
	repeater []token.TokenType,
	stopToken token.TokenType,
) ([]token.Token, error) {
	var repeatingSequence []token.Token

	for {
		if parser.curTokenIs(stopToken) {
			break
		}

		for _, tokenType := range repeater {
			if parser.curTokenIs(tokenType) {
				repeatingSequence = append(repeatingSequence, parser.curToken)
			} else {
				return repeatingSequence,
					fmt.Errorf(`expected "%s", got "%s",`, tokenType, parser.curToken.Type)
				// apparently a compiler warning bruh???
				//errors.New(fmt.Sprintf(`Expected "%s", got "%s"`, tokenType, parser.curToken.Type))
			}
		}

		parser.nextToken()
	}

	return repeatingSequence, nil
}
