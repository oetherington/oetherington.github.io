package main

import (
	"fmt"
	"os/exec"
	"sync"
	"time"

	. "github.com/oetherington/smetana"
)

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
	url := fmt.Sprintf("%s%s", baseUrl, pageUrl)
	modified, err := time.Parse(time.RFC3339, string(stdout))
	if err != nil {
		return SitemapLocation{}, err
	}
	return SitemapLocationMod(url, modified), nil
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

func getSitemap(
	baseUrl string,
	staticRoutes []StaticRoute,
	articles []ArticleInfo,
) (Sitemap, error) {
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
