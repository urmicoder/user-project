package repositories

import (
	"log"
	"strconv"
	"urmi-arch/common/settings"
	"urmi-arch/user/domain"
)

var users = []domain.User{}
var counterId = 1

type repo struct{}

func NewUserRepo() *repo {
	return &repo{}
}

func (r *repo) AddUser(user domain.User) (userId string, err error) {
	user.Id = strconv.Itoa(counterId)
	counterId++
	users = append(users, user)
	log.Println(user)
	return
}
func (r *repo) GetUser(id string, sort string) (user domain.User, err error) {
	log.Println(id, sort)
	user = domain.User{
		Name:     "urmi",
		Password: "urmi@la",
		Email:    "urmi@la.com",
		Roles:    []string{settings.ADMINROLE},
	}
	return
}
func (r *repo) Signin(userId string) (user domain.User, err error) {
	log.Println(userId)
	user = domain.User{
		Name:     "urmi",
		Password: "urmi@la",
		Email:    "urmi@la.com",
		Roles:    []string{settings.ADMINROLE},
	}
	return
}

func (r *repo) GetAdmin(userName string) (user domain.User, err error) {
	log.Println(userName)
	user = domain.User{
		Name:     "urmi",
		Password: "urmi@la",
		Email:    "urmi@la.com",
		Roles:    []string{settings.ADMINROLE},
	}
	return
}
