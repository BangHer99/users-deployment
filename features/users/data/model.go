package data

import (
	"be12/deploy/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Hp       string
	Email    string `gorm:"primary key"`
	Password string
}

func ModelInsert(data users.UserCore) User {

	userData := User{
		Name:     data.Name,
		Hp:       data.Hp,
		Email:    data.Email,
		Password: data.Password,
	}

	return userData

}

func (data *User) toUserCore(u []User) users.UserCore {

	dataUser := users.UserCore{
		ID:        data.ID,
		Name:      data.Name,
		Hp:        data.Hp,
		Email:     data.Email,
		Password:  data.Password,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return dataUser
}

func toUserCoreList(data []User) []users.UserCore {
	var dataCore []users.UserCore
	for i := 0; i < len(data); i++ {
		dataCore = append(dataCore, data[i].toUserCore(data))
	}
	return dataCore
}
