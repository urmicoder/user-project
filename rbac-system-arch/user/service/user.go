package service

import (
	"errors"
	"strings"
	"urmi-arch/common/settings"
	"urmi-arch/user/domain"
	"urmi-arch/user/interfaces"
)

type service struct {
	db interfaces.RepositoryInterfaces
}

func NewUserService(repo interfaces.RepositoryInterfaces) *service {
	return &service{
		db: repo,
	}
}

func (svc *service) AddUser(user domain.User) (userId string, err error) {
	if strings.TrimSpace(user.Name) == "" || strings.TrimSpace(user.Password) == "" {
		return "", errors.New("name and password cannot be empty")
	}

	var roles []string
	if !strings.Contains(user.Email, "la.com") {
		for _, role := range user.Roles {
			if role != settings.ADMINROLE && role != settings.SUPERADMINROLE {
				roles = append(roles, role)
			}
		}
		user.Roles = roles
	}
	svc.db.AddUser(user)
	return
}
func (svc *service) GetUser(id string, sort string) (user domain.User, err error) {
	svc.db.GetUser(id, sort)
	return
}
func (svc *service) Signin(userId string, password string) (user domain.User, err error) {
	user, err = svc.db.Signin(userId)
	if err != nil {
		return
	}
	// matched, err := common.Decrypt(user.Password, password)
	// if !matched {
	// 	return domain.User{}, common.ErrUserCredWrong
	// }
	return
}

func (svc *service) GetAdmin(userName string, password string) (user domain.User, err error) {
	svc.db.GetAdmin(userName)
	if err != nil {
		return
	}
	// matched, err := common.Decrypt(user.Password, password)  //When I add the DB, I will uncomment the code.
	// if !matched {
	// 	return domain.User{}, common.ErrUserCredWrong
	// }
	return
}
