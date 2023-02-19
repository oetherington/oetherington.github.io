package main

import (
	. "github.com/oetherington/smetana"
)

type Palette struct {
	black       Color
	offBlack    Color
	white       Color
	offWhite    Color
	darkGrey    Color
	mediumGrey  Color
	lightGrey   Color
	darkYellow  Color
	lightYellow Color
	darkPink    Color
	lightPink   Color
	darkBlue    Color
	lightBlue   Color
	purple      Color
	green       Color
	red         Color
}

func createPalette() Palette {
	return Palette{
		black:       Hex("#000"),
		offBlack:    Hex("#111"),
		white:       Hex("#fff"),
		offWhite:    Hex("#f8f8f2"),
		darkGrey:    Hex("#75715e"),
		mediumGrey:  Hex("#a8a8a2"),
		lightGrey:   Hex("#cdcdcd"),
		darkYellow:  Hex("#e6db74"),
		lightYellow: Hex("#cdcd00"),
		darkPink:    Hex("#f0f"),
		lightPink:   Hex("#cd00cd"),
		darkBlue:    Hex("#66d9ef"),
		lightBlue:   Hex("#00cdcd"),
		purple:      Hex("#ae81ff"),
		green:       Hex("#a6e22e"),
		red:         Hex("#f92672"),
	}
}
