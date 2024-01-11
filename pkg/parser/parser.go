package parser

import (
	"fmt"

	"github.com/gophoria/gophoria/pkg/ast"
	"github.com/gophoria/gophoria/pkg/lexer"
)

var VariableTypeToken = map[lexer.TokenType]struct{}{
	lexer.TokenTypeTInt:      {},
	lexer.TokenTypeTReal:     {},
	lexer.TokenTypeTString:   {},
	lexer.TokenTypeTBool:     {},
	lexer.TokenTypeTDateTime: {},
	lexer.TokenTypeIdent:     {},
}

var ArgumentTypeToken = map[lexer.TokenType]struct{}{
	lexer.TokenTypeString: {},
	lexer.TokenTypeInt:    {},
	lexer.TokenTypeTrue:   {},
	lexer.TokenTypeFalse:  {},
	lexer.TokenTypeIdent:  {},
}

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
		if p.curTokenIs(lexer.TokenTypeDb) || p.curTokenIs(lexer.TokenTypeUi) {
			cfg, err := p.parseConfig()
			if err != nil {
				return nil, err
			}

			ast.Config = append(ast.Config, cfg)
		} else if p.curTokenIs(lexer.TokenTypeEnum) {
			en, err := p.parseEnum()
			if err != nil {
				return nil, err
			}

			ast.Enums = append(ast.Enums, en)
		} else if p.curTokenIs(lexer.TokenTypeModel) {
			en, err := p.parseModel()
			if err != nil {
				return nil, err
			}

			ast.Models = append(ast.Models, en)
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

func (p *Parser) parseEnum() (*ast.Enum, error) {
	if !p.peekTokenIs(lexer.TokenTypeIdent) {
		return nil, fmt.Errorf("[line: %d, col: %d]: expected identifier but found %s", p.peekToken.Row, p.peekToken.Col, p.peekToken.Literal)
	}

	ident := ast.NewIdentifier(p.peekToken)
	enum := ast.NewEnum(p.currToken, ident)

	p.nextToken()
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

		enum.AddItem(item)
	}

	return enum, nil
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

func (p *Parser) parseModel() (*ast.Model, error) {
	if !p.peekTokenIs(lexer.TokenTypeIdent) {
		return nil, fmt.Errorf("[line: %d, col: %d]: expected identifier but found %s", p.peekToken.Row, p.peekToken.Col, p.peekToken.Literal)
	}

	ident := ast.NewIdentifier(p.peekToken)
	model := ast.NewModel(p.currToken, ident)

	p.nextToken()
	p.nextToken()

	if !p.curTokenIs(lexer.TokenTypeLBrace) {
		return nil, fmt.Errorf("[line: %d, col: %d]: expected { but found %s", p.peekToken.Row, p.peekToken.Col, p.peekToken.Literal)
	}
	p.nextToken()

	for !p.curTokenIs(lexer.TokenTypeRBrace) {
		item, err := p.parseDeclaration()
		if err != nil {
			return nil, err
		}

		model.AddItem(item)
	}

	return model, nil
}

func (p *Parser) parseDeclaration() (*ast.Declaration, error) {
	if !p.curTokenIs(lexer.TokenTypeIdent) {
		return nil, fmt.Errorf("[line: %d, col: %d]: expected identifier but found %s", p.currToken.Row, p.currToken.Col, p.currToken.Literal)
	}

	ident := ast.NewIdentifier(p.currToken)
	p.nextToken()

	declType, err := p.parseType()
	if err != nil {
		return nil, err
	}

	decl := ast.NewDeclaration(ident, declType)

	for p.curTokenIs(lexer.TokenTypeDecorator) {
		dec, err := p.parseDecorator()
		if err != nil {
			return nil, err
		}

		decl.Decorators = append(decl.Decorators, dec)
	}

	return decl, nil
}

func (p *Parser) parseDecorator() (*ast.Decorator, error) {
	if !p.curTokenIs(lexer.TokenTypeDecorator) {
		return nil, fmt.Errorf("[line: %d, col: %d]: expected @ but found %s", p.currToken.Row, p.currToken.Col, p.currToken.Literal)
	}

	decToken := p.currToken
	p.nextToken()

	if !p.curTokenIs(lexer.TokenTypeIdent) {
		return nil, fmt.Errorf("[line: %d, col: %d]: expected identifier but found %s", p.currToken.Row, p.currToken.Col, p.currToken.Literal)
	}

	name := ast.NewIdentifier(p.currToken)
	dec := ast.NewDecorator(decToken, name)

	if p.peekTokenIs(lexer.TokenTypeLParen) {
		callable, err := p.parseCallable()
		if err != nil {
			return nil, err
		}

		dec.SetCallable(callable)
	} else {
		p.nextToken()
	}

	return dec, nil
}

func (p *Parser) parseCallable() (*ast.Callable, error) {
	if !p.curTokenIs(lexer.TokenTypeIdent) {
		return nil, fmt.Errorf("[line: %d, col: %d]: expected identifier but found %s", p.currToken.Row, p.currToken.Col, p.currToken.Literal)
	}
	if !p.peekTokenIs(lexer.TokenTypeLParen) {
		return nil, fmt.Errorf("[line: %d, col: %d]: expected ( but found %s", p.currToken.Row, p.currToken.Col, p.currToken.Literal)
	}

	ident := ast.NewIdentifier(p.currToken)
	p.nextToken()
	p.nextToken()

	callable := ast.NewCallable(ident)

	for !p.curTokenIs(lexer.TokenTypeRParen) {
		arg, err := p.parseArgument()
		if err != nil {
			return nil, err
		}
		callable.AddArgument(arg)

		if p.curTokenIs(lexer.TokenTypeComma) {
			p.nextToken()
		}
	}
	if !p.curTokenIs(lexer.TokenTypeRParen) {
		return nil, fmt.Errorf("[line: %d, col: %d]: expected ) but found %s", p.currToken.Row, p.currToken.Col, p.currToken.Literal)
	}
	p.nextToken()

	return callable, nil
}

func (p *Parser) parseArgument() (*ast.Argument, error) {
	if !p.isValidArgument(p.currToken) {
		return nil, fmt.Errorf("[line: %d, col: %d]: expected argument but found %s", p.currToken.Row, p.currToken.Col, p.currToken.Literal)
	}

	if p.peekTokenIs(lexer.TokenTypeLParen) {
		callable, err := p.parseCallable()
		if err != nil {
			return nil, err
		}

		arg := ast.NewArgument(nil, nil, callable)
		return arg, nil
	} else if p.peekTokenIs(lexer.TokenTypeColon) {
		name := ast.NewIdentifier(p.currToken)
		p.nextToken()
		p.nextToken()
		if !p.curTokenIs(lexer.TokenTypeIdent) {
			return nil, fmt.Errorf("[line: %d, col: %d]: expected identifier but found %s", p.currToken.Row, p.currToken.Col, p.currToken.Literal)
		}
		value := ast.NewValue(p.currToken)
		p.nextToken()

		arg := ast.NewArgument(name, value, nil)
		return arg, nil
	}

	value := ast.NewValue(p.currToken)
	arg := ast.NewArgument(nil, value, nil)
	p.nextToken()

	return arg, nil
}

func (p *Parser) parseType() (*ast.DeclarationType, error) {
	if !p.isValidvariableType(p.currToken) {
		return nil, fmt.Errorf("[line: %d, col: %d]: expected type but found %s", p.currToken.Row, p.currToken.Col, p.currToken.Literal)
	}
	typeToken := p.currToken
	isArray := false

	p.nextToken()
	if p.curTokenIs(lexer.TokenTypeLSquareBrace) {
		isArray = true

		p.nextToken()
		if !p.curTokenIs(lexer.TokenTypeRSquareBrace) {
			return nil, fmt.Errorf("[line: %d, col: %d]: expected ] but found %s", p.currToken.Row, p.currToken.Col, p.currToken.Literal)
		}
		p.nextToken()
	}

	declType := ast.NewDeclarationType(typeToken, isArray)

	return declType, nil
}

func (p *Parser) isValidvariableType(token *lexer.Token) bool {
	_, ok := VariableTypeToken[token.Type]
	return ok
}

func (p *Parser) isValidArgument(token *lexer.Token) bool {
	_, ok := ArgumentTypeToken[token.Type]
	return ok
}
