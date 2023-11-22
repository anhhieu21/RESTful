package serviceimpl

import (
	"main/data/req"
	"main/data/res"
	"main/helper"
	"main/model"
	"main/repository"
	"main/service"

	"github.com/go-playground/validator/v10"
)

type TagsServiceImpl struct {
	TagRepository repository.TagsRepository
	Validate      *validator.Validate
}

func NewTagServiceImpl(tagRepository repository.TagsRepository,
	validate *validator.Validate) service.TagsService {
	return &TagsServiceImpl{
		TagRepository: tagRepository,
		Validate:      validate,
	}
}

func (t *TagsServiceImpl) Create(tag req.CreateTagsRequest) {
	err := t.Validate.Struct(tag)
	helper.ErrorPanic(err)
	tagModel := model.Tags{
		Name: tag.Name,
	}
	t.TagRepository.Save(tagModel)
}
func (t *TagsServiceImpl) Update(tag req.UpdateTagsRequest) {
	tagData, err := t.TagRepository.FindById(tag.Id)
	helper.ErrorPanic(err)
	tagData.Name = tag.Name
	t.TagRepository.Update(tagData)

}
func (t *TagsServiceImpl) Delete(tagId int) {
	t.TagRepository.Delete(tagId)
}
func (t *TagsServiceImpl) FindById(tagId int) res.TagsResponse {
	tagData, err := t.TagRepository.FindById(tagId)
	helper.ErrorPanic(err)

	tagResponse := res.TagsResponse{
		Id:   tagData.Id,
		Name: tagData.Name,
	}
	return tagResponse
}
func (t *TagsServiceImpl) FindAll() []res.TagsResponse {
	result := t.TagRepository.FindAll()

	var tags []res.TagsResponse
	for _, value := range result {
		tag := res.TagsResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		tags = append(tags, tag)
	}
	return tags
}
