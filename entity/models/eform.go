package entity

type Eform struct {
	ID             uint64 `gorm:"primary_key:auto_increment" json:"id"`
	NoFormulir     string `gorm:"type:varchar(255)" json:"no_formulir"`
	Jenis_Form     string `gorm:"type:varchar(255)" json:"jenis_form"`
	Nomor_Rekening string `gorm:"type:varchar(255)" json:"nomor_rekening"`
	Nomor_Ktp      string `gorm:"type:varchar(255)" json:"no_ktp"`
	Nama_Ibu       string `gorm:"type:varchar(255)" json:"nama_ibu"`
	Pekerjaan      string `gorm:"type:varchar(255)" json:"pekerjaan"`
	Pendidikan     string `gorm:"type:varchar(255)" json:"pendidikan"`
	Alamat         string `gorm:"type:varchar(255)" json:"alamat"`
	UserID         uint64 `gorm:"not null" json:"-"`
	User           User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
