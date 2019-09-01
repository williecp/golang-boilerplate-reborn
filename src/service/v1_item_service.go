package services

import (
	"../objects"
	"../repositories"
	"github.com/jinzhu/copier"
)

type V1ItemService struct {
	request        objects.V1ItemObjectResponse
	itemRepository repositories.V1ItemRepository
}

func V1ItemServiceHandler() V1ItemService {
	service := V1ItemService{
		itemRepository: repositories.V1ItemRepositoryHandler(),
	}
	return service
}

func (service *V1ItemService) GetByID(id uint) (objects.V1ItemObjectResponse, error) {

	item, err := service.itemRepository.GetByID(id)
	if nil != err {
		return objects.V1ItemObjectResponse{}, err
	}

	result := objects.V1ItemObjectResponse{}
	copier.Copy(&result, &item)

	return result, nil
}

func (service *V1ItemService) Create(requestObject objects.V1ItemObjectRequest) (objects.V1ItemObjectResponse, error) {

	item, err := service.itemRepository.Create(requestObject)
	if nil != err {
		return objects.V1ItemObjectResponse{}, err
	}

	result := objects.V1ItemObjectResponse{}
	copier.Copy(&result, &item)

	return result, nil
}

func (service *V1ItemService) UpdateByID(id int, requestObject objects.V1ItemObjectRequest) (objects.V1ItemObjectResponse, error) {

	item, err := service.itemRepository.UpdateByID(id, requestObject)
	if nil != err {
		return objects.V1ItemObjectResponse{}, err
	}

	result := objects.V1ItemObjectResponse{}
	copier.Copy(&result, &item)

	return result, nil
}

func (service *V1ItemService) GetByName(name string) (objects.V1ItemObjectResponse, error) {

	item, err := service.itemRepository.GetByName(name)
	if nil != err {
		return objects.V1ItemObjectResponse{}, err
	}

	result := objects.V1ItemObjectResponse{}
	copier.Copy(&result, &item)

	return result, nil
}
