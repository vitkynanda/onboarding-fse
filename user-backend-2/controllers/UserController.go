package controllers

import (
	"fmt"

	"go-api/connection"
	"go-api/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func GetAllUsers(c *gin.Context) {
	
	users := []models.User{}

	err := connection.DB.Find(&users).Error

	// response := []models.ResponseUsers{}
	response := []models.Response{}

	for _, user := range users {

	contacts := models.Contacts{Email:user.Email, Phone: user.Phone}

	arrayHobbies := strings.Split(user.Hobbies, ",")
	
	responseData := models.Response{
		Id : user.Id, 
		Firstname : user.Firstname, 
		Lastname : user.Lastname,
		Gender : user.Gender,
		Birthdate: user.Birthdate,
		Active : user.Active,
		Hobbies : arrayHobbies,
		Contacts : contacts,
	}
		// responseData := models.ResponseUsers{
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
	
	user := models.User{}
	
	err := connection.DB.Where("id = ?", id).Find(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message" : "Internal server error"})
	}

	contacts := models.Contacts{Email:user.Email, Phone: user.Phone}

	arrayHobbies := strings.Split(user.Hobbies, ",")
	
	response := models.Response{
		Id : user.Id, 
		Firstname : user.Firstname, 
		Lastname : user.Lastname,
		Gender : user.Gender,
		Birthdate: user.Birthdate,
		Active : user.Active,
		Hobbies : arrayHobbies,
		Contacts : contacts,
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Get user by id successfully", "data" : response})
}

func CreateNewUser(c *gin.Context) {
	
	id := uuid.New()
	response := models.Response{
		Id: id,
	}

	if err := c.ShouldBindJSON(&response); err != nil {
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

	stringHobbies := strings.Join(response.Hobbies,",")
	
	user := models.User{
		Id: response.Id,
		Email: response.Contacts.Email,
		Phone: response.Contacts.Phone,
		Hobbies: stringHobbies,
		Active: response.Active,
		Gender: response.Gender,
		Firstname: response.Firstname,
		Lastname: response.Lastname,
		Birthdate: response.Birthdate,
	}

	if errDB := connection.DB.Create(&user).Error; errDB != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Internal Server Error", "errors": "Failed create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Create user  successfully",  "data" : response})
}

func UpdateUserData(c *gin.Context) {

	id := c.Param("id")
	
	response := models.Response{}
	
	if err := c.ShouldBindJSON(&response); err != nil {

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

	stringHobbies := strings.Join(response.Hobbies,",")
	
	if errDB := connection.DB.Model(&models.User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"firstName": response.Firstname,
		"lastName": response.Lastname,
		"gender": response.Gender,
		"birthdate": response.Birthdate,
		"active": response.Active,
		"email": response.Contacts.Email,
		"phone": response.Contacts.Phone,
		"hobbies": stringHobbies,
	}).Error; errDB != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Internal Server Error", "errors": "Failed update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Update user  successfully",  "data" : response})
}

func DeletUserById(c *gin.Context) {
	
	sql := "DELETE FROM users"
	id := c.Param("id")

	sql = fmt.Sprintf("%s WHERE id = '%s'", sql, id)

	if err := connection.DB.Raw(sql).Scan(models.User{}).Error; err != nil  {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Internal Server Error", "errors": "Failed delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User deleted successfully"})
}
