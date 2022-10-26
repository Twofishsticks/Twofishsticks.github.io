// models.article.go

package main

import (
	"errors"
)

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// For this demo, we're storing the article list in memory
// In a real application, this list will most likely be fetched
// from a database or from static files
var articleList = []article{
	article{ID: 1, Title: "CSCE 247: LinkedIn - In Console", Content: "With a group of 3 other people, we worked together to build a functional console program in Java utilizing json to access saved information. "},
	article{ID: 2, Title: "CSCE 490: Website using GoLang", Content: "Every single aspect of this website was built using Go, using gin as the framework. What you see here is the final product!"},
}

// Return a list of all the articles
func getAllArticles() []article {
	return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}
