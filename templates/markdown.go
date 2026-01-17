package templates

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type CLIRenderer struct{}

// RenderMarkdownForCobra converts markdown input into terminal-friendly output
// suitable for use in [cobra.Command] help text.
func RenderMarkdownForCobra(input string) string {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithRenderer(
			renderer.NewRenderer(
				renderer.WithNodeRenderers(
					util.Prioritized(&CLIRenderer{}, 100),
				),
			),
		),
	)

	var buf bytes.Buffer
	_ = md.Convert([]byte(input), &buf)
	return strings.TrimSpace(buf.String())
}

// RegisterFuncs registers rendering functions for various markdown AST node
// types.
func (r *CLIRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindParagraph, r.renderParagraph)
	reg.Register(ast.KindText, r.renderText)
	reg.Register(ast.KindList, r.renderList)
	reg.Register(ast.KindListItem, r.renderListItem)
	reg.Register(ast.KindCodeBlock, r.renderCodeBlock)
	reg.Register(ast.KindFencedCodeBlock, r.renderCodeBlock)
}

func (r *CLIRenderer) renderParagraph(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		w.WriteString("\n\n")
	}
	return ast.WalkContinue, nil
}

func (r *CLIRenderer) renderText(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}

	t := node.(*ast.Text)
	text := string(t.Segment.Value(source))

	w.WriteString(text)
	if t.SoftLineBreak() {
		w.WriteString(" ")
	}

	return ast.WalkContinue, nil
}

func (r *CLIRenderer) renderList(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		w.WriteString("\n")
	}
	return ast.WalkContinue, nil
}

func (r *CLIRenderer) renderListItem(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		w.WriteString("\n")
		return ast.WalkContinue, nil
	}

	item := node.(*ast.ListItem)
	list := item.Parent().(*ast.List)

	if list.IsOrdered() {
		start := list.Start
		idx := orderedIndex(item)
		fmt.Fprintf(w, "  %d. ", start+idx)
	} else {
		w.WriteString("  * ")
	}

	return ast.WalkContinue, nil
}

func (r *CLIRenderer) renderCodeBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}

	var lines *text.Segments

	// Go doesn’t allow shared methods on a type switch variable so the following
	// two cases are necessary.
	switch n := node.(type) {
	case *ast.CodeBlock:
		lines = n.Lines()
	case *ast.FencedCodeBlock:
		lines = n.Lines()
	}

	for i := range lines.Len() {
		line := lines.At(i)
		w.WriteString("      ") // 6 spaces
		w.Write(line.Value(source))
	}
	w.WriteString("        \n")

	return ast.WalkContinue, nil
}

func orderedIndex(item *ast.ListItem) int {
	index := 0
	for sibling := item.PreviousSibling(); sibling != nil; sibling = sibling.PreviousSibling() {
		index++
	}
	return index
}
