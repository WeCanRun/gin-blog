package model

import (
	"github.com/go-playground/assert/v2"
	"github.com/jinzhu/gorm"
	"testing"
)

func TestAddTag(t *testing.T) {
	err := AddTag(Tag{
		Name:      "test",
		CreatedBy: "test",
		State:     1,
	})
	if err != nil {
		t.Log("err", err)
	}
	t.Log("success")
}

func TestEditTag(t *testing.T) {
	err := EditTag(Tag{
		Model:     gorm.Model{ID: 11},
		Name:      "test2",
		UpdatedBy: "test2",
		State:     1,
	})
	if err != nil {
		t.Log(err)
	}
	t.Log("success")
}

func TestGetTags(t *testing.T) {
	tags, err := GetTags(1, 10)
	if err != nil {
		t.Log(err)
	}
	t.Log("success", tags)
}

func TestDeleteTag(t *testing.T) {
	err := DeleteTag(11)
	if err != nil {
		t.Log("err", err)
	}
	t.Log("success")
}

func TestGetTagTotal(t *testing.T) {
	tags, err := GetTagsByName("test")
	if err != nil {
		t.Log("err", err)
	}
	t.Log("tags", tags)
	t.Log("total", len(tags))
}

func TestExitTagWithName(t *testing.T) {
	is := ExitTagWithName("test")
	assert.Equal(t, is, true)
}

func TestGetTagById(t *testing.T) {
	tag, err := GetTagById(12)
	if err != nil {
		t.Log("Err: ", err)
	}
	t.Log(tag)
}

func TestGetTagsByIds(t *testing.T) {
	names, err := GetTagsByIds([]uint{11, 12, 13})
	if err != nil {
		assert.Equal(t, true, gorm.IsRecordNotFoundError(err))
	}

	t.Log(names)

}
