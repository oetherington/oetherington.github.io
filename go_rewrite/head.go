package main

import (
	"fmt"
	. "github.com/oetherington/smetana"
)

const DEFAULT_TITLE = "Ollie Etherington"

func formatTitle(title string) string {
	if len(title) < 1 || title == DEFAULT_TITLE {
		return DEFAULT_TITLE
	}
	return fmt.Sprintf("%s | %s", title, DEFAULT_TITLE)
}

func createHead(palette Palette, title string) DomNode {
	return Head(
		Charset(""),
		Equiv("x-ua-compatible", "ie=edge"),
		Title(formatTitle(title)),
		Description("Ollie Etherington"),
		Author("Ollie Etherington"),
		Keywords("ollie,oliver,etherington,london,oxford,programmer"),
		Viewport(""),
		Meta("msapplication-TileColor", palette.lightPink.ToCssColor()),
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
			"color": palette.darkGrey.ToCssColor(),
		}),
		LinkHref("manifest", "/site.webmanifest"),
		LinkHref("stylesheet", "css/styles.css"),
	)
}
