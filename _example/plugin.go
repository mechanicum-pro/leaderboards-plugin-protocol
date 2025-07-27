package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"slices"
	"strconv"
	"strings"

	protocol "github.com/mechanicum-pro/leaderboards-plugin-protocol"
)

type ExampleUserPlugin struct {
}

type ExampleAppKeyPlugin struct {
	untrustedAppKeys []string
}

// Compile time checks for
// the plugins implements protocol.HttpRedirectPlugin.
var _ protocol.UserPlugin = &ExampleUserPlugin{}
var _ protocol.AppKeyPlugin = &ExampleAppKeyPlugin{}

// GetUserID returns the user id from the request's `X-User-Id` HTTP header.
func (p *ExampleUserPlugin) GetUserID(ctx context.Context, appID uint32, r *http.Request) (*int64, error) {
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
func (p *ExampleAppKeyPlugin) GetAppKey(ctx context.Context, appID uint32, r *http.Request) (int8, error) {
	key := r.FormValue("key")
	if key == "" {
		key = r.Header.Get("X-App-Key")
	}
	if key != "" {
		if slices.Contains(p.untrustedAppKeys, key) {
			return 1, nil
		}
		return 2, nil
	}
	return 0, errors.New("key is required")
}

// Example Application API key validator. Each non-empty key is considered trusted in exception of `untrusted`.
// Empty key is considered invalid.
func (p *ExampleAppKeyPlugin) ValidateKey(ctx context.Context, key string, appID uint32) (int8, error) {
	if key == "" {
		return 0, errors.New("key is required")
	}
	if slices.Contains(p.untrustedAppKeys, key) {
		return 1, nil
	}
	return 2, nil
}

// Example Application API key validator also implements protocol.Initializable.
// It loads the list of untrusted keys from the `UNTRUSTED_APP_KEYS` environment variable.
// The keys are comma-separated. If the environment variable is not set, the default value
// is a single value `untrusted`.
func (p *ExampleAppKeyPlugin) Init() error {
	untrustedAppKeys, ok := os.LookupEnv("UNTRUSTED_APP_KEYS")
	if !ok {
		untrustedAppKeys = "untrusted"
	}
	if untrustedAppKeys != "" {
		p.untrustedAppKeys = strings.Split(untrustedAppKeys, ",")
	}
	return nil
}

// Export the plugins. The variable names are matter for the plugin system.
var UserPlugin = ExampleUserPlugin{}
var AppKeyPlugin = ExampleAppKeyPlugin{}

func main() {}
