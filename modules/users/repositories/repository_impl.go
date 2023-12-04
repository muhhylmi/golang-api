package repositories

import (
	"golang-api/modules/users/models/domain"
	"golang-api/utils"
)

func (repository *RepositoryImpl) Save(user *domain.Users) (*domain.Users, error) {
	log := utils.LogWithContext(repository.logger, contextName, "Save")
	result := repository.db.Create(&user)
	if result.Error != nil {
		log.Error(result.Error)
	}
	return user, result.Error
}

func (repository *RepositoryImpl) FindByUsername(username string) (*domain.Users, error) {
	var user *domain.Users
	result := repository.db.Where(&domain.Users{Username: username}).First(&user)
	return user, result.Error
}

func (repository *RepositoryImpl) FindById(id string) (*domain.Users, error) {
	var user *domain.Users
	result := repository.db.Where(&domain.Users{Id: id}).First(&user)
	return user, result.Error
}
