package repositories

import "golang-api/modules/users/models/domain"

const contextName = "modules.users.repository"

type Repository interface {
	Save(user *domain.Users) (*domain.Users, error)
	FindByUsername(username string) (*domain.Users, error)
	FindById(Id string) (*domain.Users, error)
}
