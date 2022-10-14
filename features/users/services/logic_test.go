package services

import (
	"be12/deploy/features/users"
	"be12/deploy/mocks"

	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostData(t *testing.T) {

	userMock := new(mocks.UserData)
	input := users.UserCore{Name: "Hery", Hp: "087123", Email: "Hery@gmail.com", Password: "asdf"}

	t.Run("create success", func(t *testing.T) {

		userMock.On("CreateData", mock.Anything).Return(1, nil).Once()

		useCase := New(userMock)
		res, _ := useCase.PostData(input)
		assert.Equal(t, 1, res)
		userMock.AssertExpectations(t)
	})

	t.Run("create failed", func(t *testing.T) {

		userMock.On("CreateData", mock.Anything).Return(-1, errors.New("error")).Once()

		useCase := New(userMock)
		res, err := useCase.PostData(input)
		assert.Equal(t, -1, res)
		assert.Error(t, err)
		userMock.AssertExpectations(t)

	})

	t.Run("create failed", func(t *testing.T) {

		input.Name = ""
		input.Hp = ""
		input.Email = ""
		input.Password = ""
		useCase := New(userMock)
		res, err := useCase.PostData(input)
		assert.Equal(t, -1, res)
		assert.Error(t, err)
		userMock.AssertExpectations(t)

	})

}

func TestGetById(t *testing.T) {

	userMock := new(mocks.UserData)
	returnData := users.UserCore{Name: "Hery", ID: 1, Email: "Hery@gmail.com", Password: "asdf"}
	param := 1

	t.Run("Get by id success", func(t *testing.T) {
		userMock.On("SelectById", param).Return(returnData, nil).Once()

		useCase := New(userMock)
		res, _ := useCase.GetById(param)
		assert.Equal(t, param, int(res.ID))
		userMock.AssertExpectations(t)

	})

	t.Run("Get by id failed", func(t *testing.T) {

		userMock.On("SelectById", param).Return(users.UserCore{}, errors.New("error")).Once()

		useCase := New(userMock)
		param := 1
		res, err := useCase.GetById(param)
		assert.Error(t, err)
		assert.NotEqual(t, param, int(res.ID))
		userMock.AssertExpectations(t)

	})

}
func TestPut(t *testing.T) {

	userMock := new(mocks.UserData)
	data := users.UserCore{Name: "Hery", ID: 1, Email: "Hery@gmail.com", Password: "asdf"}
	param := 1

	t.Run("update succes", func(t *testing.T) {

		userMock.On("UpdateData", param, data).Return(1, nil).Once()

		useCase := New(userMock)
		res, _ := useCase.PutData(param, data)
		assert.Equal(t, 1, res)
		userMock.AssertExpectations(t)

	})

	t.Run("update failed", func(t *testing.T) {

		userMock.On("UpdateData", param, data).Return(-1, errors.New("error")).Once()

		useCase := New(userMock)
		res, err := useCase.PutData(param, data)
		assert.Equal(t, -1, res)
		assert.Error(t, err)
		userMock.AssertExpectations(t)

	})

}

func TestGetAll(t *testing.T) {

	userMock := new(mocks.UserData)
	returnData := []users.UserCore{{Name: "Hery", ID: 1, Email: "Hery@gmail.com", Password: "asdf"}}

	t.Run("Get All Success", func(t *testing.T) {

		userMock.On("SelectAll").Return(returnData, nil).Once()

		useCase := New(userMock)
		res, err := useCase.GetAll()
		assert.NoError(t, err)
		assert.Equal(t, returnData[0].ID, res[0].ID)
		userMock.AssertExpectations(t)

	})

	t.Run("Get All Failed", func(t *testing.T) {
		userMock.On("SelectAll").Return(nil, nil).Once()

		useCase := New(userMock)
		res, _ := useCase.GetAll()
		assert.Equal(t, 0, len(res))
		userMock.AssertExpectations(t)

	})

	t.Run("Get All Failed", func(t *testing.T) {
		userMock.On("SelectAll").Return(nil, errors.New("error")).Once()

		useCase := New(userMock)
		_, err := useCase.GetAll()
		assert.Error(t, err)
		userMock.AssertExpectations(t)

	})

}

func TestDelete(t *testing.T) {

	userMock := new(mocks.UserData)

	param := 1

	t.Run("delete succes", func(t *testing.T) {

		userMock.On("DelData", param).Return(1, nil).Once()

		useCase := New(userMock)
		res, _ := useCase.DeleteData(param)
		assert.Equal(t, 1, res)
		userMock.AssertExpectations(t)

	})

	t.Run("delete failed", func(t *testing.T) {

		userMock.On("DelData", param).Return(-1, errors.New("error")).Once()

		useCase := New(userMock)
		res, err := useCase.DeleteData(param)
		assert.Equal(t, -1, res)
		assert.Error(t, err)
		userMock.AssertExpectations(t)

	})

}
