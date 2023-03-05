package main

import (
	"fmt"
	"log"
	"os"

	"github.com/oetherington/smetana"
	"github.com/otiai10/copy"
)

var CONTEXT = smetana.NewSmetanaWithPalettes(smetana.Palettes{
	"dark":  createDarkPalette(),
	"light": createLightPalette(),
})

const PUBLIC_DIR = "./public"
const OUTPUT_DIR = "./build"

func writeString(value string, targetName string) error {
	targetPath := fmt.Sprintf("%s/%s", OUTPUT_DIR, targetName)

	file, err := os.Create(targetPath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(value)
	return err
}

func generateHtml(node smetana.Node, targetName string) error {
	html := smetana.RenderHtml(node)
	return writeString(html, targetName)
}

func generateCss(stylesheet smetana.StyleSheet, targetName string) error {
	css := smetana.RenderCss(stylesheet, CONTEXT.Palettes["dark"])
	return writeString(css, targetName)
}

func generateSitemap(sitemap smetana.Sitemap, targetName string) error {
	css := smetana.RenderSitemap(sitemap)
	return writeString(css, targetName)
}

const BASE_URL = "https://www.etherington.io/"

func main() {
	fmt.Println("Removing old build")
	if err := os.RemoveAll(OUTPUT_DIR); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Copying public files")
	if err := copy.Copy(PUBLIC_DIR, OUTPUT_DIR); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Copying fonts")
	fontSrcDir := "./font/"
	fontOutDir := fmt.Sprintf("%s/css", OUTPUT_DIR)
	if err := os.MkdirAll(fontOutDir, 0777); err != nil {
		log.Fatalln(err)
	}
	fontFormats := []string{"ttf", "woff", "woff2"}
	for _, fontFormat := range fontFormats {
		src := fontSrcDir + "UnifontLatin." + fontFormat
		dest := fontOutDir + "/UnifontLatin." + fontFormat
		if err := copy.Copy(src, dest); err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Println("Loading articles")
	articleInfo, err := loadArticleInfo()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Compiling styles")
	CONTEXT.Styles = createStyles()
	styles := CONTEXT.RenderStyles()
	for name, css := range styles {
		path := fmt.Sprintf("css/%s.css", name)
		err = writeString(css, path)
		if err != nil {
			log.Fatalln(err)
		}
	}
	if err := generateCss(createPrintStyles(), "css/print.css"); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Compiling index")
	index := Layout("", Index(articleInfo))
	if err := generateHtml(index, "index.html"); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Compiling articles...")
	for _, article := range articleInfo {
		if !article.Published {
			continue
		}
		fmt.Println("...", article.Name)
		articleHtml := MdArticle(article)
		fileName := fmt.Sprintf("%s.html", article.Path)
		if err = generateHtml(articleHtml, fileName); err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Println("Compiling sitemap")
	staticRoutes := []StaticRoute{
		{"", "./index.go"},
	}
	sitemap, err := getSitemap(BASE_URL, staticRoutes, articleInfo)
	if err != nil {
		log.Fatalln(err)
	}
	if err := generateSitemap(sitemap, "sitemap.xml"); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Done! 🎉")
}
