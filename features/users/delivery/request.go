package delivery

import "be12/deploy/features/users"

type UserReq struct {
	Name     string `json:"name" form:"name"`
	Hp       string `json:"hp" form:"hp"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ToCore(req UserReq) users.UserCore {

	return users.UserCore{
		Name:     req.Name,
		Hp:       req.Hp,
		Email:    req.Email,
		Password: req.Password,
	}

}
