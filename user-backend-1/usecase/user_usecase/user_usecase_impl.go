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
	response := []dto.User{}

	for _, user := range userlist {

	contact := dto.Contact{Email:user.Email, Phone: user.Phone}

	arrayHobbies := strings.Split(user.Hobbies, ",")
	
	responseData := dto.User{
		Id : user.Id, 
		Firstname : user.Firstname, 
		Lastname : user.Lastname,
		Gender : user.Gender,
		Birthdate: user.Birthdate,
		Active : user.Active,
		Hobbies : arrayHobbies,
		Contact : contact,
	}
		
		response = append(response, responseData)
	}

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", 500)
	}
	return helpers.ResponseSuccess("Get all users successfully", 200, response)
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
	
	 if err != nil {
		return helpers.ResponseError("Internal server error", 500)
	}

	newUser.Id = userData.Id
	return helpers.ResponseSuccess("New user created successfully", 200, newUser)
}

func (user *userUsecase) UpdateUserData(userUpdate dto.User, id string) dto.Response {
	
	stringHobbies := strings.Join(userUpdate.Hobbies, ",")

	userInsert := entity.User{
		Email: userUpdate.Contact.Email,
		Phone: userUpdate.Contact.Phone,
		Gender: userUpdate.Gender,
		Active: userUpdate.Active,
		Firstname: userUpdate.Firstname,
		Lastname: userUpdate.Lastname,
		Birthdate: userUpdate.Birthdate,
		Hobbies: stringHobbies,
	}
		
	_, err := user.userRepo.UpdateUserData(userInsert, id)
	 
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", 500)
	}



	return helpers.ResponseSuccess("User data updated successfully", 200, userUpdate)
}

func (user *userUsecase) DeleteUserById(id string) dto.Response {
	
 err := user.userRepo.DeleteUserById(id)
 
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", 500)
	}
	return helpers.ResponseSuccess("User deleted successfully", 200, nil)
}
