package lexer

import (
	"Monkey-D-Lang/src/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
	line         int  // current line in input
	column       int  // current column in input
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.line = 1
	l.column = 0
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
		l.column += 1
	}
	l.position = l.readPosition

	l.readPosition += 1
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		if l.ch == '\n' {
			l.line += 1
			l.column = 0
		}
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()
	// l.readChar()
	// fmt.Println(l.ch)
	currentLine := l.line
	currentColumn := l.column

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			// startPosition := l.readPosition - 1

			tok = newToken(token.EQ, ch, currentColumn, currentLine)
			tok.Value = "=="

		} else if l.peekChar() == '>' {
			ch := l.ch
			l.readChar()
			// startPosition := l.readPosition - 1

			tok = newToken(token.ARROW, ch, currentColumn, currentLine)
			tok.Value = "=>"

		} else {
			tok = newToken(token.ASSIGN, l.ch, currentColumn, currentLine)

		}
	case '+':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			// startPosition := l.readPosition - 1

			tok = newToken(token.PLUS_EQ, ch, currentColumn, currentLine)
			tok.Value = "+="

		} else {
			tok = newToken(token.PLUS, l.ch, currentColumn, currentLine)
		}

	case +'-':
		if l.peekChar() == '>' {
			ch := l.ch
			l.readChar()
			// startPosition := l.readPosition - 1

			tok = newToken(token.ARROW_SMALL, ch, currentColumn, currentLine)
			tok.Value = "->"

		} else if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			// startPosition := l.readPosition - 1

			tok = newToken(token.MINUS_EQ, ch, currentColumn, currentLine)
			tok.Value = "-="

		} else {
			tok = newToken(token.MINUS, l.ch, currentColumn, currentLine)
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			// startPosition := l.readPosition - 1

			tok = newToken(token.NOT_EQ, ch, currentColumn, currentLine)
			tok.Value = "!="

		} else {
			tok = newToken(token.BANG, l.ch, currentColumn, currentLine)

		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch, currentColumn, currentLine)
	case '/':
		tok = newToken(token.SLASH, l.ch, currentColumn, currentLine)
	case '<':
		tok = newToken(token.LT, l.ch, currentColumn, currentLine)
	case '>':
		if l.peekChar() == '>' {
			ch := l.ch
			l.readChar()
			// startPosition := l.readPosition - 1

			tok = newToken(token.PIPE, ch, currentColumn, currentLine)
			tok.Value = ">>"

		} else {
			tok = newToken(token.GT, l.ch, currentColumn, currentLine)
		}
	case '|':
		tok = newToken(token.UNION, l.ch, currentColumn, currentLine)
	// case '%':
	// 	tok = newToken(token.MODULO, l.ch, currentColumn)
	// case '&':
	// 	tok = newToken(token.AND, l.ch, currentColumn)
	// case '^':
	// 	tok = newToken(token.XOR, l.ch, currentColumn)
	// case '~':
	// 	tok = newToken(token.NOT, l.ch, currentColumn)
	case ',':
		tok = newToken(token.COMMA, l.ch, currentColumn, currentLine)
	case '.':
		tok = newToken(token.DOT, l.ch, currentColumn, currentLine)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch, currentColumn, currentLine)
	case '(':
		tok = newToken(token.LPAREN, l.ch, currentColumn, currentLine)
	case ')':
		tok = newToken(token.RPAREN, l.ch, currentColumn, currentLine)
	case '{':
		tok = newToken(token.LBRACE, l.ch, currentColumn, currentLine)
	case '}':
		tok = newToken(token.RBRACE, l.ch, currentColumn, currentLine)
	case '[':
		tok = newToken(token.LBRACKET, l.ch, currentColumn, currentLine)
	case ']':
		tok = newToken(token.RBRACKET, l.ch, currentColumn, currentLine)
	case 0:
		tok.Value = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			// startPosition := l.readPosition
			// tok.Value = l.readIdentifier()
			// tok.Type = token.LookupIdent(tok.Value)

			// return newTokenString(tok.Type, tok.Value, startPosition)

			tok.Column = l.readPosition
			tok.Value = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Value)
			tok.Line = l.line
			tok.Column = currentColumn
			return tok

		} else if isDigit(l.ch) {
			tok.Column = l.readPosition
			tok.Type = token.INT
			tok.Value = l.readNumber()
			tok.Line = l.line
			tok.Column = currentColumn
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch, currentColumn, l.line)
		}

	}
	l.readChar()

	return tok
}

func (l *Lexer) readIdentifier() string {
	posistion := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[posistion:l.position]
}

// isLetter returns true if the given byte is a letter or underscore.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenToInsert token.TokenType, ch byte, column int, line int) token.Token {
	return token.Token{
		Type:     tokenToInsert,
		Value:    string(ch),
		FileName: "todo logic for this",
		Line:     line,
		Column:   column,
	}
}

// func newTokenString(tokenToInsert token.TokenType, ch string, column int) token.Token {
// 	return token.Token{
// 		Type:     tokenToInsert,
// 		Value:    ch,
// 		FileName: "todo logic for this",
// 		Line:     1,
// 		Column:   column,
// 	}
// }
