package parser

import (
	"monkeyInterpreter/src/ast"
	"monkeyInterpreter/src/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("Program.Statements does not contain 3 statements. Got=%d",
			len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"}, {"y"}, {"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]

		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}


func testLetStatement(
	t *testing.T,
	s ast.Statement,
	name string) bool {

	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. Got=%q",
			s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("s not *ast.LetStatement. Got=%T",
			s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. Got=%s",
			name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. Got=%s",
			name, letStmt.Name.Value)
		return false
	}

	return true
}

func TestConstStatements(t *testing.T) {
	input := `
	const x = 5;
	const y = 10;
	const foobar = 838383;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("Program.Statements does not contain 3 statements. Got=%d",
			len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"}, {"y"}, {"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]

		if !testConstStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testConstStatement(
	t *testing.T,
	s ast.Statement,
	name string) bool {

	if s.TokenLiteral() != "const" {
		t.Errorf("s.TokenLiteral not 'const'. Got=%q",
			s.TokenLiteral())
		return false
	}

	constStmt, ok := s.(*ast.ConstStatement)

	if !ok {
		t.Errorf("s not *ast.ConstStatement. Got=%T",
			s)
		return false
	}

	if constStmt.Name.Value != name {
		t.Errorf("constStmt.Name.Value not '%s'. Got=%s",
			name, constStmt.Name.Value)
		return false
	}

	if constStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. Got=%s",
			name, constStmt.Name.Value)
		return false
	}

	return true
}

// wtf is this
// func TestGetRepeatingSequence(t *testing.T) {
// 	input := `
// 		enum Color = red|green|blue;
// 		enum State = on|off;
// 		enum Direction = up|down|left|right;
// 		enum Me = me;
// 	`

// 	l := lexer.New(input)
// 	p := New(l)

// 	program := p.ParseProgram() // todo remove this from test

// 	if program == nil {
// 		t.Fatalf("ParseProgram() returned nil")
// 	}

// 	if len(program.Statements) != 3 {
// 		t.Fatalf("Program.Statements does not contain 3 statements. Got=%d",
// 			len(program.Statements))
// 	}

// 	tests := []struct {
// 		repeater []token.TokenType
// 		stop     token.TokenType
// 		expected []token.Token
// 	}{
// 		{
// 			[]token.TokenType{token.IDENT, token.PIPE},
// 			token.SEMICOLON,
// 			[]token.Token{
// 				{Type: token.IDENT, Value: "red"},
// 				{Type: token.PIPE, Value: "|"},
// 				{Type: token.IDENT, Value: "green"},
// 				{Type: token.PIPE, Value: "|"},
// 				{Type: token.IDENT, Value: "blue"},
// 			},
// 		},
// 		{
// 			[]token.TokenType{token.IDENT, token.PIPE},
// 			token.SEMICOLON,
// 			[]token.Token{
// 				{Type: token.IDENT, Value: "on"},
// 				{Type: token.PIPE, Value: "|"},
// 				{Type: token.IDENT, Value: "off"},
// 			},
// 		},
// 		{
// 			[]token.TokenType{token.IDENT, token.PIPE},
// 			token.SEMICOLON,
// 			[]token.Token{
// 				{Type: token.IDENT, Value: "up"},
// 				{Type: token.PIPE, Value: "|"},
// 				{Type: token.IDENT, Value: "down"},
// 				{Type: token.PIPE, Value: "|"},
// 				{Type: token.IDENT, Value: "left"},
// 				{Type: token.PIPE, Value: "|"},
// 				{Type: token.IDENT, Value: "right"},
// 			},
// 		},
// 		{
// 			[]token.TokenType{token.IDENT, token.PIPE},
// 			token.SEMICOLON,
// 			[]token.Token{
// 				{Type: token.IDENT, Value: "me"},
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		repeatingSequence := p.getRepeatingSequence(tt.repeater, tt.stop)

// 		if len(repeatingSequence) != len(tt.expected) {
// 			t.Errorf("getRepeatingSequence() returned wrong number of tokens. Got=%d, Expected=%d",
// 				len(repeatingSequence), len(tt.expected))
// 		}

// 		for i, token := range repeatingSequence {
// 			if token.Type != tt.expected[i].Type {
// 				t.Errorf("getRepeatingSequence() returned wrong token type at index %d. Got=%s, Expected=%s",
// 					i, token.Type, tt.expected[i].Type)
// 			}

// 			if token.Value != tt.expected[i].Value {
// 				t.Errorf("getRepeatingSequence() returned wrong token value at index %d. Got=%s, Expected=%s",
// 					i, token.Value, tt.expected[i].Value)
// 			}
// 		}
// 	}
// }
