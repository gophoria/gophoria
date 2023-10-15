package parser

import (
	"fmt"

	"github.com/gophoria/gophoria/pkg/ast"
	"github.com/gophoria/gophoria/pkg/lexer"
)

type Parser struct {
	lexer *lexer.Lexer

	currToken *lexer.Token
	peekToken *lexer.Token
}

func NewParser(lexer *lexer.Lexer) *Parser {
	p := Parser{
		lexer:     lexer,
		currToken: nil,
		peekToken: nil,
	}

	p.nextToken()
	p.nextToken()

	return &p
}

func (p *Parser) Parse() (*ast.Ast, error) {
	ast := ast.NewAst()

	for !p.curTokenIs(lexer.TokenTypeEof) {
		if p.curTokenIs(lexer.TokenTypeDb) {
			cfg, err := p.parseConfig()
			if err != nil {
				return nil, err
			}

			ast.Config = append(ast.Config, cfg)
		}
		p.nextToken()
	}

	return ast, nil
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.lexer.Next()
}

func (p *Parser) curTokenIs(tokenType lexer.TokenType) bool {
	return p.currToken.Type == tokenType
}

func (p *Parser) peekTokenIs(tokenType lexer.TokenType) bool {
	return p.peekToken.Type == tokenType
}

func (p *Parser) parseConfig() (*ast.Config, error) {
	config := ast.NewConfig(p.currToken)

	p.nextToken()
	if !p.curTokenIs(lexer.TokenTypeLBrace) {
		return nil, fmt.Errorf("[line: %d, col: %d]: expected { but found %s", p.peekToken.Row, p.peekToken.Col, p.peekToken.Literal)
	}

	p.nextToken()

	for !p.curTokenIs(lexer.TokenTypeRBrace) {
		item, err := p.parseAssignItem()
		if err != nil {
			return nil, err
		}

		config.AddItem(item)
	}

	return config, nil
}

func (p *Parser) parseAssignItem() (*ast.AssignItem, error) {
	if !p.curTokenIs(lexer.TokenTypeIdent) {
		return nil, fmt.Errorf("[line: %d, col: %d]: expected identifier but found %s", p.currToken.Row, p.currToken.Col, p.currToken.Literal)
	}
	ident := ast.NewIdentifier(p.currToken)

	p.nextToken()

	if !p.curTokenIs(lexer.TokenTypeAssign) {
		return nil, fmt.Errorf("[line: %d, col: %d]: expected = but found %s", p.currToken.Row, p.currToken.Col, p.currToken.Literal)
	}
	if !p.peekTokenIs(lexer.TokenTypeString) && !p.peekTokenIs(lexer.TokenTypeInt) {
		return nil, fmt.Errorf("[line: %d, col: %d]: expected value but found %s", p.peekToken.Row, p.peekToken.Col, p.peekToken.Literal)
	}

	val := ast.NewValue(p.peekToken)

	item := ast.NewAssignItem(p.currToken, ident, val)

	p.nextToken()
	p.nextToken()

	return item, nil
}
