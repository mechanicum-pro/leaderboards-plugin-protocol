package main

import (
	"errors"
	"net/http"
	"strconv"

	protocol "github.com/mechanicum-pro/leaderboards-plugin-protocol"
)

type ExampleUserGetterPlugin struct {
}

type ExampleValidateAppKeyPlugin struct {
}

type ExampleAppKeyGetterPlugin struct {
}

// Compile time checks for
// the plugins implements protocol.HttpRedirectPlugin.
var _ protocol.UserGetterPlugin = &ExampleUserGetterPlugin{}
var _ protocol.AppKeyValidationPlugin = &ExampleValidateAppKeyPlugin{}
var _ protocol.AppKeyGetterPlugin = &ExampleAppKeyGetterPlugin{}

// GetUserID returns the user id from the request's `X-User-Id` HTTP header.
func (p *ExampleUserGetterPlugin) GetUserID(r *http.Request) (*int64, error) {
	rawUserID := r.Header.Get("X-User-Id")
	if rawUserID != "" {
		userID, err := strconv.ParseInt(rawUserID, 10, 64)
		if err != nil {
			return nil, err
		}
		return &userID, nil
	}
	return nil, errors.New("X-User-Id header is required")
}

// GetAppKey returns the application API key from the request's
// `X-App-Key` HTTP header or `key` form/query value.
func (p *ExampleAppKeyGetterPlugin) GetAppKey(r *http.Request) (string, error) {
	key := r.FormValue("key")

	if key != "" {
		return key, nil
	}
	key = r.Header.Get("X-App-Key")
	if key != "" {
		return key, nil
	}
	return "", errors.New("key is required")
}

// Example Application API key validator. Each non-empty key is considered trusted in exception of `untrusted`.
// Empty key is considered invalid.
func (p *ExampleValidateAppKeyPlugin) ValidateKey(key string, appID uint32) (int8, error) {
	if key == "" {
		return 0, errors.New("key is required")
	}
	if key == "untrusted" {
		return 1, nil
	}
	return 2, nil
}

// Export the plugins. The variable names are matter for the plugin system.
var UserIDGetter = ExampleUserGetterPlugin{}
var AppKeyGetter = ExampleAppKeyGetterPlugin{}
var AppKeyValidator = ExampleValidateAppKeyPlugin{}

func main() {}
