package authdto

type AuthRequest struct {
	Email       string `json:"email" gorm:"type: varchar(255)" validate:"required"`
	Password    string `json:"password" gorm:"type: varchar(255)" validate:"required"`
	Chanelname  string `json:"chanelname" gorm:"type: varchar(255)" validate:"required"`
	Fhoto       string `json:"fhoto" gorm:"type: varchar(255)"`
	Thumbnail   string `json:"thumbnail" gorm:"type: varchar(255)"`
	Description string `json:"description" gorm:"type: varchar(255)"`
	Subcribe    bool   `json:"subcribe" gorm:"type: text"`
	UserId      int    `json:"-" gorm:"type: int"`
}

type LoginRequest struct {
	Email    string `json:"email" gorm:"type: varchar(255)" validate:"required"`
	Password string `json:"password" gorm:"type: varchar(255)" validate:"required"`
}
