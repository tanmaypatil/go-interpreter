package ast

 import "example/hello/token"

type Node interface {
	tokenLiteral() string
}

type Statement interface {
	Node
	StatementNode()
}

type Expression interface {
	Node
	ExpressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].tokenLiteral()
	} else {
		return ""
	}
}

type LetStatment struct {
	Token token.Token // the token.LET token 
	Name *Identifier 
	Value Expression
}

func ( ls *LetStatment) StatementNode() {}
func ( ls *LetStatment) TokenLiteral()string { return ls.Token.Literal}

type Identifier struct { 
	Token token.Token
	Value string 
}

func ( i *Identifier) ExpressionNode() {}
func ( i *Identifier) TokenLiteral() string { return i.Token.Literal}
