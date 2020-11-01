package models

import (
	"github.com/go-playground/assert/v2"
	"github.com/jinzhu/gorm"
	"testing"
	"time"
)

func TestAddArticle(t *testing.T) {
	article := Article{
		Model: gorm.Model{
			ID:        2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		},
		TagId:     1,
		Title:     "test",
		Desc:      "test",
		Content:   "test",
		CreatedBy: "test",
		State:     1,
	}
	err := AddArticle(article)
	if err != nil {
		t.Log("err: ", err)
	}
	newArticle, _ := GetArticleById(article.ID)
	assert.Equal(t, newArticle.ID, article.ID)
}

func TestGetArticleById(t *testing.T) {
	article, err := GetArticleById(1)
	if err != nil {
		t.Log("err:", err)
	}
	t.Log(article)
}

func TestGetArticleByTitle(t *testing.T) {
	articles, err := GetArticleByTitle("test", 1)
	if err != nil {
		t.Log(err)
	}
	for _, article := range articles {
		t.Log(article)
	}
}

func TestGetArticleByTagId(t *testing.T) {
	articles, err := GetArticleByTagId(1, 1)
	if err != nil {
		t.Log(err)
	}
	for _, article := range articles {
		t.Log(article)
	}
}

func TestExitArticleWithTitle(t *testing.T) {
	isExit := ExitArticleWithTitle("test")
	assert.Equal(t, isExit, true)
}

func TestGetArticleTotal(t *testing.T) {
	total := GetArticleTotal(Article{
		State: 1,
	})
	t.Log(total)
}

func TestEditArticle(t *testing.T) {
	article := Article{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Content:   "edit",
		UpdatedBy: "edit",
		State:     0,
	}
	if err := EditArticle(article); err != nil {
		t.Log(err)
	}
	assert.Equal(t, article.Content, "edit")
}

func TestDeleteArticle(t *testing.T) {
	if err := DeleteArticle(2); err != nil {
		t.Log(err)
	} else {
		t.Log("Success")
	}

}

func TestGetArticles(t *testing.T) {
	articles, err := GetArticles(0, 10)
	if err != nil {
		t.Log(err)
	}
	for _, article := range articles {
		t.Log(article)
	}
}
