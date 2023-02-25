package main

import (
	"fmt"
	"time"

	. "github.com/oetherington/smetana"
)

const DEFAULT_TITLE = "Ollie Etherington"

func formatTitle(title string) string {
	if len(title) < 1 || title == DEFAULT_TITLE {
		return DEFAULT_TITLE
	}
	return fmt.Sprintf("%s | %s", title, DEFAULT_TITLE)
}

func Layout(palette Palette, title string, content Node) HtmlNode {
	now := time.Now()
	return Html(
		Head(
			Charset(""),
			Link(Attrs{
				"rel":   "stylesheet",
				"href":  "/css/print.css",
				"media": "print",
			}),
			Link(Attrs{
				"rel":   "stylesheet",
				"href":  "/css/light.css",
				"media": "screen and (prefers-color-scheme: light)",
			}),
			Link(Attrs{
				"rel":   "stylesheet",
				"href":  "/css/dark.css",
				"media": "screen and (prefers-color-scheme: dark)",
			}),
			Equiv("x-ua-compatible", "ie=edge"),
			Title(formatTitle(title)),
			Description("Ollie Etherington"),
			Author("Ollie Etherington"),
			Keywords("ollie,oliver,etherington,london,oxford,programmer"),
			Viewport(""),
			Meta("msapplication-TileColor", palette.pink.ToCssColor()),
			Meta("theme-color", palette.white.ToCssColor()),
			Link(Attrs{
				"rel":  "shortcut icon",
				"href": "/favicon.ico?",
				"type": "image/x-icon",
			}),
			Link(Attrs{
				"rel":   "apple-touch-icon",
				"href":  "/apple-touch-icon.png",
				"sizes": "180x180",
			}),
			Link(Attrs{
				"rel":   "icon",
				"href":  "/favicon-32x32.png",
				"sizes": "32x32",
				"type":  "image/png",
			}),
			Link(Attrs{
				"rel":   "icon",
				"href":  "/favicon-16x16.png",
				"sizes": "16x16",
				"type":  "image/png",
			}),
			Link(Attrs{
				"rel":   "mask-icon",
				"href":  "/safari-pinned-tab.svg",
				"color": palette.grey.ToCssColor(),
			}),
			LinkHref("manifest", "/site.webmanifest"),
		),
		Body(
			Script("0"), // Prevents FOUC in Firefox
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
				P(fmt.Sprintf(
					"Copyright &#169; 2009-%d Ollie Etherington.",
					now.Year(),
				)),
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
