package cache_service

import (
	"encoding/json"
	"fmt"
	"github.com/WeCanRun/gin-blog/global/constants"
	"github.com/WeCanRun/gin-blog/internal/model"
	"github.com/WeCanRun/gin-blog/pkg/logging"
)

const (
	ARTICLE_ID_KEY     = constants.CACHE_ARTICLE + ":id:%v"
	ARTICLE_CACHE_TIME = 30
)

func GetArticleCacheById(id uint) (model.Article, error) {
	article := model.Article{}
	key := fmt.Sprintf(ARTICLE_ID_KEY, id)
	data, err := Get(key)
	if err != nil {
		return article, err
	}
	if err := json.Unmarshal(data, &article); err != nil {
		logging.Error("GetArticleCacheById Unmarshal fail, err: {}", err)
	}
	return article, err
}

func SetArticleCacheById(article model.Article) error {
	key := fmt.Sprintf(ARTICLE_ID_KEY, article.ID)
	return SetWithExpire(key, article, ARTICLE_CACHE_TIME)
}

func DeleteArticleCacheById(id uint) error {
	key := fmt.Sprintf(ARTICLE_ID_KEY, id)
	exists := Exists(key)
	if exists {
		_, err := Delete(key)
		return err
	}

	return nil
}
