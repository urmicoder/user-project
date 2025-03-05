package interfaces

import "urmi-arch/user/domain"

type ServiceInterfaces interface {
	AddUser(user domain.User) (userId string, err error)
	GetUser(id string, sort string) (user domain.User, err error)
	Signin(userId string, password string) (user domain.User, err error)
	GetAdmin(userName string, password string) (user domain.User, err error)
}
