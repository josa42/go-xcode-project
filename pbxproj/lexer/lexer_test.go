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
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			l := New(fixture(tt.name))

			// ln := 0
			tokens := []token.Token{}
			for {
				tok := l.NextToken()
				if tok.Type == token.EOF {
					break
				}
				tokens = append(tokens, tok)
			}

			assert.Equal(t, tt.expect, tokens)

			// var tok token.Token
			// for tok = ""; tok = l.NextToken(); tok.Type != token.EOF {
			// 	ln += 1
			//
			// 	fmt.Println("lit", tok.Literal)
			// 	assert.Equal(t, tt.expect[ln], tok)
			// }

			// assert.Equal(t, len(tt.expect), ln)

		})
	}
	//
	// 	input := `let five = 5;
	// let ten = 10;
	//
	// let add = fn(x, y) {
	//   x + y;
	// };
	//
	// let result = add(five, ten);
	// !-/*5;
	// 5 < 10 > 5;
	//
	// if (5 < 10) {
	// 	return true;
	// } else {
	// 	return false;
	// }
	//
	// 10 == 10;
	// 10 != 9;
	// "foobar"
	// "foo bar"
	// [1, 2];
	// {"foo": "bar"}
	// `
	//
	// 	tests := []struct {
	// 		expectedType    token.TokenType
	// 		expectedLiteral string
	// 	}{}
	//
	// 	l := New(input)
	//
	// 	for i, tt := range tests {
	// 		tok := l.NextToken()
	//
	// 		if tok.Type != tt.expectedType {
	// 			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
	// 				i, tt.expectedType, tok.Type)
	// 		}
	//
	// 		if tok.Literal != tt.expectedLiteral {
	// 			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
	// 				i, tt.expectedLiteral, tok.Literal)
	// 		}
	// 	}
}
