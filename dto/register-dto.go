package dto

type RegisterDTO struct {
	NamaLengkap string `json:"namalengkap" form:"namalengkap	" binding:"required"`
	Email       string `json:"email" form:"email" binding:"required,email" `
	Password    string `json:"password" form:"password" binding:"required"`
}
