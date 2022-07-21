package cache_service

import (
	"encoding/json"
	"fmt"
	"github.com/WeCanRun/gin-blog/global/constants"
	"github.com/WeCanRun/gin-blog/internal/model"
	"github.com/WeCanRun/gin-blog/pkg/logging"
)

const (
	TAG_ID_KEY     = constants.CACHE_TAG + ":id:%v"
	TAG_CAHCE_TIME = 30
)

// 拿出缓存
func GetTagCacheById(tagId uint) (model.Tag, error) {
	tag := model.Tag{}
	key := fmt.Sprintf(TAG_ID_KEY, tagId)
	bytes, err := Get(key)
	if err != nil {
		logging.Error("GetTagCache fail, err:%v", err)
		return tag, err
	}
	if err := json.Unmarshal(bytes, &tag); err != nil {
		logging.Error("GetTagCache | Unmarshal fail, err:%v", err)
	}
	return tag, err
}

// 设置缓存
func SetTagCacheById(tag model.Tag) error {
	key := fmt.Sprintf(TAG_ID_KEY, tag.ID)
	err := SetWithExpire(key, tag, TAG_CAHCE_TIME)
	if err != nil {
		logging.Error("SetTagCache#SetWithExpire fail, err:%v ", err)
	}
	return err
}

func DeleteTagCacheById(id uint) error {
	key := fmt.Sprintf(TAG_ID_KEY, id)
	exists := Exists(key)
	if exists {
		_, err := Delete(key)
		return err
	}

	return nil
}
