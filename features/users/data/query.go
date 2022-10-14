package data

import (
	"be12/deploy/features/users"

	"gorm.io/gorm"
)

type userData struct {
	DB *gorm.DB
}

func New(conn *gorm.DB) users.DataInterface {
	return &userData{
		DB: conn,
	}
}

func (repo *userData) DelData(param int) (int, error) {

	var data User
	txDelId := repo.DB.Delete(&data, param)
	if txDelId.Error != nil {
		return -1, txDelId.Error
	}

	return int(txDelId.RowsAffected), nil

}

func (repo *userData) UpdateData(param int, dataUpdate users.UserCore) (int, error) {

	var data User
	data.Name = dataUpdate.Name
	data.Hp = dataUpdate.Hp
	data.Email = dataUpdate.Email
	data.Password = dataUpdate.Password

	var user User
	user.ID = dataUpdate.ID
	txUpdateId := repo.DB.Model(&user).Updates(data)
	if txUpdateId.Error != nil {
		return -1, txUpdateId.Error
	}

	var err error

	return int(txUpdateId.RowsAffected), err

}

func (repo *userData) CreateData(data users.UserCore) (int, error) {

	dataModel := ModelInsert(data)
	tx := repo.DB.Create(&dataModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil

}
func (repo *userData) SelectAll() ([]users.UserCore, error) {

	var dataAll []User
	tx := repo.DB.Find(&dataAll)
	if tx.Error != nil {
		return nil, tx.Error
	}

	dataCore := toUserCoreList(dataAll)
	return dataCore, nil

}

func (repo *userData) user() []User {

	var dataBookUser []User
	tx := repo.DB.Find(&dataBookUser)
	if tx.Error != nil {
		return nil
	}

	return dataBookUser

}

func (repo *userData) SelectById(param int) (users.UserCore, error) {

	var data User
	tx := repo.DB.First(&data, param)
	if tx.Error != nil {
		return users.UserCore{}, tx.Error
	}

	userList := repo.user()

	userId := data.toUserCore(userList)
	return userId, nil

}
