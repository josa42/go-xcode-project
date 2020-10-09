package formatter

import (
	"fmt"
	"strings"

	"github.com/josa42/go-xcode-project/pbxproj/ast"
)

func Format(n ast.Node) string {
	lines := []string{}

	lines = append(lines, "// !$*UTF8*$!")
	lines = append(lines, formatNode(n, 0, "")...)

	return strings.Join(lines, "\n")
}

func formatNode(n ast.Node, depth int, sign string) []string {
	// pre := strings.Repeat("\t", depth)
	lines := []string{}

	switch n.Type() {
	case ast.LIST:
		if l, ok := n.(*ast.ListNode); ok {
			lines = append(lines, formatList(l, depth)...)
		}
	case ast.DICT:
		if l, ok := n.(*ast.DictNode); ok {
			lines = append(lines, formatDict(l, depth)...)
		}
	case ast.STRING:
		if l, ok := n.(ast.StringNode); ok {
			lines = append(lines, formatString(l.Value, depth))
		}
	}

	if sign != "" && len(lines) > 0 {
		lines[len(lines)-1] += sign
	}

	return lines
}

func formatList(n *ast.ListNode, depth int) []string {
	pre := strings.Repeat("\t", depth)
	lines := []string{}

	lines = append(lines, "(")
	for _, s := range n.Values {
		lines = append(lines, fmt.Sprintf("%s\t%s", pre, strings.Join(formatNode(s, depth+1, ";"), "\n")))
	}
	lines = append(lines, fmt.Sprintf("%s}", pre))

	return lines
}

func formatDict(n *ast.DictNode, depth int) []string {
	pre := strings.Repeat("\t", depth)
	lines := []string{}

	lines = append(lines, "{")
	for key, s := range n.Values {
		lines = append(
			lines,
			fmt.Sprintf("%s\t%s = %s", pre, key, strings.Join(formatNode(s, depth+1, ";"), "\n")),
		)

	}
	lines = append(lines, fmt.Sprintf("%s}", pre))

	return lines
}

func formatString(s string, depth int) string {
	return fmt.Sprintf(`"%s"`, strings.ReplaceAll(s, "\"", "\\\""))
}
