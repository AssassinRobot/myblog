package repo

import (
	"blog/pkg/data/database"
	"blog/pkg/data/models"
)

func GetUser(id string) (models.User, error) {
	var u models.User
	err := database.DB.Find(&u, id).Error
	return u, err
}

func CreateUser(fristname, lastname, username, password, email, phonenumber string) (models.User, error) {
	var u models.User
	u.Fristname = fristname
	u.Lastname = lastname
	u.Username = username
	u.Password = password
	u.Email = email
	u.PhoneNumber = phonenumber

	err := database.DB.Create(&u).Error

	return u, err
}
func UpdateUserById(id, fristname, lastname, username, password, email, phonenumber string) error {
	u, err := GetUser(id)

	if err != nil {
		return err
	}
	u.Fristname = fristname
	u.Lastname = lastname
	u.Username = username
	u.Password = password
	u.Email = email
	u.PhoneNumber = phonenumber

	err = database.DB.Save(&u).Error

	return err
}
func DeleteUser(id string) error {
	var u models.User
	err := database.DB.Delete(&u, id).Error

	return err
}