package parser

import (
	"errors"
	"fmt"
	"strings"

	"github.com/josa42/go-xcode-project/pbxproj/ast"
	"github.com/josa42/go-xcode-project/pbxproj/lexer"
	"github.com/josa42/go-xcode-project/pbxproj/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{}
	p.l = l

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Parse() (ast.Node, error) {

	if p.curTokenIs(token.LBRACE) {
		return p.parseDict()

	} else if p.curTokenIs(token.LPAREN) {
		return p.parseList()

	} else {
		return nil, errors.New("Invalid Start Token")
	}
}

func (p *Parser) parseList() (*ast.ListNode, error) {

	if err := p.expectTokenIs(token.LPAREN); err != nil {
		return nil, err
	}

	p.nextToken()

	list := &ast.ListNode{}

	for !p.curTokenIs(token.EOF) {
		switch p.curToken.Type {
		case token.RPAREN:
			p.nextToken()
			return list, nil
		case token.COMMA:
			p.nextToken()
		default:

			v, err := p.parseValue()
			if err != nil {
				return nil, err
			}

			list.Append(v)

			if err := p.expectTokenIs(token.COMMA, token.RPAREN); err != nil {
				return nil, err
			}
		}

	}

	return nil, errors.New("Syntax Error: Missing }")
}

func (p *Parser) parseDict() (*ast.DictNode, error) {

	if err := p.expectTokenIs(token.LBRACE); err != nil {
		return nil, err
	}

	p.nextToken()

	dict := &ast.DictNode{}

	for !p.curTokenIs(token.EOF) {
		switch p.curToken.Type {
		case token.KEY:
			key := p.curToken
			p.nextToken()
			if err := p.expectTokenIs(token.ASSIGN); err != nil {
				return nil, err
			}

			p.nextToken()

			v, err := p.parseValue()
			if err != nil {
				return nil, err
			}

			dict.Set(key.Literal, v)

			if err := p.expectTokenIs(token.SEMICOLON, token.RBRACE); err != nil {
				return nil, err
			}
		case token.SEMICOLON:
			p.nextToken()

		case token.RBRACE:
			p.nextToken()
			return dict, nil

		default:
			return nil, fmt.Errorf("Syntax Error: Unexpected '%s' [%s]", p.curToken.Literal, p.curToken.Type)
		}
	}

	return nil, errors.New("Syntax Error: Missing }")
}

func (p *Parser) parseValue() (ast.Node, error) {
	switch p.curToken.Type {
	case token.LBRACE:
		return p.parseDict()
	case token.LPAREN:
		return p.parseList()
	case token.STRING:
		s := ast.StringNode{Value: p.curToken.Literal}
		p.nextToken()
		return s, nil
	default:
		return nil, fmt.Errorf("Syntax Error: Unxpected %s", p.curToken.Type)
	}

}

func (p *Parser) nextToken() {

	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
	// fmt.Printf("curToken = '%s' [%s]\n", p.curToken.Literal, p.curToken.Type)
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectTokenIs(ts ...token.TokenType) error {

	for _, t := range ts {
		if p.curToken.Type == t {
			return nil
		}
	}

	exp := ""
	if len(ts) == 1 {
		exp = string(ts[0])
	} else if len(ts) > 1 {
		l := ts[len(ts)-1]
		r := []string{}
		for _, t := range ts[:len(ts)-1] {
			r = append(r, string(t))
		}

		exp = fmt.Sprintf("%s or %s", strings.Join(r, ", "), l)
	}

	return fmt.Errorf("Syntax Error: Expected %s | Not: %s", exp, p.curToken.Literal)
}
