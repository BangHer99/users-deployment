package delivery

import "be12/deploy/features/users"

type Respon struct {
	ID    uint   `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Hp    string `json:"hp" form:"hp"`
	Email string `json:"email" form:"email"`
}

func toResponId(data users.UserCore) Respon {

	return Respon{
		ID:    data.ID,
		Name:  data.Name,
		Hp:    data.Hp,
		Email: data.Email,
	}

}

func toResponList(data []users.UserCore) []Respon {

	var respon []Respon
	for _, v := range data {
		respon = append(respon, Respon{
			ID:    v.ID,
			Name:  v.Name,
			Hp:    v.Hp,
			Email: v.Email,
		})
	}

	return respon
}
