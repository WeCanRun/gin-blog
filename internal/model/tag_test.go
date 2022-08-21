package model

import (
	"context"
	"github.com/go-playground/assert/v2"
	"github.com/jinzhu/gorm"
	"testing"
)

func TestAddTag(t *testing.T) {
	err := AddTag(context.Background(), Tag{
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
	err := EditTag(context.Background(), Tag{
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
	tags, err := GetTags(context.Background(), 1, 10)
	if err != nil {
		t.Log(err)
	}
	t.Log("success", tags)
}

func TestDeleteTag(t *testing.T) {
	err := DeleteTag(context.Background(), 13)
	if err != nil {
		t.Log("err", err)
	}
	t.Log("success")
}

func TestGetTagTotal(t *testing.T) {
	tags, err := GetTagsByName(context.Background(), "test")
	if err != nil {
		t.Log("err", err)
	}
	t.Log("tags", tags)
	t.Log("total", len(tags))
}

func TestExitTagWithName(t *testing.T) {
	is := ExitTagWithName(context.Background(), "test")
	assert.Equal(t, is, true)
}

func TestGetTagById(t *testing.T) {
	tag, err := GetTagById(context.Background(), 12)
	if err != nil {
		t.Log("Err: ", err)
	}
	t.Log(tag)
}

func TestGetTagsByIds(t *testing.T) {
	names, err := GetTagsByIds(context.Background(), []uint{11, 12, 13})
	if err != nil {
		assert.Equal(t, true, gorm.IsRecordNotFoundError(err))
	}

	t.Log(names)

}
