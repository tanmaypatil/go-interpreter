package parser

import (
	"example/hello/ast"
	"example/hello/lexer"
	"testing"
)

func TestLetstatements( t  *testing.T) {
	input := `
	let x = 5; 
	let y = 10;
	let foobar = 838383 ;
	`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("parseProgram() returned nil ")
	}

	if len(program.Statements)  != 3 {
		t.Fatalf("program Statements does not contain 3 statements , got = %d ",len(program.Statements))
	}

	checkParserErrors(t,p)

	tests := [] struct {
		expectedIdentifier string
	} {
		{"x"},
		{ "y"},
		{ "foobar" },
	}
	for i,tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t,stmt,tt.expectedIdentifier) {
			return 
		}
			
	}
}

// each statement has parsing error
func TestLetstatementParseErrors( t  *testing.T) {
	input := `
	let x  5; 
	let = 10;
	let  838383 ;
	`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("parseProgram() returned nil ")
	}

	if len(program.Statements)  != 3 {
		t.Fatalf("program Statements does not contain 3 statements , got = %d ",len(program.Statements))
	}

	checkParserErrors(t,p)

	tests := [] struct {
		expectedIdentifier string
	} {
		{"x"},
		{ "y"},
		{ "foobar" },
	}
	for i,tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t,stmt,tt.expectedIdentifier) {
			return 
		}
			
	}
}

func TestReturnStatement( t  *testing.T) {
	input := `
	return 5; 
	return 10;
	return 838383 ;
	`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("parseProgram() returned nil ")
	}

	checkParserErrors(t,p)
	
	if len(program.Statements)  != 3 {
		t.Fatalf("program Statements does not contain 3 statements , got = %d ",len(program.Statements))
	}

	for _,stmt := range program.Statements {
		returnStmt ,ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not ast.*ReturnStatement ,got %T ",stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral != return , got %q",returnStmt.TokenLiteral())
		}
	}


	
}



func testLetStatement( t *testing.T ,s ast.Statement , name string ) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let' ,got = %q ,s.TokenLiteral()",s.TokenLiteral())
		return false
	}

	letStmt , ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *.ast.Statement , got = %T",s)
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not %s , got %s",name,letStmt.Name.Value)
		return false 
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s' got ='%s'",name,letStmt.Name.TokenLiteral() )
		return false
	}
	return true
}

func checkParserErrors(t *testing.T , p *Parser) {
	errors := p.Errors()
	if len(errors) == 0  {
		return
	}
	t.Errorf("parser had %d errors",len(p.errors))
	for _,msg := range errors {
		t.Errorf("parser error : %q",msg)
	}
	t.FailNow()
}

func TestIdentifierExpression( t *testing.T  ){
	input := "foobar;"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t,p)

	if  len(program.Statements) != 1 {
		t.Fatalf("program has not had enough statements,got =%d",len(program.Statements))
	}
	stmt,ok := program.Statements[0].(*ast.ExpressionStatement) 
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got =%T ",program.Statements[0])
	}

	ident,ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not ast.Identifier , got =%T ",stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Fatalf("ident.Value not [%s] , got %s","foobar",ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Fatalf("ident.TokenLiteral() not [%s] , got %s","foobar",ident.TokenLiteral())
	}

}