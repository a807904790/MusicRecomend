package handler

import (
	"MusicRecomend/User/model"
	Proto "MusicRecomend/proto"
)

type UserHandler struct {

}



func (u *UserHandler) SignIn(request Proto.Request)  Proto.Response {
	var users []model.User
	model.Db.Find(&users)
	var res Proto.Response
	for _, item := range users {
		if item.Name == request.UserName && item.PassWord == request.PassWord {
			res.Code = 0
			res.Msg = "SignIn success"
			return res
		}
	}
	res.Code = -1
	res.Msg = "username or passoword incorret"
	return res
}


