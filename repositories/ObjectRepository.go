package repositories

import (
	"github.com/nezor-dev/go-structure/model"
	"gorm.io/gorm"
)

type ObjectRepository struct {
	db *gorm.DB
}

func NewObjectRepository(db *gorm.DB) *ObjectRepository {
	return &ObjectRepository{db: db}
}

func (repo *ObjectRepository) CreateObject(object *model.Object) (*model.Object, error) {

	result := repo.db.Create(&object)

	if result.Error != nil {
		return &model.Object{}, result.Error
	}

	return object, nil
}

func (repo *ObjectRepository) GetExistingObject(id uint) *model.Object {
	var object *model.Object
	repo.db.Where("id = ?", id).First(object)
	return object
}
