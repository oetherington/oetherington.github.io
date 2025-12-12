package main

import (
	"fmt"
	. "github.com/oetherington/smetana"
)

func Project(name string, href string, description string) DomNode {
	return P(
		AHref(href, name),
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
		),
		Project(
			"Smetana",
			"https://github.com/oetherington/smetana",
			"An HTML and CSS generator for Go",
		),
		Project(
			"Pongboot",
			"https://github.com/oetherington/pongboot",
			"Pong squeezed into an x86 bootloader",
		),
		Project(
			"Bluesky Embed React",
			"/bluesky-embed-react",
			"Embed posts and feeds from Bluesky in React",
		),
		Project(
			"JSkorost",
			"https://github.com/oetherington/jskorost",
			"A fast single-header JSON parser",
		),
		Project(
			"Chrd",
			"https://www.etherington.xyz/chrd/",
			"A LaTeX inspired chord chart editor",
		),
		Project(
			"Glinka",
			"https://glinka.io",
			"A fast Typescript compiler written in Zig",
		),
		Project(
			"ws",
			"https://github.com/oetherington/ws",
			"A simple CLI workspace manager",
		),
		Project(
			"ChessPieceSwitcher",
			"https://github.com/oetherington/ChessPieceSwitcher",
			"An ill-advised web extension",
		),
	)
}
