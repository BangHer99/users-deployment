package services

import (
	"be12/deploy/features/users"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	dataUser users.DataInterface
}

func New(data users.DataInterface) users.ServiceInterface {
	return &userService{
		dataUser: data,
	}

}

func (service *userService) DeleteData(param int) (int, error) {

	// if param != token {
	// 	return -1, errors.New("not have access")
	// }

	_, err := service.dataUser.DelData(param)
	if err != nil {
		return -1, err
	}

	return 1, nil

}

func (service *userService) PutData(param int, data users.UserCore) (int, error) {

	row, err := service.dataUser.UpdateData(param, data)
	if err != nil || row == 0 {
		return -1, err
	}

	return 1, nil

}
func (service *userService) GetById(param int) (users.UserCore, error) {

	dataId, err := service.dataUser.SelectById(param)
	if err != nil {
		return users.UserCore{}, err
	}

	return dataId, nil

}

func (service *userService) PostData(data users.UserCore) (int, error) {

	if data.Email != "" && data.Name != "" && data.Password != "" && data.Hp != "" {
		passByte := []byte(data.Password)
		hashPass, _ := bcrypt.GenerateFromPassword(passByte, bcrypt.DefaultCost)
		data.Password = string(hashPass)
		add, err := service.dataUser.CreateData(data)
		if err != nil || add == 0 {
			return -1, err
		} else {
			return 1, nil
		}
	} else {
		return -1, errors.New("all input data must be filled")
	}

}

func (service *userService) GetAll() ([]users.UserCore, error) {

	dataAll, err := service.dataUser.SelectAll()
	if err != nil {
		return nil, errors.New("gagal melihat data")
	} else if len(dataAll) == 0 {
		return nil, errors.New("data tidak ada")
	} else {
		return dataAll, nil
	}

}
