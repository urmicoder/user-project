package handler

import (
	"fmt"
	"net/http"
	"urmi-arch/common/middlewares"
	"urmi-arch/common/settings"
	"urmi-arch/user/domain"
	"urmi-arch/user/interfaces"

	"github.com/gorilla/mux"
)

var ISAPIKEY bool

type handler struct {
	svc interfaces.ServiceInterfaces
}

func newHandler(svc interfaces.ServiceInterfaces) *handler {
	return &handler{
		svc: svc,
	}
}

// AddUser
func (h *handler) AddUser(w http.ResponseWriter, r *http.Request) {
	user, err := DecodeAddUser(r)
	if err != nil {
		EncodeAddUser(w, "", err)
		return
	}
	id, err := h.svc.AddUser(user)
	if err != nil {
		EncodeAddUser(w, "", err)
		return
	}
	EncodeAddUser(w, id, nil)
}

// GetUser
func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, sort, err := DecodeGetUser(r)
	if err != nil {
		EncodeAddUser(w, "", err)
		return
	}
	user, err := h.svc.GetUser(id, sort)
	if err != nil {
		EncodeGetUser(w, domain.User{}, err)
		return
	}
	EncodeGetUser(w, user, nil)
}

// SignInUser
func (h *handler) SignInUser(w http.ResponseWriter, r *http.Request) {
	userId, password, err := DecodeSignIn(r)
	if err != nil {
		EncodeSignIn(w, "", err)
		return
	}
	user, err := h.svc.Signin(userId, password)
	if err != nil {
		EncodeSignIn(w, "", err)
		return
	}

	if ISAPIKEY {
		fmt.Println("APPI KEY")
		key := middlewares.CreateAPIKey(user.Roles)
		w.Header().Set("API-KEY", key)
	} else {
		token, err := middlewares.CreateToken(user.Id, user.Roles)
		if err != nil {
			EncodeSignIn(w, "", err)
			return
		}
		w.Header().Set("Authorization", "Bearer "+token)
	}

	EncodeSignIn(w, user.Id, nil)
}

// GetAdmin
func (h *handler) GetAdmin(w http.ResponseWriter, r *http.Request) {
	username, password, err := DecodeGetAdmin(r)
	if err != nil {
		EncodeAddUser(w, "", err)
		return
	}
	user, err := h.svc.GetAdmin(username, password)
	if err != nil {
		EncodeGetAdmin(w, domain.User{}, err)
		return
	}
	EncodeGetAdmin(w, user, nil)
}
func UserRoutes(svc interfaces.ServiceInterfaces, rl middlewares.RateLimiting, isApiKey bool) *mux.Router {
	h := newHandler(svc)
	r := mux.NewRouter()
	r.HandleFunc("/user/add", h.AddUser).Methods(http.MethodPost)
	if !isApiKey {
		r.Handle("/user/get", middlewares.AuthMiddleware(rl.RateLimiting(http.HandlerFunc(h.GetUser)), []string{settings.ADMINROLE})).Methods(http.MethodGet)
	} else {
		r.Handle("/user/get", middlewares.APPIKeyCheck(rl.RateLimiting(http.HandlerFunc(h.GetUser)))).Methods(http.MethodGet)
	}
	if !isApiKey {
		r.Handle("/user/admin", middlewares.AuthMiddleware(rl.RateLimiting(http.HandlerFunc(h.GetAdmin)), []string{settings.ADMINROLE})).Methods(http.MethodGet)
	} else {
		r.Handle("/user/admin", middlewares.APPIKeyCheck(rl.RateLimiting(http.HandlerFunc(h.GetAdmin)))).Methods(http.MethodGet)
	}
	//r.Handle("/user/admin", middlewares.AuthMiddleware(middlewares.APPIKeyCheck(rl.RateLimiting(http.HandlerFunc(h.GetAdmin))), []string{})).Methods(http.MethodGet)
	ISAPIKEY = isApiKey
	r.HandleFunc("/user/signin", h.SignInUser).Methods(http.MethodPost)
	return r
}
