package main

import (
	"encoding/json"
	"io"
	"os"
)

type ArticleInfo struct {
	Name      string `json:"name"`
	Path      string `json:"path"`
	Published bool   `json:"published"`
}

func loadArticleInfo() ([]ArticleInfo, error) {
	file, err := os.Open("articles/manifest.json")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var articles []ArticleInfo
	err = json.Unmarshal(bytes, &articles)
	if err != nil {
		return nil, err
	}

	return articles, nil
}
