package dto

type EformUpdateDTO struct {
	ID             uint64 `json:"id" form:"id" binding:"required"`
	NoFormulir     string `json:"nomor_formulir" form:"nomor_formulir" binding:"required"`
	Jenis_Form     string `json:"jenis_form" form:"jenis_form" binding:"required"`
	Nomor_Rekening string `json:"nomor_rekening" form:"nomor_rekening" binding:"required"`
	Nomor_Ktp      string `json:"nomor_ktp" form:"nomor_ktp" binding:"required"`
	Nama_Ibu       string `json:"nama_ibu" form:"nama_ibu" binding:"required"`
	Pekerjaan      string `json:"pekerjaan" form:"pekerjaan" binding:"required"`
	Pendidikan     string `json:"pendidikan" form:"pendidikan" binding:"required"`
	Alamat         string `json:"alamat" form:"alamat" binding:"required"`
	UserID         uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
}

type EformCreateDTO struct {
	NoFormulir     string `json:"nomor_formulir" form:"nomor_forUsermulir" binding:"required"`
	Jenis_Form     string `json:"jenis_form" form:"jenis_form" binding:"required"`
	Nomor_Rekening string `json:"nomor_rekening" form:"nomor_rekening" binding:"required"`
	Nomor_Ktp      string `json:"nomor_ktp" form:"nomor_ktp" binding:"required"`
	Nama_Ibu       string `json:"nama_ibu" form:"nama_ibu" binding:"required"`
	Pekerjaan      string `json:"pekerjaan" form:"pekerjaan" binding:"required"`
	Pendidikan     string `json:"pendidikan" form:"pendidikan" binding:"required"`
	Alamat         string `json:"alamat" form:"alamat" binding:"required"`
	UserID         uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
}
