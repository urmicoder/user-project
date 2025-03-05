package common

import "errors"

var (
	ErrUserCredWrong = errors.New("user invailed credentails, please use correct username password")
	ErrRateLimiting  = errors.New("you are sending too many request, please wait sometime")
	ErrAPIKey        = errors.New("api key is missing")
)
