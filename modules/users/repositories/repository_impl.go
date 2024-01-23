package repositories

import (
	"golang-api/modules/users/models/domain"
)

func (repository *RepositoryImpl) Save(user *domain.Users) (*domain.Users, error) {
	log := repository.Logger.LogWithContext(contextName, "Save")
	result := repository.DB.Gorm.Create(&user)
	if result.Error != nil {
		log.Error(result.Error)
	}
	return user, result.Error
}

func (repository *RepositoryImpl) FindByUsername(username string) (*domain.Users, error) {
	var user *domain.Users
	result := repository.DB.Gorm.Where(&domain.Users{Username: username}).First(&user)
	return user, result.Error
}

func (repository *RepositoryImpl) FindById(id string) (*domain.Users, error) {
	var user *domain.Users
	result := repository.DB.Gorm.Where(&domain.Users{Id: id}).First(&user)
	return user, result.Error
}
