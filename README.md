# Leaderboards Plugin Protocol

The package contains interfaces for the Leaderboards service plugins

## Plugins

> **Note:**
>
> Plugin system supports only POSIX operation systems like Linux, MacOS, *BSD. On Windows you can use [WSL (Windows Subsystem for Linux)](https://learn.microsoft.com/en-us/windows/wsl/) to run the service with plugins.

Plugins allow to change some predefined behavior of the service, including:

1. Application key receiving from the Web API requests and it's validation
2. UserID receiving and validating for the Client API requests
3. Friends list retrieving for a user

Single plugin may implement all of the possibilities or just some subset.

### How To Develop a Plugin

Each plugin (dynamic library) may implement one or more plugin interfaces for the service. To implement an interface use `github.com/mechanicum-pro/leaderboards-plugin-protocol` package (run `go get github.com/mechanicum-pro/leaderboards-plugin-protocol`). The package contains all necessary interfaces to implement a plugin:

1. `UserPlugin` - interface for plugins that retrieve a UserID from the HTTP request. The plugin should implement the `GetUserID` method as defined by the `UserGetter` protocol.

2. `AppKeyPlugin` - interface for plugins that retrieve an Application API Key from the HTTP request. The plugin should implement the `GetAppKey` method as defined by the `AppKeyGetter` protocol.
   * `0` - the `key` is invalid
   * `1` - the `key` is valid but is not trusted for the application
   * `2` - the `key` is valid and trusted for the application

3. `FriendsPlugin` - interface for plugins that retrieve list of friends of the user within the application 

The implemented plugin structures should be assigned to the correspondent global variables:

* `UserPlugin` to the `UserPlugin`
* `AppKeyPlugin` to the `AppKeyPlugin`
* `FriendsPlugin` to the `FriendsPlugin`

For implementation example see [`_example/plugin.go`](_example/plugin.go)

### How to make a stateful Plugin

In some cases, it may be useful to have some state in the plugin. For example, for caching purposes or to load some configuration and keep it during the lifecycle. To achieve this, you can implement an `Initializable` interface for your plugin with a single method `Init() error` which is called once after the plugin is loaded.

> **Attention**: if the Init method returns an error, the application will be terminated.

### How to build a Plugin

In the plugin directory run the following command:

```shell
go build -buildmode=plugin -o your-plugin-name.so plugin.go
```
