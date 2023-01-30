package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/michaeldeven/microserviceuniversalbpr/dto"
	entity "github.com/michaeldeven/microserviceuniversalbpr/entity/models"
	"github.com/michaeldeven/microserviceuniversalbpr/repository"
)

type PengaduanService interface {
	Insert(b dto.PengaduanCreateDTO) entity.Pengaduan
	Update(b dto.PengaduanUpdateDTO) entity.Pengaduan
	Delete(b entity.Pengaduan)
	All() []entity.Pengaduan
	FindByID(pengaduanID uint64) entity.Pengaduan
	IsAllowedToEdit(userID string, pengaduanID uint64) bool
}

type pengaduanService struct {
	pengaduanRepository repository.PengaduanRepository
}

//NewPengaduanService .....
func NewPengaduanService(pengaduanRepo repository.PengaduanRepository) PengaduanService {
	return &pengaduanService{
		pengaduanRepository: pengaduanRepo,
	}
}

func (service *pengaduanService) Insert(b dto.PengaduanCreateDTO) entity.Pengaduan {
	pengaduan := entity.Pengaduan{}
	err := smapping.FillStruct(&pengaduan, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.pengaduanRepository.InsertPengaduan(pengaduan)
	return res
}

func (service *pengaduanService) Update(b dto.PengaduanUpdateDTO) entity.Pengaduan {
	pengaduan := entity.Pengaduan{}
	err := smapping.FillStruct(&pengaduan, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.pengaduanRepository.UpdatePengaduan(pengaduan)
	return res
}

func (service *pengaduanService) Delete(b entity.Pengaduan) {
	service.pengaduanRepository.DeletePengaduan(b)
}

func (service *pengaduanService) All() []entity.Pengaduan {
	return service.pengaduanRepository.AllPengaduan()
}

func (service *pengaduanService) FindByID(pengaduanID uint64) entity.Pengaduan {
	return service.pengaduanRepository.FindPengaduanByID(pengaduanID)
}

func (service *pengaduanService) IsAllowedToEdit(userID string, pengaduanID uint64) bool {
	b := service.pengaduanRepository.FindPengaduanByID(pengaduanID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
