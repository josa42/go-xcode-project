package lexer

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/josa42/go-xcode-project/pbxproj/token"
	"github.com/stretchr/testify/assert"
)

func fixture(name string) string {
	b, _ := ioutil.ReadFile(filepath.Join("testdata", name+".pbxproj"))
	return string(b)
}

func TestNextToken(t *testing.T) {

	tests := []struct {
		name   string
		expect []token.Token
	}{
		{"number", []token.Token{
			{token.LBRACE, "{"},
			{token.KEY, "number"},
			{token.ASSIGN, "="},
			{token.STRING, "1"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
		}},
		{"numbers", []token.Token{
			{token.LBRACE, "{"},
			{token.KEY, "number1"},
			{token.ASSIGN, "="},
			{token.STRING, "1"},
			{token.SEMICOLON, ";"},
			{token.KEY, "number2"},
			{token.ASSIGN, "="},
			{token.STRING, "2"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
		}},
		{"comment", []token.Token{
			{token.LBRACE, "{"},
			{token.KEY, "number"},
			{token.ASSIGN, "="},
			{token.STRING, "1"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
		}},
		{"strings", []token.Token{
			{token.LBRACE, "{"},
			{token.KEY, "str1"},
			{token.ASSIGN, "="},
			{token.STRING, "MyString"},
			{token.SEMICOLON, ";"},
			{token.KEY, "str2"},
			{token.ASSIGN, "="},
			{token.STRING, "My String"},
			{token.SEMICOLON, ";"},
			{token.KEY, "str3"},
			{token.ASSIGN, "="},
			{token.STRING, "My \"complex\" String"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
		}},
		{"lists", []token.Token{
			{token.LBRACE, "{"},
			{token.KEY, "list1"},
			{token.ASSIGN, "="},
			{token.LPAREN, "("},
			{token.STRING, "1"},
			{token.COMMA, ","},
			{token.STRING, "2"},
			{token.RPAREN, ")"},
			{token.SEMICOLON, ";"},
			{token.KEY, "list2"},
			{token.ASSIGN, "="},
			{token.LPAREN, "("},
			{token.STRING, "one"},
			{token.COMMA, ","},
			{token.STRING, "two"},
			{token.COMMA, ","},
			{token.RPAREN, ")"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
		}},
		{"dict", []token.Token{
			{token.LBRACE, "{"},
			{token.KEY, "dict1"},
			{token.ASSIGN, "="},
			{token.LBRACE, "{"},
			{token.KEY, "key1"},
			{token.ASSIGN, "="},
			{token.STRING, "1"},
			{token.SEMICOLON, ";"},
			{token.KEY, "key2"},
			{token.ASSIGN, "="},
			{token.STRING, "two"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
		}},
		{"list-numbers", []token.Token{
			{token.LPAREN, "("},
			{token.STRING, "1"},
			{token.COMMA, ","},
			{token.STRING, "2"},
			{token.COMMA, ","},
			{token.STRING, "3"},
			{token.RPAREN, ")"},
		}},
		{"filename", []token.Token{
			{token.LBRACE, "{"},
			{token.KEY, "name"},
			{token.ASSIGN, "="},
			{token.STRING, "file.swift"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
		}},
		{"comments", []token.Token{
			{token.LBRACE, "{"},
			{token.KEY, "key"},
			{token.ASSIGN, "="},
			{token.STRING, "1"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
		}},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			l := New(fixture(tt.name))

			tokens := []token.Token{}
			for {
				tok := l.NextToken()
				if tok.Type == token.EOF {
					break
				}
				tokens = append(tokens, tok)
			}

			assert.Equal(t, tt.expect, tokens)
		})
	}
}
