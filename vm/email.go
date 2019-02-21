package vm

import (
	"../model"
	"../config"
)

type EmailViewModel struct {
	Username string
	Token    string
	Server   string
}

type EmailViewModelOp struct{}

func (EmailViewModelOp) GetVM(email string) EmailViewModel {
	v := EmailViewModel{}
	u, _ := model.GetUserByEmail(email)
	v.Username = u.Username
	v.Server = config.GetServerURL()

	return v
}
