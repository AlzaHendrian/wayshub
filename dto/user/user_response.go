package usersdto

type UserResponse struct {
	ID          int    `json:"id"`
	Chanelname  string `json:"chanelname" form:"chanelname" validate:"required"`
	Email       string `json:"email" form:"email" validate:"required"`
	Password    string `json:"password" form:"password" validate:"required"`
	Fhoto       string `json:"fhoto" form:"fhoto" validate:"required"`
	Thumbnail   string `json:"thumbnail" form:"thumbnail" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Subcribe    bool   `json:"subcribe" form:"subcribe"`
}

type ProfileResponse struct {
	ID          int    `json:"id"`
	Chanelname  string `json:"chanelname"`
	Fhoto       string `json:"fhoto"`
	Thumbnail   string `json:"thumbnail"`
	Description string `json:"description"`
}
