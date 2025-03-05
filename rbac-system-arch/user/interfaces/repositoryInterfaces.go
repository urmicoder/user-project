package interfaces

import "urmi-arch/user/domain"

type RepositoryInterfaces interface {
	AddUser(user domain.User) (userId string, err error)
	GetUser(id string, sort string) (user domain.User, err error)
	Signin(userId string) (user domain.User, err error)
	GetAdmin(username string) (user domain.User, err error)
}
