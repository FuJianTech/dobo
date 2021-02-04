package service

import (
	"dobo/database"
	"dobo/database/model"
	"fmt"

)

// GetUserByEmail ...
func GetUserByEmail(email string) (*model.Auth, error) {
	db := database.GetDB()

	var auth model.Auth

	if err := db.Where("email = ? ", email).Find(&auth).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &auth, nil
}
