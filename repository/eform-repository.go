package repository

import (
	entity "github.com/michaeldeven/microserviceuniversalbpr/entity/models"
	"gorm.io/gorm"
)

type EformRepository interface {
	InsertEform(b entity.Eform) entity.Eform
	UpdateEform(b entity.Eform) entity.Eform
	DeleteEform(b entity.Eform)
	AllEform() []entity.Eform
	FindEformByID(pengaduanID uint64) entity.Eform
}

type eformConnection struct {
	connection *gorm.DB
}

func NewEformRepository(dbConn *gorm.DB) EformRepository {
	return &eformConnection{
		connection: dbConn,
	}
}

func (db *eformConnection) InsertEform(b entity.Eform) entity.Eform {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *eformConnection) UpdateEform(b entity.Eform) entity.Eform {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *eformConnection) DeleteEform(b entity.Eform) {
	db.connection.Delete(&b)
}

func (db *eformConnection) FindEformByID(eformID uint64) entity.Eform {
	var eform entity.Eform
	db.connection.Preload("User").Find(&eform, eformID)
	return eform
}

func (db *eformConnection) AllEform() []entity.Eform {
	var eforms []entity.Eform
	db.connection.Preload("User").Find(&eforms)
	return eforms
}
