package protocol

import (
	"net/http"
)

type UserGetter = func(r *http.Request) (*int64, error)
type AppKeyGetter = func(r *http.Request) (string, error)
type AppKeyValidator = func(key string, appID uint32) (int8, error)

const Version = "0.0.1"

type UserGetterPlugin interface {
	GetUserID(r *http.Request) (*int64, error)
}

type AppKeyGetterPlugin interface {
	GetAppKey(r *http.Request) (string, error)
}

type AppKeyValidationPlugin interface {
	ValidateKey(key string, appID uint32) (int8, error)
}
