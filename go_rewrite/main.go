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
	html := smetana.RenderCss(stylesheet)
	return writeString(html, targetName)
}

func main() {
	err := os.RemoveAll(OUTPUT_DIR)
	if err != nil {
		log.Fatalln(err)
	}

	err = copy.Copy(PUBLIC_DIR, OUTPUT_DIR)
	if err != nil {
		log.Fatalln(err)
	}

	palette := createPalette()

	styles := createStyles(palette)
	err = generateCss(styles, "css/styles.css")
	if err != nil {
		log.Fatalln(err)
	}

	index := Layout(palette, "", Index())
	err = generateHtml(index, "index.html")
	if err != nil {
		log.Fatalln(err)
	}
}
