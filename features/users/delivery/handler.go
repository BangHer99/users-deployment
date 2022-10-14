package delivery

import (
	users "be12/deploy/features/users"
	"strconv"

	// "be12/deploy/utils/helper"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	data users.ServiceInterface
	// date users.UserCore
}

func New(e *echo.Echo, usecase users.ServiceInterface) {

	handler := UserHandler{
		data: usecase,
	}

	e.GET("/users", handler.GetAll)
	e.GET("/users/:id", handler.GetById)
	e.POST("/users", handler.AddUser)
	e.PUT("/users/:id", handler.PutData)
	e.DELETE("/users/:id", handler.Deldata)

}
func (users *UserHandler) Deldata(e echo.Context) error {

	id, _ := strconv.Atoi(e.Param("id"))
	if id == -1 {
		return e.JSON(400, map[string]interface{}{
			"message": "id tidak di temukan",
		})
	}

	row, _ := users.data.DeleteData(id)
	if row == 1 {
		return e.JSON(200, map[string]interface{}{
			"message": "data berhasil di hapus",
		})
	} else {
		return e.JSON(400, map[string]interface{}{
			"message": "kamu tidak bisa akses",
		})
	}

}

func (uh *UserHandler) PutData(e echo.Context) error {

	id, _ := strconv.Atoi(e.Param("id"))
	if id == -1 {
		return e.JSON(400, map[string]interface{}{
			"message": "data tidak di temukan",
		})
	}

	var req UserReq
	err := e.Bind(&req)
	if err != nil {
		return e.JSON(400, map[string]interface{}{
			"message": "gagal bind data",
		})
	}

	var add users.UserCore
	if req.Email != "" {
		add.Email = req.Email
	}
	if req.Hp != "" {
		add.Hp = req.Hp
	}
	if req.Name != "" {
		add.Name = req.Name
	}
	if req.Password != "" {
		add.Password = req.Password
	}

	add.ID = uint(id)

	row, _ := uh.data.PutData(id, add)
	if row == 1 {
		return e.JSON(200, map[string]interface{}{
			"message": "sukses update data",
		})
	} else {
		return e.JSON(400, map[string]interface{}{
			"message": "kamu tidak bisa akses",
		})
	}

}

func (users *UserHandler) GetById(e echo.Context) error {

	idconv, _ := strconv.Atoi(e.Param("id"))

	res, err := users.data.GetById(idconv)
	if err != nil {
		return e.JSON(400, map[string]interface{}{
			"message": "tidak di temukan",
		})
	}

	respon := toResponId(res)

	return e.JSON(200, map[string]interface{}{
		"message": "sukses mendapatkan data by id",
		"data":    respon,
	})

}

func (users *UserHandler) AddUser(e echo.Context) error {

	var req UserReq
	err := e.Bind(&req)
	if err != nil {
		return e.JSON(400, map[string]interface{}{
			"message": " Gagal",
		})
	}

	add := ToCore(req)
	row, _ := users.data.PostData(add)
	if row == 1 {
		return e.JSON(200, map[string]interface{}{
			"message": "sukses membuat data",
		})
	} else {
		return e.JSON(400, map[string]interface{}{
			"message": "gagal membuat data",
		})
	}

}

func (users *UserHandler) GetAll(e echo.Context) error {

	res, err := users.data.GetAll()
	if err != nil {
		return e.JSON(400, map[string]interface{}{
			"message": "Gagal Melihat Data",
		})
	}

	respon := toResponList(res)

	return e.JSON(200, map[string]interface{}{
		"message": "Berhasil Melihat Data",
		"data":    respon,
	})

}
