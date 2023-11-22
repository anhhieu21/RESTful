package service

import (
	"main/data/req"
	"main/data/res"
)

type TagsService interface {
	Create(tags req.CreateTagsRequest)
	Update(tags req.UpdateTagsRequest)
	Delete(tagsId int)
	FindById(tagsId int) res.TagsResponse
	FindAll() []res.TagsResponse
}
