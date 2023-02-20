package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	. "github.com/oetherington/smetana"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"go.abhg.dev/goldmark/toc"
)

func renderMarkdownFile(path string) (Node, Node, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	defer file.Close()

	src, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, nil, err
	}

	markdown := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
	)

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
	markdown.Renderer().Render(&output, src, doc)

	return Text(output.String()), contents, nil
}

func MdArticle(palette Palette, articleInfo ArticleInfo) HtmlNode {
	mdPath := fmt.Sprintf("./articles/%s.md", articleInfo.Path)
	md, contents, err := renderMarkdownFile(mdPath)
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
				contents,
				Hr(),
			),
			Div(
				ClassNames("content", "centered"),
				md,
			),
		),
	)
}
