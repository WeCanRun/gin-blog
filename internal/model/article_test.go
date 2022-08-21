package model

import (
	"context"
	"github.com/WeCanRun/gin-blog/global"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/WeCanRun/gin-blog/pkg/setting"
	"github.com/go-playground/assert/v2"
	"github.com/jinzhu/gorm"
	"testing"
)

func init() {
	global.Setting = setting.Setup("../../conf/app-test.yaml")
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
		TagId:     1,
		Title:     "test",
		Desc:      "test",
		Content:   "test",
		CreatedBy: "test",
		State:     1,
	}
	err := AddArticle(context.Background(), article)
	if err != nil {
		t.Log("err: ", err)
	}
}

func TestGetArticleById(t *testing.T) {
	article, err := GetArticleById(context.Background(), 1)
	if err != nil {
		t.Log("err:", err)
	}
	t.Log(article)
}

func TestGetArticleByTitle(t *testing.T) {
	articles, err := GetArticleByTitle(context.Background(), "test", 1)
	if err != nil {
		t.Log(err)
	}
	for _, article := range articles {
		t.Log(article)
	}
}

func TestGetArticleByTagId(t *testing.T) {
	articles, err := GetArticleByTagId(context.Background(), 1, 1)
	if err != nil {
		t.Log(err)
	}
	for _, article := range articles {
		t.Log(article)
	}
}

func TestExitArticleWithTitle(t *testing.T) {
	isExit := ExitArticleWithTitle(context.Background(), "test")
	assert.Equal(t, isExit, true)
}

func TestGetArticleTotal(t *testing.T) {
	total := GetArticleTotal(context.Background(), Article{
		State: 1,
	})
	t.Log(total)
}

func TestEditArticle(t *testing.T) {
	article := Article{
		Model: gorm.Model{
			ID: 11,
		},
		Content:   "edit",
		UpdatedBy: "edit",
		State:     0,
	}
	if err := EditArticle(context.Background(), article); err != nil {
		t.Log(err)
	}
	assert.Equal(t, article.Content, "edit")
	t.Log(article.UpdatedAt)
}

func TestDeleteArticle(t *testing.T) {
	if err := DeleteArticle(context.Background(), 6); err != nil {
		t.Log(err)
	} else {
		t.Log("Success")
	}

}

func TestGetArticles(t *testing.T) {
	articles, err := GetArticles(context.Background(), 0, 10)
	if err != nil {
		t.Log(err)
	}
	for _, article := range articles {
		t.Log(article)
	}
}
