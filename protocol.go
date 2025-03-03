package protocol

import (
	"context"
	"net/http"
)

type UserGetter = func(r *http.Request) (*int64, error)
type AppKeyGetter = func(r *http.Request) (string, error)
type AppKeyValidator = func(ctx context.Context, key string, appID uint32) (int8, error)
type UserIDValidator = func(ctx context.Context, userID int64, appID uint32) (bool, error)

const Version = "0.0.2"

type Initializable interface {
	Init() error
}

type UserGetterPlugin interface {
	GetUserID(r *http.Request) (*int64, error)
}

type AppKeyGetterPlugin interface {
	GetAppKey(r *http.Request) (string, error)
}

type AppKeyValidationPlugin interface {
	ValidateKey(ctx context.Context, key string, appID uint32) (int8, error)
}

type UserIDValidatorPlugin interface {
	ValidateUserID(ctx context.Context, userID int64, appID uint32) (bool, error)
}
