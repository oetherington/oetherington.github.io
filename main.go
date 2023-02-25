package main

import (
	"fmt"
	"log"
	"os"

	"github.com/oetherington/smetana"
	"github.com/otiai10/copy"
)

const PUBLIC_DIR = "./public"
const OUTPUT_DIR = "./docs"

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
	css := smetana.RenderCss(stylesheet)
	return writeString(css, targetName)
}

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
	fontSrcs := []string{
		"UnifontLatin.ttf",
		"UnifontLatin.woff",
		"UnifontLatin.woff2",
	}
	for _, fontSrc := range fontSrcs {
		src := fontSrcDir + fontSrc
		dest := fontOutDir + "/" + fontSrc
		if err := copy.Copy(src, dest); err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Println("Loading articles")
	articleInfo, err := loadArticleInfo()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Generating palette")
	palette := createPalette()

	fmt.Println("Compiling styles")
	styles := createStyles(palette)
	if err := generateCss(styles, "css/styles.css"); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Compiling index")
	index := Layout(palette, "", Index(articleInfo))
	if err := generateHtml(index, "index.html"); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Compiling articles...")
	for _, article := range articleInfo {
		if !article.Published {
			continue
		}
		fmt.Println("...", article.Name)
		articleHtml := MdArticle(palette, article)
		fileName := fmt.Sprintf("%s.html", article.Path)
		if err = generateHtml(articleHtml, fileName); err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Println("Compiling sitemap")
	baseUrl := "https://www.etherington.io/"
	staticRoutes := []StaticRoute{
		{"", "./index.go"},
	}
	locations, err := getSitemapLocations(baseUrl, staticRoutes, articleInfo)
	if err != nil {
		log.Fatalln(err)
	}
	sitemap := Sitemap(locations)
	if err := generateHtml(sitemap, "sitemap.xml"); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Done! 🎉")
}
