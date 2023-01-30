package repository

import (
	entity "github.com/michaeldeven/microserviceuniversalbpr/entity/models"
	"gorm.io/gorm"
)

type PengaduanRepository interface {
	InsertPengaduan(b entity.Pengaduan) entity.Pengaduan
	UpdatePengaduan(b entity.Pengaduan) entity.Pengaduan
	DeletePengaduan(b entity.Pengaduan)
	AllPengaduan() []entity.Pengaduan
	FindPengaduanByID(pengaduanID uint64) entity.Pengaduan
}

type pengaduanConnection struct {
	connection *gorm.DB
}

func NewPengaduanRepository(dbConn *gorm.DB) PengaduanRepository {
	return &pengaduanConnection{
		connection: dbConn,
	}
}

func (db *pengaduanConnection) InsertPengaduan(b entity.Pengaduan) entity.Pengaduan {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *pengaduanConnection) UpdatePengaduan(b entity.Pengaduan) entity.Pengaduan {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *pengaduanConnection) DeletePengaduan(b entity.Pengaduan) {
	db.connection.Delete(&b)
}

func (db *pengaduanConnection) FindPengaduanByID(pengaduanID uint64) entity.Pengaduan {
	var pengaduan entity.Pengaduan
	db.connection.Preload("User").Find(&pengaduan, pengaduanID)
	return pengaduan
}

func (db *pengaduanConnection) AllPengaduan() []entity.Pengaduan {
	var pengaduans []entity.Pengaduan
	db.connection.Preload("User").Find(&pengaduans)
	return pengaduans
}
