package parser

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/josa42/go-xcode-project/pbxproj/ast"
	"github.com/josa42/go-xcode-project/pbxproj/lexer"
	"github.com/stretchr/testify/assert"
)

func fixture(name string) string {
	b, _ := ioutil.ReadFile(filepath.Join("testdata", name+".pbxproj"))
	return string(b)
}

func TestParser_Parse(t *testing.T) {

	tests := []struct {
		name   string
		expect ast.Node
	}{
		{"dict-empty", &ast.DictNode{}},
		{"dict", &ast.DictNode{
			Values: map[string]ast.Node{
				"key1": ast.StringNode{"1"},
				"key2": ast.StringNode{"two"},
			},
		}},
		{"list-empty", &ast.ListNode{}},
		{"list-numbers", &ast.ListNode{
			Values: []ast.Node{ast.StringNode{"1"}, ast.StringNode{"2"}, ast.StringNode{"3"}},
		}},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(fixture(tt.name))
			p := New(l)
			n, err := p.Parse()

			assert.Equal(t, nil, err)
			assert.Equal(t, tt.expect, n)
		})
	}

	t.Run("project", func(t *testing.T) {
		l := lexer.New(fixture("project"))
		p := New(l)
		_, err := p.Parse()

		assert.Equal(t, nil, err)
	})
}
