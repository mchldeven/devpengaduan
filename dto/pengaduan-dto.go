package dto

type PengaduanUpdateDTO struct {
	ID             uint64 `json:"id" form:"id" binding:"required"`
	Noaduan        string `json:"noaduan" form:"noaduan" binding:"required"`
	Nomor_Rekening string `json:"nomor_rekening" form:"nomorrekening" binding:"required"`
	Description    string `json:"description" form:"description" binding:"required"`
	UserID         uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
}

type PengaduanCreateDTO struct {
	Noaduan        string `json:"noaduan" form:"noaduan" binding:"required"`
	Nomor_Rekening string `json:"nomor_rekening" form:"nomorrekening" binding:"required"`
	Description    string `json:"description" form:"description" binding:"required"`
	UserID         uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
}
