# Leaderboards Plugin Protocol

The package contains interfaces for the Leaderboards service plugins

## Plugins

> **Note:**
>
> Plugin system supports only POSIX operation systems like Linux, MacOS, *BSD. On Windows you can use [WSL (Windows Subsystem for Linux)](https://learn.microsoft.com/en-us/windows/wsl/) to run the service with plugins.

Plugins allow to change some predefined behavior of the service, including:

1. Application key receiving from the Web API requests
2. Application key validation for the Web API requests
3. UserID receiving for the Client API requests

Single plugin may implement all of the possibilities or just some subset.

### How To Develop a Plugin

Each plugin (dynamic library) may implement one or more plugin interfaces for the service. To implement an interface use `github.com/mechanicum-pro/leaderboards-plugin-protocol` package (run `go get github.com/mechanicum-pro/leaderboards-plugin-protocol`). The package contains all necessary interfaces to implement a plugin:

1. `UserGetterPlugin` - interface for plugins that retrieve a UserID from the HTTP request. The plugin should implement the `GetUserID` method as defined by the `UserGetter` protocol.
   `GetUserID` receives an `*http.Request` and should return a sender's UserID in `*int64` or an `error`. An `error` should be returned only in cases of issues like network problems. If the user is not found or the request lacks the required data, return `nil` for the UserID.

2. `AppKeyGetterPlugin` - interface for plugins that retrieve an Application API Key from the HTTP request. The plugin should implement the `GetAppKey` method as defined by the `AppKeyGetter` protocol.
   `GetAppKey` receives an `*http.Request` and should return an Application API key used for the request.

3. `AppKeyValidationPlugin` - interface for plugins that verify the Application API Key and an AppID and return a value indicating whether the key is trusted for the application. The plugin should implement the `ValidateKey` method, which must satisfy the `AppKeyValidator` protocol.
   `ValidateKey` receives a `key` (`string`) and an `appID` (`uint32`) and expected to return an error or one of the following values:
   * `0` - the `key` is invalid
   * `1` - the `key` is valid but is not trusted for the application
   * `2` - the `key` is valid and trusted for the application

The implemented plugin structures should be assigned to the correspondent global variables:

* `UserGetterPlugin` to the `UserIDGetter`
* `AppKeyGetterPlugin` to the `AppKeyGetter`
* `AppKeyValidationPlugin` to the `AppKeyGetter`

For implementation example see [`_example/plugin.go`](_example/plugin.go)

### How to build a Plugin

In the plugin directory run the following command:

```shell
go build -buildmode=plugin -o your-plugin-name.so plugin.go
```
