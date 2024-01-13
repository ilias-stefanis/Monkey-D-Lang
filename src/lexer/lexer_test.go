package lexer

import (
	"Monkey-D-Lang/src/token"
	"fmt"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
	let five = 5;
	let ten = 10;
	let add = func(x, y) {
	x + y;
	};
	let result = add(five, ten);
	!-/*5;
	5 < 10 > 5;
	if ((5 < 10) and (10 > 1)) {
		return not (five or ten);
	}
	10 == 10;
	10 != 9;
	`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "func"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.AND, "and"},
		{token.LPAREN, "("},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "1"},
		{token.RPAREN, ")"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.NOT, "not"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.OR, "or"},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},

		{token.EOF, ""},
	}

	l := New(input)
	// l.readChar()
	for i, tt := range tests {
		tok := l.NextToken()
		fmt.Println(tok.Column)
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q in column %d",
				i, tt.expectedType, tok.Type, tok.Column)
		}
		if tok.Value != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Value)
		}
	}
}

func TestColumnsOfLine(t *testing.T) {
	input := `=+(){},;`
	columns := []int{1, 2, 3, 4, 5, 6, 7, 8}

	l := New(input)
	// l.readChar()
	for i, tt := range columns {
		tok := l.NextToken()
		// fmt.Println(tok.Column)
		if tok.Column != tt {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%d, got=%d",
				i, tt, tok.Column)
		}
		// if tok.Value != tt.expectedLiteral {
		// 	t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
		// 		i, tt.expectedLiteral, tok.Value)
		// }
	}
}

func TestColumnsOfLineWithSpaces(t *testing.T) {
	input := `= + ( ) { } , ;`
	columns := []int{1, 3, 5, 7, 9, 11, 13, 15}

	l := New(input)
	// l.readChar()
	for i, tt := range columns {
		tok := l.NextToken()
		// fmt.Println(tok.Column)
		if tok.Column != tt {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%d, got=%d",
				i, tt, tok.Column)
		}
		// if tok.Value != tt.expectedLiteral {
		// 	t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
		// 		i, tt.expectedLiteral, tok.Value)
		// }
	}
}

func TestColumnsOfLineWithNewLines(t *testing.T) {
	input :=
		`=+
(){},;
let a = 1;
`
	columns := []int{1, 2,
		1, 2, 3, 4, 5, 6,
		1, 5, 7, 9, 10,
	}

	l := New(input)
	// l.readChar()
	for i, tt := range columns {
		tok := l.NextToken()
		// fmt.Println(tok.Column)
		if tok.Column != tt {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%d, got=%d",
				i, tt, tok.Column)
		}
		// if tok.Value != tt.expectedLiteral {
		// 	t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
		// 		i, tt.expectedLiteral, tok.Value)
		// }
	}
}

func TestLexerLineWithWhitespace(t *testing.T) {
	input :=
		`let a = 1;
		let b == 2;
		+
		==
		*
		/`

	// lines := []int{0, 0, 1, 0, 1, 0, 0, 2, 0}

	lines := []int{1, 1, 1, 1, 1,
		2, 2, 2, 2, 2,
		3, 4,
		5, 6,
	}

	l := New(input)
	// t.Fatalf("tests[%d] - starting line=%d", 0, l.line)
	// l.readChar()
	for i, tt := range lines {
		l.NextToken()

		if l.line != tt {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%d, got=%d",
				i, tt, l.line)
		}

	}
}

func TestLineWithWhitespace(t *testing.T) {
	input :=
		`let a = 1;
		let b == 2;
		+
		==
		*
		/`

	lines := []int{1, 1, 1, 1, 1,
		2, 2, 2, 2, 2,
		3, 4,
		5, 6,
	}

	l := New(input)

	for i, tt := range lines {
		var tok = l.NextToken()
		// t.Fatalf("tests[%d] - starting line=%d", 0, l.line)
		// fmt.Println(tok.Column)
		if tok.Line != tt {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%d, got=%d",
				i, tt, tok.Line)
		}

	}
}

func TestLexerLineWithWhitespaceAndNewLines(t *testing.T) {
	input :=
		`let a = 1;

let b = 2;

+
-

*
/`

	// lines := []int{0, 0, 1, 0, 1, 0, 0, 2, 0}

	lines := []int{1, 1, 1, 1, 1,
		3, 3, 3, 3, 3,
		5, 6,
		8, 9,
	}

	l := New(input)
	// t.Fatalf("tests[%d] - starting line=%d", 0, l.line)
	// l.readChar()
	for i, tt := range lines {

		l.NextToken()
		// t.Fatalf("tests[%d] - starting line=%d", 0, l.line)
		// fmt.Println(tok.Column)
		if l.line != tt {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%d, got=%d",
				i, tt, l.line)
		}
		// if tok.Value != tt.expectedLiteral {
		// 	t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
		// 		i, tt.expectedLiteral, tok.Value)
		// }
	}
}

func TestColumnsAndLinesDetailedTest(t *testing.T) {
	input :=
		`let a = 1;

let b = 2;

+
-

enum Test = a|b|c;
func test(x, y) {
    x + y;
    x
}

if (x < y) {
    return x;
} else {
    return y;
}
		`
	columns := []int{1, 5, 7, 9, 10,
		1, 5, 7, 9, 10,
		1,
		1,
		1, 6, 11, 13, 14, 15, 16, 17, 18,
		1, 6, 10, 11, 12, 14, 15, 17,
		5, 7, 9, 10,
		5,
		1,
		1, 4, 5, 7, 9, 10, 12,
		5, 12, 13,
		1, 3, 8,
		5, 12, 13,
		1,
	}

	lines := []int{1, 1, 1, 1, 1,
		3, 3, 3, 3, 3,
		5, 6,
		8, 8, 8, 8, 8, 8, 8, 8, 8,
		9, 9, 9, 9, 9, 9, 9, 9,
		10, 10, 10, 10,
		11,
		12,
		14, 14, 14, 14, 14, 14, 14,
		15, 15, 15,
		16, 16, 16,
		17, 17, 17,
		18,
	}

	l := New(input)
	// t.Fatalf("tests[%d] - starting line=%d", 0, l.line)
	// l.readChar()
	for i, tt := range lines {

		tok := l.NextToken()
		// t.Fatalf("tests[%d] - starting line=%d", 0, l.line)
		// fmt.Println(tok.Column)
		t.Logf(`%d:%d %s`, tok.Line, tok.Column, tok.Value)
		if tok.Line != tt {
			t.Fatalf("tests[%d] - line is wrong. expected=%d, got=%d (btw current column is %d and token is '%s')",
				i, tt, tok.Line, tok.Column, tok.Value)
		}
		if tok.Column != columns[i] {
			t.Fatalf("tests[%d] - column is wrong. expected=%d, got=%d (btw current line is %d and token is '%s')",
				i, columns[i], tok.Column, tok.Line, tok.Value)
		}
		// if tok.Value != tt.expectedLiteral {
		// 	t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
		// 		i, tt.expectedLiteral, tok.Value)
		// }

	}
}
func Test_isLetter(t *testing.T) {
	type args struct {
		ch byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.

		{"a", args{'a'}, true},
		{"b", args{'b'}, true},
		{"c", args{'c'}, true},
		{"d", args{'d'}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLetter(tt.args.ch); got != tt.want {
				t.Errorf("isLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
