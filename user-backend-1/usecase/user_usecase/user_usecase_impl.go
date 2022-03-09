package user_usecase

import (
	"errors"
	"go-api/helpers"
	"go-api/models/dto"
	"go-api/models/entity"
	"strings"

	"gorm.io/gorm"
)

func (user *userUsecase) GetAllUsers() dto.Response {
	userlist, err := user.userRepo.GetAllUsers()

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", 500)
	}
	return helpers.ResponseSuccess("Get all users successfully", 200, userlist)
}

func (user *userUsecase) GetUserById(id string) dto.Response {
	userData, err := user.userRepo.GetUserById(id)
	
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", 500)
	}
	return helpers.ResponseSuccess("Get user successfully", 200, userData)
}

func (user *userUsecase) CreateNewUser(newUser dto.User) dto.Response {
	
	stringHobbies := strings.Join(newUser.Hobbies, ",")

	userInsert := entity.User{
		Id: newUser.Id,
		Email: newUser.Contact.Email,
		Phone: newUser.Contact.Phone,
		Gender: newUser.Gender,
		Active: newUser.Active,
		Firstname: newUser.Firstname,
		Lastname: newUser.Lastname,
		Birthdate: newUser.Birthdate,
		Hobbies: stringHobbies,
	}
		
	userData, err := user.userRepo.CreateNewUser(userInsert)

	
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", 500)
	}

	newUser.Id = userData.Id
	return helpers.ResponseSuccess("Create new user successfully", 200, newUser)
}

func (user *userUsecase) UpdateUserData(userUpdate dto.User, id string) dto.Response {
	
	stringHobbies := strings.Join(userUpdate.Hobbies, ",")

	userInsert := entity.User{
		Id: userUpdate.Id,
		Email: userUpdate.Contact.Email,
		Phone: userUpdate.Contact.Phone,
		Gender: userUpdate.Gender,
		Active: userUpdate.Active,
		Firstname: userUpdate.Firstname,
		Lastname: userUpdate.Lastname,
		Birthdate: userUpdate.Birthdate,
		Hobbies: stringHobbies,
	}
		
	userData, err := user.userRepo.UpdateUserData(userInsert, id)
	 
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", 500)
	}

	userUpdate.Id = userData.Id

	return helpers.ResponseSuccess("Create new user successfully", 200, userUpdate)
}

func (user *userUsecase) DeleteUserById(id string) dto.Response {
	
 err := user.userRepo.DeleteUserById( id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", 500)
	}
	return helpers.ResponseSuccess("Create new user successfully", 200, nil)
}
