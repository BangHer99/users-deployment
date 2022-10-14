package users

import "time"

type UserCore struct {
	ID        uint
	Name      string
	Hp        string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type ServiceInterface interface {
	GetAll() (data []UserCore, err error)
	GetById(param int) (data UserCore, err error)
	PostData(data UserCore) (int, error)
	PutData(param int, data UserCore) (int, error)
	DeleteData(param int) (int, error)
}

type DataInterface interface {
	SelectAll() (data []UserCore, err error)
	SelectById(param int) (data UserCore, err error)
	CreateData(data UserCore) (int, error)
	UpdateData(param int, data UserCore) (int, error)
	DelData(param int) (int, error)
}
