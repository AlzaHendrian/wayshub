package usersdto

type CreateUserRequest struct {
	Chanelname  string `json:"chanelname" form:"chanelname" validate:"required"`
	Email       string `json:"email" form:"email" validate:"required"`
	Password    string `json:"password" form:"password" validate:"required"`
	Fhoto       string `json:"fhoto" form:"fhoto" validate:"required"`
	Thumbnail   string `json:"thumbnail" form:"thumbnail" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Subcribe    bool   `json:"subcribe" form:"password" `
}

type UpdateUserRequest struct {
	Chanelname  string `json:"chanelname" form:"chanelname"`
	Email       string `json:"email" form:"email" `
	Password    string `json:"password" form:"password" `
	Fhoto       string `json:"fhoto" form:"fhoto"`
	Thumbnail   string `json:"thumbnail" form:"thumbnail"`
	Description string `json:"description" form:"description"`
}
