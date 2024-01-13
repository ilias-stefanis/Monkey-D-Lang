package ast

import (
	"Monkey-D-Lang/src/token"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Value
}

// ------- Let Statement -------
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Value
}

// ------- Const Statement -------
type ConstStatement struct {
	Token token.Token // the token.CONST token
	Name  *Identifier
	Value Expression
}

func (cs *ConstStatement) statementNode() {}
func (cs *ConstStatement) TokenLiteral() string {
	return cs.Token.Value
}

// ------- Return Statement -------
type ReturnStatement struct {
	Token token.Token // the token.RETURN token
	Value Expression
}

func (cs *ReturnStatement) statementNode() {}
func (cs *ReturnStatement) TokenLiteral() string {
	return cs.Token.Value
}

// ------- Enum Statement -------
type EnumStatement struct {
	Token token.Token // the token.ENUM token
	Name  *Identifier
	Value Expression
}

func (cs *EnumStatement) statementNode() {}
func (cs *EnumStatement) TokenLiteral() string {
	return cs.Token.Value
}
