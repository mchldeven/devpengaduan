package service

import (
	"fmt"
	"github.com/mashingan/smapping"
	"github.com/michaeldeven/microserviceuniversalbpr/dto"
	entity "github.com/michaeldeven/microserviceuniversalbpr/entity/models"
	"github.com/michaeldeven/microserviceuniversalbpr/repository"
	"log"
)

type EformService interface {
	Insert(b dto.EformCreateDTO) entity.Eform
	Update(b dto.EformUpdateDTO) entity.Eform
	Delete(b entity.Eform)
	All() []entity.Eform
	FindByID(eformID uint64) entity.Eform
	IsAllowedToEdit(userID string, eformID uint64) bool
}

type eformService struct {
	eformRepository repository.EformRepository
}

// NewEformService .....
func NewEformService(eformRepo repository.EformRepository) EformService {
	return &eformService{
		eformRepository: eformRepo,
	}
}

func (service *eformService) Insert(b dto.EformCreateDTO) entity.Eform {
	eform := entity.Eform{}
	err := smapping.FillStruct(&eform, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.eformRepository.InsertEform(eform)
	return res
}

func (service *eformService) Update(b dto.EformUpdateDTO) entity.Eform {
	eform := entity.Eform{}
	err := smapping.FillStruct(&eform, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.eformRepository.UpdateEform(eform)
	return res
}

func (service *eformService) Delete(b entity.Eform) {
	service.eformRepository.DeleteEform(b)
}

func (service *eformService) All() []entity.Eform {
	return service.eformRepository.AllEform()
}

func (service *eformService) FindByID(eformID uint64) entity.Eform {
	return service.eformRepository.FindEformByID(eformID)
}

func (service *eformService) IsAllowedToEdit(userID string, eformID uint64) bool {
	b := service.eformRepository.FindEformByID(eformID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
