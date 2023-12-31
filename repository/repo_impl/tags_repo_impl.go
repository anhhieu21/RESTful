package repoimpl

import (
	"errors"
	"main/data/req"
	"main/helper"
	"main/model"
	"main/repository"

	"gorm.io/gorm"
)

type TagsRepositoryImpl struct {
	Db *gorm.DB
}

func NewTagsRepositoryImpl(Db *gorm.DB) repository.TagsRepository {
	return &TagsRepositoryImpl{Db: Db}
}

func (t *TagsRepositoryImpl) Save(tags model.Tags) {
	result := t.Db.Create(&tags)
	helper.ErrorPanic(result.Error)
}
func (t *TagsRepositoryImpl) Update(tags model.Tags) {
	var updateTag = req.UpdateTagsRequest{
		Id:   tags.Id,
		Name: tags.Name,
	}
	result := t.Db.Model(&tags).Updates(updateTag)
	helper.ErrorPanic(result.Error)
}
func (t *TagsRepositoryImpl) Delete(tagsId int) {
	var tags model.Tags
	result := t.Db.Where("id= ?", tagsId).Delete(&tags)
	helper.ErrorPanic(result.Error)
}
func (t *TagsRepositoryImpl) FindById(tagsId int) (tags model.Tags, err error) {
	var tag model.Tags
	result := t.Db.Find(&tag, tagsId)
	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("not found")
	}
}
func (t TagsRepositoryImpl) FindAll() []model.Tags {
	var tags []model.Tags
	result := t.Db.Find(&tags)
	helper.ErrorPanic(result.Error)
	return tags
}
