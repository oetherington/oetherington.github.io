package main

import (
	. "github.com/oetherington/smetana"
)

type Palette struct {
	background  Color
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

/*
New palette?:
	#282a36
	#686868
	#FF5C57
	#5AF78E
	#F3F99D
	#57C7FF
	#FF6AC1
	#9AEDFE
	#EFF0EB
*/

func createPalette() Palette {
	return Palette{
		background:  Hex("#282a36"),
		black:       Hex("#000"),
		offBlack:    Hex("#222"),
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
