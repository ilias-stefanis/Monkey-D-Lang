package parser

import (
	"monkeyInterpreter/src/ast"
	"monkeyInterpreter/src/token"
)

func (parser *Parser) parseStatement() ast.Statement {
	switch parser.curToken.Type {
	case token.LET:
		return parser.parseLetStatement()
	case token.CONST:
		return parser.parseConstStatement()
	case token.ENUM:
		return parser.parseEnumStatement()
	default:
		return nil
	}
}

// to remember what to fix later
func TODO(message string) {
	panic("TODO (" + message + ")")
}

func (parser *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{
		Token: parser.curToken,
	}

	if parser.peekIsNot(token.IDENT) {
		return nil
	}

	statement.Name = &ast.Identifier{
		Token: parser.curToken,
		Value: parser.curToken.Value,
	}

	if parser.peekIsNot(token.ASSIGN) {
		return nil
	}

	for !parser.curTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return statement
}

func (parser *Parser) parseConstStatement() *ast.ConstStatement {
	statement := &ast.ConstStatement{
		Token: parser.curToken,
	}

	if parser.peekIsNot(token.IDENT) {
		return nil
	}

	statement.Name = &ast.Identifier{
		Token: parser.curToken,
		Value: parser.curToken.Value,
	}

	if parser.peekIsNot(token.ASSIGN) {
		return nil
	}

	for !parser.curTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return statement
}

// example enum color = red|green|blue
func (parser *Parser) parseEnumStatement() *ast.EnumStatement {
	statement := &ast.EnumStatement{
		Token: parser.curToken,
	}

	// Parse enum name
	if parser.peekIsNot(token.IDENT) {
		return nil
	}

	statement.Name = &ast.Identifier{
		Token: parser.curToken,
		Value: parser.curToken.Value,
	}

	if parser.peekIsNot(token.ASSIGN) {
		return nil
	}

	enumVariants := []ast.Expression{}

	if parser.peekIsNot(token.IDENT) {
		return nil
	}

	enumVariants = append(enumVariants, &ast.Identifier{
		Token: parser.curToken,
		Value: parser.curToken.Value,
	})
	sequence, err := parser.getRepeatingSequence([]token.TokenType{token.IDENT, token.PIPE}, token.SEMICOLON)

	if err != nil {
		TODO("getRepeatingSequence failed, message: " + err.Error())
	}

	for _, token := range sequence {
		enumVariants = append(enumVariants, &ast.Identifier{
			Token: token,
			Value: token.Value,
		})
	}
	// for !parser.curTokenIs(token.SEMICOLON) {
	// 	parser.nextToken()
	// }
	TODO("decide how to structure enum variants in ast.go")
	return statement
}
