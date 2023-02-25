package main

import (
	"github.com/alecthomas/chroma/v2"
	"github.com/oetherington/smetana"
)

func convertColor(c smetana.Color) chroma.Colour {
	rgba := c.ToRgba()
	value := (int32(rgba.R) << 16) | (int32(rgba.G) << 8) | int32(rgba.B)
	return chroma.Colour(value)
}

func fgHighlight(fg smetana.Color) chroma.StyleEntry {
	return chroma.StyleEntry{
		Colour: convertColor(fg),
	}
}

func bgHighlight(bg smetana.Color) chroma.StyleEntry {
	return chroma.StyleEntry{
		Background: convertColor(bg),
	}
}

func fgBgHighlight(fg smetana.Color, bg smetana.Color) chroma.StyleEntry {
	return chroma.StyleEntry{
		Colour:     convertColor(fg),
		Background: convertColor(fg),
	}
}

func italicHighlight() chroma.StyleEntry {
	return chroma.StyleEntry{
		Italic: chroma.Yes,
	}
}

func boldHighlight() chroma.StyleEntry {
	return chroma.StyleEntry{
		Bold: chroma.Yes,
	}
}

func createHighlightStyles(palette Palette) (*chroma.Style, error) {
	builder := chroma.NewStyleBuilder("oetherington")
	builder.AddEntry(chroma.Error, fgBgHighlight(palette.white, palette.red))
	builder.AddEntry(chroma.Background, bgHighlight(palette.black))
	builder.AddEntry(chroma.Keyword, fgHighlight(palette.darkBlue))
	builder.AddEntry(chroma.KeywordNamespace, fgHighlight(palette.pink))
	builder.AddEntry(chroma.Name, fgHighlight(palette.white))
	builder.AddEntry(chroma.NameAttribute, fgHighlight(palette.green))
	builder.AddEntry(chroma.NameClass, fgHighlight(palette.green))
	builder.AddEntry(chroma.NameConstant, fgHighlight(palette.darkBlue))
	builder.AddEntry(chroma.NameDecorator, fgHighlight(palette.green))
	builder.AddEntry(chroma.NameException, fgHighlight(palette.green))
	builder.AddEntry(chroma.NameFunction, fgHighlight(palette.green))
	builder.AddEntry(chroma.NameOther, fgHighlight(palette.green))
	builder.AddEntry(chroma.NameTag, fgHighlight(palette.pink))
	builder.AddEntry(chroma.Literal, fgHighlight(palette.lightBlue))
	builder.AddEntry(chroma.LiteralDate, fgHighlight(palette.yellow))
	builder.AddEntry(chroma.LiteralString, fgHighlight(palette.yellow))
	builder.AddEntry(chroma.LiteralStringEscape, fgHighlight(palette.lightBlue))
	builder.AddEntry(chroma.LiteralNumber, fgHighlight(palette.lightBlue))
	builder.AddEntry(chroma.Operator, fgHighlight(palette.pink))
	builder.AddEntry(chroma.Punctuation, fgHighlight(palette.white))
	builder.AddEntry(chroma.Comment, fgHighlight(palette.grey))
	builder.AddEntry(chroma.GenericDeleted, fgHighlight(palette.red))
	builder.AddEntry(chroma.GenericInserted, fgHighlight(palette.green))
	builder.AddEntry(chroma.GenericSubheading, fgHighlight(palette.grey))
	builder.AddEntry(chroma.GenericEmph, italicHighlight())
	builder.AddEntry(chroma.GenericStrong, boldHighlight())
	builder.AddEntry(chroma.Text, fgHighlight(palette.white))
	return builder.Build()
}

