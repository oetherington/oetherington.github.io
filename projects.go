package main

import (
	"fmt"
	. "github.com/oetherington/smetana"
)

func Project(name string, href string, description string, extra Node) DomNode {
	return P(
		AHref(href, name),
		extra,
		Span(fmt.Sprintf("&nbsp; %s", description)),
	)
}

func Projects() DomNode {
	return Div(
		ClassName("content"),
		Project(
			"Spockfish",
			"https://spockfish.com",
			"A 3D chess engine and interface",
			nil,
		),
		Project(
			"Smetana",
			"https://github.com/oetherington/smetana",
			"An HTML and CSS generator for Go",
			nil,
		),
		Project(
			"JSkorost",
			"https://github.com/oetherington/jskorost",
			"A fast single-header JSON parser",
			nil,
		),
		Project(
			"Chrd",
			"https://www.etherington.xyz/chrd/",
			"A LaTeX inspired chord chart editor",
			nil,
		),
		Project(
			"Glinka",
			"https://glinka.io",
			"A fast Typescript compiler written in Zig",
			AHref(
				"https://www.npmjs.com/package/glinka",
				ClassName("inline-icon"),
				NPM_ICON,
			),
		),
		Project(
			"ws",
			"https://github.com/oetherington/ws",
			"A simple CLI workspace manager",
			nil,
		),
		Project(
			"ChessPieceSwitcher",
			"https://github.com/oetherington/ChessPieceSwitcher",
			"An ill-advised web extension",
			Fragment(
				AHref(
					"https://addons.mozilla.org/en-US/firefox/addon/chesspieceswitcher/",
					ClassName("inline-icon"),
					FIREFOX_ICON,
				),
				AHref(
					"https://chrome.google.com/webstore/detail/lichesspieceswitcher/koifaekonhpknbmkbgcacnfcplmhaehn",
					ClassName("inline-icon"),
					CHROME_ICON,
				),
			),
		),
	)
}
