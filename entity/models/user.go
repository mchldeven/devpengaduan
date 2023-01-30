package entity

type User struct {
	ID          uint64       `gorm:"primary_key:auto_increment" json:":id"`
	NamaLengkap string       `gorm:"type:varchar(255)" json:"namalengkap"`
	Email       string       `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password    string       `gorm:"->;<-;not null" json:"-"`
	Token       string       `gorm:"-" json:"token,omitempty"`
	Pengaduans  *[]Pengaduan `json:"pengaduans,omitempty"`
	Eforms      *[]Eform     `json:"eforms,omitempty"`
}
