package controllers

import (
	"fmt"

	"go-api/connection"
	"go-api/models/dto"
	"go-api/models/entity"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func GetAllUsers(c *gin.Context) {
	
	users := []entity.User{}

	err := connection.DB.Find(&users).Error

	// response := []entity.ResponseUsers{}
	response := []dto.User{}

	for _, user := range users {

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
		// responseData := entity.ResponseUsers{
		// 	Id : user.Id,
		// 	Gender : user.Gender,
		// 	Firstname : user.Firstname,
		// 	Lastname : user.Lastname,
		// 	Active: user.Active,
		// 	Birthdate : user.Birthdate,
		// }
		response = append(response, responseData)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message" : "Internal server error"})
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Get all users successfully", "data": response})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	
	user := entity.User{}
	
	err := connection.DB.Where("id = ?", id).Find(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message" : "Internal server error"})
	}

	contact := dto.Contact{Email:user.Email, Phone: user.Phone}

	arrayHobbies := strings.Split(user.Hobbies, ",")
	
	response := dto.User{
		Id : user.Id, 
		Firstname : user.Firstname, 
		Lastname : user.Lastname,
		Gender : user.Gender,
		Birthdate: user.Birthdate,
		Active : user.Active,
		Hobbies : arrayHobbies,
		Contact : contact,
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Get user by id successfully", "data" : response})
}

func CreateNewUser(c *gin.Context) {
	
	id := uuid.New()
	request := dto.User{
		Id: id,
	}

	if err := c.ShouldBindJSON(&request); err != nil {
	errorMessages :=  []string{}

	for _, e :=  range err.(validator.ValidationErrors) {
		tag := e.ActualTag()
		field := e.Field()

		errorMessage := fmt.Sprintf("Error on Field %s, condition: %s", field, tag)
		errorMessages = append(errorMessages,  errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad request", "errors": errorMessages})
		return
	}

	stringHobbies := strings.Join(request.Hobbies,",")
	
	user := entity.User{
		Id: request.Id,
		Email: request.Contact.Email,
		Phone: request.Contact.Phone,
		Hobbies: stringHobbies,
		Active: request.Active,
		Gender: request.Gender,
		Firstname: request.Firstname,
		Lastname: request.Lastname,
		Birthdate: request.Birthdate,
	}

	if errDB := connection.DB.Create(&user).Error; errDB != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Internal Server Error", "errors": "Failed create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Create user  successfully",  "data" : request})
}

func UpdateUserData(c *gin.Context) {

	id := c.Param("id")
	
	request := dto.User{}
	
	if err := c.ShouldBindJSON(&request); err != nil {

	errorMessages :=  []string{}
	
	for _, e :=  range err.(validator.ValidationErrors) {
		tag := e.ActualTag()
		field := e.Field()

		errorMessage := fmt.Sprintf("Error on Field %s, condition: %s", field, tag)
		errorMessages = append(errorMessages,  errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad request", "errors": errorMessages})
		return
	}

	stringHobbies := strings.Join(request.Hobbies,",")
	
	if errDB := connection.DB.Model(&entity.User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"firstName": request.Firstname,
		"lastName": request.Lastname,
		"gender": request.Gender,
		"birthdate": request.Birthdate,
		"active": request.Active,
		"email": request.Contact.Email,
		"phone": request.Contact.Phone,
		"hobbies": stringHobbies,
	}).Error; errDB != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Internal Server Error", "errors": "Failed update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Update user  successfully",  "data" : request})
}

func DeletUserById(c *gin.Context) {
	
	sql := "DELETE FROM users"
	id := c.Param("id")

	sql = fmt.Sprintf("%s WHERE id = '%s'", sql, id)

	if err := connection.DB.Raw(sql).Scan(entity.User{}).Error; err != nil  {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Internal Server Error", "errors": "Failed delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User deleted successfully"})
}
