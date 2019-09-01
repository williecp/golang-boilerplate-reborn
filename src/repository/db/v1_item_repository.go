package repositories

import (
	"../database"
	"../models"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type V1ItemRepository struct {
	DB gorm.DB
}

func V1ItemRepositoryHandler() V1ItemRepository {
	repository := V1ItemRepository{DB: *database.GetConnection()}
	return repository
}

func (repository *V1ItemRepository) GetByID(id uint) (models.Item, error) {

	itemModel := models.Item{}
	query := repository.DB.Table("rl_items")
	query = query.Where("id = ?", id)
	query = query.First(&itemModel)

	return itemModel, query.Error
}

func (repository *V1ItemRepository) Create(itemData interface{}) (models.Item, error) {
	itemModel := models.Item{}
	copier.Copy(&itemModel, &itemData)

	query := repository.DB.Begin()
	query = query.Create(&itemModel)

	if err := query.Error; err != nil {
		query.Rollback()
		return itemModel, err
	}

	return itemModel, query.Commit().Error
}

func (repository *V1ItemRepository) UpdateByID(id int, itemData interface{}) (models.Item, error) {

	itemModel := models.Item{}
	copier.Copy(&itemModel, &itemData)

	query := repository.DB.Model(&itemModel)
	query = query.Where("id = ?", id)
	query = query.Update("name", "BLICKLE BS-PS 310K heavy duty wheel with pneumatic tyre Fixed castors")
	query = query.Scan(&itemModel)
	return itemModel, query.Error

}

func (repository *V1ItemRepository) GetByName(name string) (models.Item, error) {

	itemModel := models.Item{}

	query := repository.DB.Table("rl_items")
	query = query.Where("name LIKE ?", "%"+name+"%")
	query = query.Find(&itemModel)

	if query.RecordNotFound() {
		return itemModel, nil
	}

	return itemModel, query.Error
}
