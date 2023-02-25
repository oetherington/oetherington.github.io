package main

import (
	"fmt"
	. "github.com/oetherington/smetana"
	"os/exec"
)

type SitemapLocation struct {
	url      string
	modified string
}

func getFileModifiedDate(filePath string) (string, error) {
	args := []string{
		"log",
		"-1",
		"--pretty=format:%aI",
		"--follow",
		"--",
		filePath,
	}
	result := exec.Command("git", args...)
	stdout, err := result.Output()
	if err != nil {
		return "", err
	}
	return string(stdout), nil
}

type StaticRoute struct {
	url     string
	srcFile string
}

func getSitemapLocations(
	baseUrl string,
	staticRoutes []StaticRoute,
	articles []ArticleInfo,
) ([]SitemapLocation, error) {
	locations := []SitemapLocation{}

	for _, route := range staticRoutes {
		modified, err := getFileModifiedDate(route.srcFile)
		if err != nil {
			return nil, err
		}
		locations = append(locations, SitemapLocation{
			url:      fmt.Sprintf("%s%s", baseUrl, route.url),
			modified: modified,
		})
	}

	for _, article := range articles {
		if !article.Published {
			continue
		}
		filePath := fmt.Sprintf("./articles/%s.md", article.Path)
		modified, err := getFileModifiedDate(filePath)
		if err != nil {
			return nil, err
		}
		locations = append(locations, SitemapLocation{
			url:      fmt.Sprintf("%s%s", baseUrl, article.Path),
			modified: modified,
		})
	}

	return locations, nil
}

type SitemapNode struct {
	locations []SitemapLocation
}

func (node SitemapNode) ToHtml(builder *Builder) {
	builder.Buf.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n")
	builder.Buf.WriteString("<urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\">\n")
	for _, loc := range node.locations {
		builder.Buf.WriteString(fmt.Sprintf(
			" <url>\n  <loc>%s</loc>\n  <lastmod>%s</lastmod>\n </url>\n",
			loc.url,
			loc.modified,
		))
	}
	builder.Buf.WriteString("</urlset>\n")
}

func Sitemap(locations []SitemapLocation) SitemapNode {
	return SitemapNode{locations}
}
