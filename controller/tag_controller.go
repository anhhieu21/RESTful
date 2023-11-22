package controller

import (
	"main/data/req"
	"main/data/res"
	"main/helper"
	"main/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TagController struct {
	tagService service.TagsService
}

func NewTagController(service service.TagsService) *TagController {
	return &TagController{
		tagService: service,
	}
}

func (c *TagController) Create(ctx *gin.Context) {
	createTagRequest := req.CreateTagsRequest{}
	err := ctx.ShouldBindJSON(&createTagRequest)
	helper.ErrorPanic(err)
	c.tagService.Create(createTagRequest)

	webResponse := res.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *TagController) Update(ctx *gin.Context) {
	updateTagRequest := req.UpdateTagsRequest{}
	err := ctx.ShouldBindJSON(&updateTagRequest)
	helper.ErrorPanic(err)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	updateTagRequest.Id = id
	c.tagService.Update(updateTagRequest)
	webResponse := res.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
func (c *TagController) Delete(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	c.tagService.Delete(id)

	webResponse := res.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
func (c *TagController) FindById(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	tagResponse := c.tagService.FindById(id)
	webResponse := res.Response{
		Code:   200,
		Status: "Ok",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
func (c *TagController) FindAll(ctx *gin.Context) {
	tagResponse := c.tagService.FindAll()

	webResponse := res.Response{
		Code:   200,
		Status: "Ok",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
