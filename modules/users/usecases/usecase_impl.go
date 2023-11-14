package usecases

import (
	"context"
	"golang-api/modules/users/models/domain"
	"golang-api/modules/users/models/web"
	"golang-api/modules/users/repositories"
	"golang-api/utils"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UsecaseImpl struct {
	logger     *logrus.Logger
	repository repositories.Repository
}

func NewUsecaseImpl(logger *logrus.Logger, repository repositories.Repository) Usecases {
	return &UsecaseImpl{
		logger:     logger,
		repository: repository,
	}
}

func (usecase *UsecaseImpl) CreateUser(ctx context.Context, payload *web.RequestCreateUser) utils.Result {
	log := utils.LogWithContext(usecase.logger, contextName, "CreateBook")
	var result utils.Result

	checkUser, _ := usecase.repository.FindByUsername(payload.Username)
	if checkUser != nil {
		error := utils.NewBadRequest("User Already Exists")
		result.Error = error
		return result
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err)
		error := utils.NewBadRequest("Error Hashing Password")
		result.Error = error
		return result
	}
	userData := domain.Users{
		Id:       uuid.New().String(),
		Username: payload.Username,
		Password: string(hashedPassword),
		Gender:   payload.Gender,
	}
	user, err := usecase.repository.Save(&userData)
	if err != nil {
		error := utils.NewBadRequest("Cannot Create User")
		result.Error = error
		return result
	}
	result.Data = user
	return result
}
