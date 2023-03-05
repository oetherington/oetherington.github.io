package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	. "github.com/oetherington/smetana"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
	"go.abhg.dev/goldmark/toc"
)

type HeadingHTMLRenderer struct {
	html.Config
}

func NewHeadingHTMLRenderer(opts ...html.Option) renderer.NodeRenderer {
	r := &HeadingHTMLRenderer{
		Config: html.NewConfig(),
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

func (r *HeadingHTMLRenderer) RegisterFuncs(
	reg renderer.NodeRendererFuncRegisterer,
) {
	reg.Register(ast.KindHeading, r.renderHeading)
}

func (r *HeadingHTMLRenderer) renderHeading(
	w util.BufWriter,
	source []byte,
	node ast.Node,
	entering bool,
) (ast.WalkStatus, error) {
	n := node.(*ast.Heading)

	id := ""
	attrs := n.Attributes()
	for _, attr := range attrs {
		if bytes.Equal(attr.Name, []byte{'i', 'd'}) {
			switch value := attr.Value.(type) {
			case string:
				id = value
			case []byte:
				id = string(value)
			default:
				break
			}
		}
	}

	if entering {
		_, _ = w.WriteString("<h")
		_ = w.WriteByte("0123456"[n.Level])
		if attrs != nil {
			html.RenderAttributes(w, node, html.HeadingAttributeFilter)
		}
		if len(id) > 0 {
			output := fmt.Sprintf("><a href=\"#%s\">&gt; ", id)
			_, _ = w.WriteString(output)
		} else {
			_, _ = w.WriteString(">&gt; ")
		}
	} else {
		if len(id) > 0 {
			_, _ = w.WriteString("</a></h")
		} else {
			_, _ = w.WriteString("</h")
		}
		_ = w.WriteByte("0123456"[n.Level])
		_, _ = w.WriteString(">\n")
	}
	return ast.WalkContinue, nil
}

func renderMarkdownFile(path string) (Node, Node, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	defer file.Close()

	src, err := io.ReadAll(file)
	if err != nil {
		return nil, nil, err
	}

	markdown := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			highlighting.NewHighlighting(
				highlighting.WithFormatOptions(
					chromahtml.WithClasses(true),
					chromahtml.WithLineNumbers(false),
				),
			),
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)
	markdown.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewHeadingHTMLRenderer(), 1),
	))

	doc := markdown.Parser().Parse(text.NewReader(src))

	tree, err := toc.Inspect(doc, src)
	if err != nil {
		return nil, nil, err
	}

	contents := Fragment()
	for i, item := range tree.Items {
		node := P(
			AHref(
				fmt.Sprintf("#%s", item.ID),
				fmt.Sprintf("%d. %s", i+1, item.Title),
			),
		)
		contents.Children = append(contents.Children, node)
	}

	var output strings.Builder
	err = markdown.Renderer().Render(&output, src, doc)
	if err != nil {
		return nil, nil, err
	}

	return Text(output.String()), contents, nil
}
