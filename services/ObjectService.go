package services

import (
	"strconv"

	"github.com/nezor-dev/go-structure/model"
	"github.com/nezor-dev/go-structure/repositories"
)

type ObjectService struct {
	objectRepository *repositories.ObjectRepository
}

func InitObjectService(objectRepository *repositories.ObjectRepository) *ObjectService {
	return &ObjectService{
		objectRepository: objectRepository,
	}
}

func (s *ObjectService) CreateObject(inputObject *model.InputObject) (*model.Object, error) {
	//data validation etc etc

	object, err := s.objectRepository.CreateObject(inputObject.ToObject())

	return object, err

}

func (s *ObjectService) GetExistingObject(objectId string) (*model.Object, error) {
	//data validation etc etc

	oId, err := strconv.ParseUint(objectId, 10, 32)

	if err != nil {
		return nil, err
	}

	object := s.objectRepository.GetExistingObject(uint(oId))

	return object, nil
}
