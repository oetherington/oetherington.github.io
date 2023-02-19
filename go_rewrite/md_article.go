package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	. "github.com/oetherington/smetana"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

func renderMarkdownFile(path string) (Node, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	md, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	renderer := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
	)

	var buf bytes.Buffer
	if err := renderer.Convert(md, &buf); err != nil {
		return nil, err
	}

	return Text(buf.String()), nil
}

func MdArticle(palette Palette, articleInfo ArticleInfo) HtmlNode {
	mdPath := fmt.Sprintf("./articles/%s.md", articleInfo.Path)
	md, err := renderMarkdownFile(mdPath)
	if err != nil {
		fmt.Println(err)
	}

	return Layout(
		palette,
		articleInfo.Name,
		Fragment(
			Div(
				ClassName("center"),
				H2(
					Span(
						ClassName("no-mobile"),
						"==== ",
					),
					Text(articleInfo.Name),
					Text(" "),
					Span(
						ClassName("no-mobile"),
						" ====",
					),
				),
			),
			Div(
				ClassNames("content", "centered", "center"),
				Hr(),
				P("TODO: TOC"),
			),
			Div(
				ClassNames("content", "centered"),
				P(md),
			),
		),
	)
}
