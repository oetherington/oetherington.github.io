package main

import (
	. "github.com/oetherington/smetana"
)

func Layout(palette Palette, title string, content Node) HtmlNode {
	return Html(
		createHead(palette, title),
		Body(
			Div(
				ClassName("content-full"),
				Div(
					ClassName("link-icons"),
					AHref(
						"https://github.com/oetherington",
						GITHUB_ICON,
					),
				),
				H2(
					AHref(
						"/",
						"Ollie Etherington<span class=\"cursor\">&nbsp;</span>",
					),
				),
				H4(
					"Programmer | Oxford, UK",
				),
			),
			Hr(),
			content,
			Div(
				ClassNames("content", "footer"),
				P("Copyright &#169; 2009-2022 Ollie Etherington."),
				P("All content is <a href=\"https://creativecommons.org/licenses/by-nc-sa/4.0/\">CC BY-NC-SA 4.0</a> unless otherwise stated."),
			),
			Br(),
			Script(Attrs{
				"src":  "/instantpage.5.1.0.js",
				"type": "module",
			}),
		),
	)
}
