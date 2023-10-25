package repository

import (
	model "Laundary-Backend/model"
)

func Signup(user model.User) error {
	res := Database.Create(&user)
	return res.Error
}
