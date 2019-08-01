package services

import (
	"github.com/jinzhu/copier"
	"github.com/lukmanralali/rl-ms-boilerplate-go/objects"
	"github.com/lukmanralali/rl-ms-boilerplate-go/repositories"
)

type V1UserService struct {
	userRepository repositories.V1UserRepositoryInterface
}

func V1UserServiceHandler() V1UserService {
	return V1UserService{repositories.V1UserRepositoryHandler()}
}

type Use interface {
	GetById(id int) (objects.V1UserObjectResponse, error)
	UpdateById(id int, requestObject objects.V1UserObjectRequest) (objects.V1UserObjectResponse, error)
}

func (service *V1UserService) GetById(id int) (objects.V1UserObjectResponse, error) {
	user, err := service.userRepository.GetById(id)
	if nil != err {
		return objects.V1UserObjectResponse{}, err
	}
	result := objects.V1UserObjectResponse{}
	copier.Copy(&result, &user)
	return result, nil
}

func (service *V1UserService) UpdateById(id int, requestObject objects.V1UserObjectRequest) (objects.V1UserObjectResponse, error) {

	user, err := service.userRepository.UpdateById(id, requestObject)
	if nil != err {
		return objects.V1UserObjectResponse{}, err
	}

	result := objects.V1UserObjectResponse{}
	copier.Copy(&result, &user)

	return result, nil

}
