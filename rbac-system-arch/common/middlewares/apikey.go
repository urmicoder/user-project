package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
	"urmi-arch/common"
	"urmi-arch/common/settings"
)

var APIKEYS map[string]string

func APIKEYS_INIT() {
	APIKEYS = make(map[string]string)
}

func APIKeyAddedInDB(key, val string) {
	APIKEYS[key] = val
}
func GetAPIKeyFromDB(key string) string {
	return APIKEYS[key]
}

func GenerateApiKey(role string) string {
	now := time.Now().UnixNano()
	hashed := common.Encrypt(strings.ToUpper(role))
	apikey := fmt.Sprintf("%v-%v", now, hashed)
	return apikey
}

func CreateAPIKey(roles []string) string {
	isAdmin := false
	isSuper := false
	isUser := false
	isGuest := false
	for _, role := range roles {
		if role == settings.ADMINROLE {
			isAdmin = true
		}
		if role == settings.SUPERADMINROLE {
			isSuper = true
		}
		if role == settings.USERROLE {
			isUser = true
		}
		if role == settings.GUESTROLE {
			isGuest = true
		}
	}
	if isSuper {
		key := GenerateApiKey(settings.SUPERADMINROLE)
		APIKeyAddedInDB(key, "SUPERADMINROLE")
		return key
	}
	if isAdmin {
		key := GenerateApiKey(settings.ADMINROLE)
		APIKeyAddedInDB(key, "ADMINROLE")
		return key
	}
	if isGuest {
		key := GenerateApiKey(settings.GUESTROLE)
		APIKeyAddedInDB(key, "GUESTROLE")
		return key
	}
	if isUser {
		key := GenerateApiKey(settings.USERROLE)
		APIKeyAddedInDB(key, "USERROLE")
		return key
	}
	return ""
}

func APPIKeyCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("API-KEY")
		if apiKey == "" {
			http.Error(w, common.ErrAPIKey.Error(), http.StatusBadRequest)
			return
		}

		role := GetAPIKeyFromDB(apiKey)
		if role == "" {
			http.Error(w, common.ErrAPIKey.Error(), http.StatusBadRequest)
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), "role", role))
		r = r.WithContext(context.WithValue(r.Context(), "isApiKey", true))
		fmt.Println("API KEY Called: apiKey", apiKey, " ROle ", role)
		next.ServeHTTP(w, r)
	})
}
