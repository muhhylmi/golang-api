package usecases

import (
	"context"
	"golang-api/modules/users/models/domain"
	"golang-api/modules/users/models/web"
	"golang-api/utils/jwt"
	"golang-api/utils/wrapper"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (usecase *UsecaseImpl) CreateUser(ctx context.Context, payload *web.RequestCreateUser) wrapper.Result {
	log := usecase.Logger.LogWithContext(contextName, "CreateBook")
	var result wrapper.Result

	_, err := usecase.Repository.FindByUsername(payload.Username)
	if err == nil {
		error := wrapper.NewBadRequest("User Already Exists")
		result.Error = error
		return result
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err)
		error := wrapper.NewBadRequest("Error Hashing Password")
		result.Error = error
		return result
	}
	userData := domain.Users{
		Id:       uuid.New().String(),
		Username: payload.Username,
		Password: string(hashedPassword),
		Gender:   payload.Gender,
	}
	user, err := usecase.Repository.Save(&userData)
	if err != nil {
		error := wrapper.NewBadRequest("Cannot Create User")
		result.Error = error
		return result
	}
	result.Data = user
	return result
}

func (usecase *UsecaseImpl) LoginUser(ctx context.Context, payload *web.RequestLogin) wrapper.Result {
	log := usecase.Logger.LogWithContext(contextName, "CreateBook")
	var result wrapper.Result

	checkUser, _ := usecase.Repository.FindByUsername(payload.Username)
	if checkUser == nil {
		error := wrapper.NewNotFound("User Not Found")
		result.Error = error
		return result
	}

	err := bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(payload.Password))
	if err != nil {
		log.Error(err)
		error := wrapper.NewBadRequest("Password is Invalid")
		result.Error = error
		return result
	}
	userData := domain.Users{
		Id:       checkUser.Id,
		Username: checkUser.Username,
		Gender:   checkUser.Gender,
	}
	token, err := jwt.CreateToken(&userData)
	if err != nil {
		log.Error(err)
		error := wrapper.NewBadRequest(err.Error())
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
