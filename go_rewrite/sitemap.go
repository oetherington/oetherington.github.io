package main

import (
	"fmt"
	"os/exec"
	"sync"

	. "github.com/oetherington/smetana"
)

type SitemapLocation struct {
	url      string
	modified string
}

func newSitemapLocation(
	baseUrl string,
	pageUrl string,
	srcFile string,
) (SitemapLocation, error) {
	args := []string{
		"log",
		"-1",
		"--pretty=format:%aI",
		"--follow",
		"--",
		srcFile,
	}
	result := exec.Command("git", args...)
	stdout, err := result.Output()
	if err != nil {
		return SitemapLocation{}, err
	}
	return SitemapLocation{
		url:      fmt.Sprintf("%s%s", baseUrl, pageUrl),
		modified: string(stdout),
	}, nil
}

type StaticRoute struct {
	url     string
	srcFile string
}

func countLocations(staticRoutes []StaticRoute, articles []ArticleInfo) int {
	count := len(staticRoutes)
	for _, article := range articles {
		if article.Published {
			count++
		}
	}
	return count
}

func getSitemapLocations(
	baseUrl string,
	staticRoutes []StaticRoute,
	articles []ArticleInfo,
) ([]SitemapLocation, error) {
	count := countLocations(staticRoutes, articles)
	locations := make([]SitemapLocation, count)
	errors := make([]error, count)

	var wg sync.WaitGroup

	for i, route := range staticRoutes {
		wg.Add(1)
		go func(i int, route StaticRoute) {
			loc, err := newSitemapLocation(baseUrl, route.url, route.srcFile)
			locations[i] = loc
			errors[i] = err
			wg.Done()
		}(i, route)
	}

	for i, article := range articles {
		if !article.Published {
			continue
		}
		wg.Add(1)
		go func(i int, article ArticleInfo) {
			filePath := fmt.Sprintf("./articles/%s.md", article.Path)
			loc, err := newSitemapLocation(baseUrl, article.Path, filePath)
			index := len(staticRoutes) + i
			locations[index] = loc
			errors[index] = err
			wg.Done()
		}(i, article)
	}

	wg.Wait()

	for _, err := range errors {
		if err != nil {
			return nil, err
		}
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
