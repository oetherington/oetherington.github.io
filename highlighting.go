package main

import (
	"fmt"
	"github.com/alecthomas/chroma/v2"
	. "github.com/oetherington/smetana"
)

func fgBgStyle(fg string, bg string) CssProps {
	//nolint:govet
	return CssProps{
		{"color", PaletteValue(fg)},
		{"background-color", PaletteValue(bg)},
	}
}

func bgStyle(bg string) CssProps {
	return CssProps{
		//nolint:govet
		{"background-color", PaletteValue(bg)},
	}
}

func fgStyle(fg string) CssProps {
	return CssProps{
		//nolint:govet
		{"color", PaletteValue(fg)},
	}
}

func italicStyle() CssProps {
	//nolint:govet
	return CssProps{{"font-style", "italic"}}
}

func boldStyle() CssProps {
	//nolint:govet
	return CssProps{{"font-weight", "bold"}}
}

type HighlightTheme map[chroma.TokenType]CssProps

func (theme HighlightTheme) GetForToken(token chroma.TokenType) CssProps {
	if value := theme[token]; value != nil {
		return value
	}

	token = token.SubCategory()
	if value := theme[token]; value != nil {
		return value
	}

	token = token.Category()
	if value := theme[token]; value != nil {
		return value
	}

	return CssProps{}
}

func createHighlightTheme() HighlightTheme {
	return HighlightTheme{
		chroma.Error:               fgBgStyle("white", "red"),
		chroma.Background:          bgStyle("black"),
		chroma.Keyword:             fgStyle("darkBlue"),
		chroma.KeywordNamespace:    fgStyle("pink"),
		chroma.Name:                fgStyle("white"),
		chroma.NameAttribute:       fgStyle("green"),
		chroma.NameClass:           fgStyle("green"),
		chroma.NameConstant:        fgStyle("darkBlue"),
		chroma.NameDecorator:       fgStyle("green"),
		chroma.NameException:       fgStyle("green"),
		chroma.NameFunction:        fgStyle("green"),
		chroma.NameOther:           fgStyle("green"),
		chroma.NameTag:             fgStyle("pink"),
		chroma.Literal:             fgStyle("lightBlue"),
		chroma.LiteralDate:         fgStyle("yellow"),
		chroma.LiteralString:       fgStyle("yellow"),
		chroma.LiteralStringEscape: fgStyle("lightBlue"),
		chroma.LiteralNumber:       fgStyle("lightBlue"),
		chroma.Operator:            fgStyle("pink"),
		chroma.Punctuation:         fgStyle("white"),
		chroma.Comment:             fgStyle("grey"),
		chroma.GenericDeleted:      fgStyle("red"),
		chroma.GenericInserted:     fgStyle("green"),
		chroma.GenericSubheading:   fgStyle("grey"),
		chroma.GenericEmph:         italicStyle(),
		chroma.GenericStrong:       boldStyle(),
		chroma.Text:                fgStyle("white"),
	}
}

func renderHighlightingCss() StyleSheet {
	styles := NewStyleSheet()
	theme := createHighlightTheme()
	for token, className := range chroma.StandardTypes {
		switch token {
		case chroma.Background, chroma.PreWrapper, chroma.Text:
			continue
		}
		style := theme.GetForToken(token)
		selector := fmt.Sprintf(".chroma .%s", className)
		styles.AddBlock(selector, style)
	}
	return styles
}
