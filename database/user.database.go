package database

import "akikflix/model"

func CreateUser(user *model.User) model.User {
	DB.Create(&user)
	return *user
}

func CheckIfUserExists(phone string) (model.User,bool) {
	user := model.User{};
	result := DB.Where("phone = ?",phone).First(&user)
	if result.RowsAffected == 0 {
		return model.User{},false;
	}
	return user, true;
}