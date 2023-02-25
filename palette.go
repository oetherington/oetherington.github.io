package main

import (
	. "github.com/oetherington/smetana"
)

type Palette struct {
	background Color
	black      Color
	white      Color
	grey       Color
	yellow     Color
	pink       Color
	darkBlue   Color
	lightBlue  Color
	green      Color
	red        Color
}

func createDarkPalette() Palette {
	return Palette{
		background: Hex("#282a36"),
		black:      Hex("#000"),
		white:      Hex("#EFF0EB"),
		grey:       Hex("#686868"),
		yellow:     Hex("#F3F99D"),
		pink:       Hex("#FF6AC1"),
		darkBlue:   Hex("#57C7FF"),
		lightBlue:  Hex("#9AEDFE"),
		green:      Hex("#5AF78E"),
		red:        Hex("#FF5C57"),
	}
}

func createLightPalette() Palette {
	return Palette{
		background: Hex("#fdf6e3"),
		black:      Hex("#eee8d5"),
		white:      Hex("#839496"),
		grey:       Hex("#073642"),
		yellow:     Hex("#b58900"),
		pink:       Hex("#d33682"),
		darkBlue:   Hex("#268bd2"),
		lightBlue:  Hex("#2aa198"),
		green:      Hex("#859900"),
		red:        Hex("#dc322f"),
	}
}
