package main

import (
	. "github.com/oetherington/smetana"
)

var MS_TILE_COLOR = Hex("#FF6AC1")
var THEME_COLOR = Hex("#EFF0EB")
var MASK_COLOR = Hex("#686868")

func createDefaultPalette() Palette {
	return Palette{
		"background": Hex("#282a36"),
		"black":      Hex("#000"),
		"white":      Hex("#EFF0EB"),
		"grey":       Hex("#686868"),
		"yellow":     Hex("#F3F99D"),
		"pink":       Hex("#FF6AC1"),
		"darkBlue":   Hex("#57C7FF"),
		"lightBlue":  Hex("#9AEDFE"),
		"green":      Hex("#5AF78E"),
		"red":        Hex("#FF5C57"),
	}
}
