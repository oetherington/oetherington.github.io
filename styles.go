package main

import (
	"fmt"
	. "github.com/oetherington/smetana"
)

func createStyles(palette Palette, prefixClass string) (StyleSheet, error) {
	mdCss, err := renderMarkdownCss(palette, prefixClass)
	if err != nil {
		return StyleSheet{}, err
	}
	styles := NewStyleSheet(
		StylesCss(mdCss),
		StylesFontFace(
			"Unifont",
			"UnifontLatin.woff2",
			"UnifontLatin.woff",
			"UnifontLatin.ttf",
		),
		StylesBlock("body", CssProps{
			"font-family": "Unifont, monospace",
			"background":  palette.background,
			"color":       palette.white,
			"width":       Perc(90),
			"margin":      "auto",
		}),
		StylesCss(`
			@media only screen and (max-width: 600px) {
				body { width: 100%; }
			}
		`),
		// TODO: These 2 should use OrderedCssProps
		StylesBlock("hr", CssProps{
			"border": "none",
		}),
		StylesBlock("hr", CssProps{
			"border-top": fmt.Sprintf(
				"1px dashed %s",
				palette.grey.ToCssColor(),
			),
		}),
		StylesBlock("h2", CssProps{
			"color": palette.yellow,
		}),
		StylesBlock("h4", CssProps{
			"color": palette.pink,
		}),
		StylesBlock("a", CssProps{
			"color":      palette.lightBlue,
			"transition": "color 0.2s",
		}),
		StylesBlock("a:hover", CssProps{
			"color": palette.pink,
		}),
		StylesBlock("a svg", CssProps{
			"fill":       palette.lightBlue,
			"transition": "fill 0.2s",
		}),
		StylesBlock("a svg:hover", CssProps{
			"fill": palette.pink,
		}),
		StylesBlock("h1 a", CssProps{
			"color":           palette.white,
			"text-decoration": "none !important",
		}),
		StylesBlock("h1 a:hover", CssProps{
			"color": palette.lightBlue,
		}),
		StylesBlock("h2 a", CssProps{
			"color":           palette.yellow,
			"text-decoration": "none !important",
		}),
		StylesBlock("h2 a:hover", CssProps{
			"color": palette.pink,
		}),
		StylesBlock("svg", CssProps{
			"margin": "0 0.3em -0.2em 0.3em",
		}),
		StylesBlock(".link-icons", CssProps{
			"float":       "right",
			"font-size":   EM(2),
			"display":     "flex",
			"align-items": "center",
			"gap":         EM(0.2),
		}),
		StylesBlock(".link-icons svg", CssProps{
			"fill":  palette.yellow,
			"width": EM(1.1),
		}),
		StylesBlock(".inline-icon svg", CssProps{
			"width": EM(1),
		}),
		StylesBlock("table", CssProps{
			"text-align":     "left",
			"border-spacing": "1.5em 0.1em",
		}),
		StylesBlock("th", CssProps{
			"border-bottom": fmt.Sprintf(
				"1px dashed %s",
				palette.grey.ToCssColor(),
			),
		}),
		StylesCss(`
			@keyframes blink {
				0% { opacity: 0%; }
				50% { opacity: 100%; }
			}
		`),
		StylesBlock(".cursor", CssProps{
			"animation": "blink 1.5s steps(1,end) 0s infinite none",
		}),
		StylesBlock("h2 .cursor", CssProps{
			"background": palette.yellow,
		}),
		StylesBlock("h2 a:hover .cursor", CssProps{
			"background": palette.pink,
		}),
		StylesBlock(".content-full", CssProps{
			"padding": "0 2ch",
		}),
		StylesBlock(".content", CssProps{
			"padding":       "0 2ch",
			"max-width":     CH(80),
			"overflow-wrap": "break-word",
		}),
		StylesBlock(".content h1:not(:first-child)", CssProps{
			"margin-top": EM(2),
		}),
		StylesBlock(".center", CssProps{
			"text-align": "center",
		}),
		StylesBlock(".centered", CssProps{
			"margin": "0 auto",
		}),
		StylesBlock(".tall", CssProps{
			"line-height": EM(1.5),
		}),
		StylesBlock(".footer", CssProps{
			"font-size":   EM(0.8),
			"line-height": Perc(90),
			"text-align":  "center",
			"margin":      "12ch auto 0 auto",
		}),
		StylesCss(`
			@media only screen and (max-width: 600px) {
				.no-mobile { display: none; }
			}
		`),
		StylesBlock(".todo", CssProps{
			"background":    palette.black,
			"color":         palette.yellow,
			"padding":       EM(0.15),
			"border-radius": EM(0.2),
		}),
		StylesBlock("code", CssProps{
			"background":    palette.black,
			"font-family":   "Unifont, monospace",
			"padding":       EM(0.15),
			"border-radius": EM(0.2),
		}),
		StylesBlock("pre", CssProps{
			"background":    palette.black,
			"padding":       EM(0.5),
			"border-radius": EM(1),
			"white-space":   "pre-wrap",
		}),
		StylesBlock("pre code", CssProps{
			"padding":     "0",
			"white-space": "0",
		}),
		StylesBlock("ul", CssProps{
			"list-style":      `"- "`,
			"list-style-type": `"- "`,
		}),
		StylesBlock(".code-lit", CssProps{
			"color": palette.yellow,
		}),
		StylesBlock(".code-fn", CssProps{
			"color": palette.green,
		}),
	)
	return styles, nil
}
