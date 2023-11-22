package req


type CreateTagsRequest struct {
	Name string `validate:"required,min=1,max=200" json:"name"`
}

type UpdateTagsRequest struct {
	Id int `validate:"required"`
    Name string `validate:"required,min=1,max=200" json:"name"`
}