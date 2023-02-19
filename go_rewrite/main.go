package main

import (
	"fmt"
	. "github.com/oetherington/smetana"
)

func main() {
	palette := createPalette()
	styles := createStyles(palette)

	page := Layout(palette, "", nil)

	css := RenderCss(styles)
	html := RenderHtml(page)

	fmt.Println(css)
	fmt.Println(html)
}
