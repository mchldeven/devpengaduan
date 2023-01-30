package entity

type Pengaduan struct {
	ID             uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Noaduan        string `gorm:"type:varchar(255)" json:"noaduan"`
	Nomor_Rekening string `gorm:"type:varchar(255)" json:"nomor_rekening"`
	Description    string `gorm:"type:text" json:"description"`
	UserID         uint64 `gorm:"not null" json:"-"`
	User           User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
