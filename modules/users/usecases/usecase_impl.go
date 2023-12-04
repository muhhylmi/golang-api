package usecases

import (
	"context"
	"golang-api/modules/users/models/domain"
	"golang-api/modules/users/models/web"
	"golang-api/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (usecase *UsecaseImpl) CreateUser(ctx context.Context, payload *web.RequestCreateUser) utils.Result {
	log := utils.LogWithContext(usecase.logger, contextName, "CreateBook")
	var result utils.Result

	_, err := usecase.repository.FindByUsername(payload.Username)
	if err == nil {
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

func (usecase *UsecaseImpl) LoginUser(ctx context.Context, payload *web.RequestLogin) utils.Result {
	log := utils.LogWithContext(usecase.logger, contextName, "CreateBook")
	var result utils.Result

	checkUser, _ := usecase.repository.FindByUsername(payload.Username)
	if checkUser == nil {
		error := utils.NewNotFound("User Not Found")
		result.Error = error
		return result
	}

	err := bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(payload.Password))
	if err != nil {
		log.Error(err)
		error := utils.NewBadRequest("Password is Invalid")
		result.Error = error
		return result
	}
	userData := domain.Users{
		Id:       checkUser.Id,
		Username: checkUser.Username,
		Gender:   checkUser.Gender,
	}
	token, err := utils.CreateToken(&userData)
	if err != nil {
		log.Error(err)
		error := utils.NewBadRequest(err.Error())
		result.Error = error
		return result
	}

	result.Data = web.ResoponseLogin{
		Id:       checkUser.Id,
		Username: checkUser.Username,
		Gender:   checkUser.Gender,
		Token:    token,
	}
	return result
}
