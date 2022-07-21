package model

import (
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/WeCanRun/gin-blog/pkg/setting"
	"github.com/go-playground/assert/v2"
	"github.com/jinzhu/gorm"
	"testing"
	"time"
)

func init() {
	setting.Setup("../../conf/app.yaml")
	logging.Setup()
	Setup()
}

func TestAddData(t *testing.T) {
	for i := 0; i < 5; i++ {
		TestAddArticle(t)
		TestAddTag(t)
	}
}

func TestAddArticle(t *testing.T) {
	article := Article{
		Model: gorm.Model{
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
	newArticle, _ := GetArticleById(1)
	assert.Equal(t, newArticle.ID, 1)
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
