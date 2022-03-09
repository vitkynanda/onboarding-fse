package user_repository

import (
	"go-api/models/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (repo *userRepository) GetAllUsers() ([]entity.User, error) {
	users := []entity.User{}

	err := repo.mysqlConnection.Find(&users).Error

	if err != nil {
		return nil, err
	}

	if  len(users) <= 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return users, nil
}

func (repo *userRepository) GetUserById(id string) (*entity.User, error) {
	user := entity.User{}
	
	err := repo.mysqlConnection.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) CreateNewUser(user entity.User) (*entity.User, error){
	id := uuid.New()
	user.Id = id

	if err := repo.mysqlConnection.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) UpdateUserData(user entity.User, id string ) (*entity.User, error){
	
	if err := repo.mysqlConnection.Model(&user).Where("id = ?", id).Updates(map[string]interface{}{
		"firstName": user.Firstname,
		"lastName": user.Lastname,
		"gender": user.Gender,
		"birthdate": user.Birthdate,
		"active": user.Active,
		"email": user.Email,
		"phone": user.Phone,
		"hobbies": user.Hobbies,
	}).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) DeleteUserById( id string) error {
	if err := repo.mysqlConnection.Delete(&entity.User{}, id).Error; err != nil  {
		return err
	}
	return nil
}