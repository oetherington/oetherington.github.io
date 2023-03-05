package main

import (
	"fmt"
	. "github.com/oetherington/smetana"
)

func MdArticle(articleInfo ArticleInfo) HtmlNode {
	mdPath := fmt.Sprintf("./articles/%s.md", articleInfo.Path)
	md, contents, err := renderMarkdownFile(mdPath)
	if err != nil {
		fmt.Println(err)
	}

	return Layout(
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
