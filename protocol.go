package protocol

import (
	"context"
	"net/http"
)

type (
	UserGetter    = func(ctx context.Context, appID uint32, r *http.Request) (*int64, error)
	FriendsGetter = func(ctx context.Context, appID uint32, userID int64) ([]int64, error)
	AppKeyGetter  = func(ctx context.Context, appID uint32, r *http.Request) (int8, error)
)

const Version = "0.0.3"

type KeyValidity int8

const (
	KeyValidityInvalid   KeyValidity = 0
	KeyValidityUntrusted KeyValidity = 1
	KeyValidityTrusted   KeyValidity = 2
)

type Initializable interface {
	Init() error
}

type UserPlugin interface {
	GetUserID(ctx context.Context, appID uint32, r *http.Request) (*int64, error)
}

type FriendsPlugin interface {
	GetFriends(ctx context.Context, appID uint32, userID int64) ([]int64, error)
}

type AppKeyPlugin interface {
	GetAppKey(ctx context.Context, appID uint32, r *http.Request) (int8, error)
}
