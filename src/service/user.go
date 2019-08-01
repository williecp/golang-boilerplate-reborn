package services

import (
	"github.com/jinzhu/copier"
	"github.com/lukmanralali/rl-ms-boilerplate-go/objects"
	"github.com/lukmanralali/rl-ms-boilerplate-go/repositories"
)

type UserService struct {
	// userRepository repositories.V1UserRepositoryInterface
}

func User() UserService {
	return V1UserService{repositories.V1UserRepositoryHandler()}
}

type User interface {
	GetById(id int)
	GetAllUser(page int,count int)
	UpdateById(id int)
}

func (service *V1UserService) GetById(id int){
	
}

func (service *V1UserService) GetAllUser(page int,count int){

}

func (service *V1UserService) UpdateById(id int){

}