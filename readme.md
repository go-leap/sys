# usys
--
    import "github.com/go-leap/sys"


## Usage

```go
var (
	// Env aliases `os.Getenv` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `usys` to not have to repeatedly pull in and out the extra `os` import.
	Env = os.Getenv
)
```

#### func  Arg

```go
func Arg(i int) (argValOrEmpty string)
```

#### func  EnvBool

```go
func EnvBool(name string, defaultValue bool) bool
```
EnvBool returns `defaultValue` unless there was an environment variable with the
given `name` and a value that `strconv.ParseBool`ed successfully.

#### func  OnSigint

```go
func OnSigint(do func())
```

#### func  UserDataDirPath

```go
func UserDataDirPath(preferCacheOverConfig bool) string
```
UserDataDirPath looks for the user's local configuration or cache directory,
probing for a variety of common folder-name idioms and environment variables. If
no match, it returns the result of `UserHomeDirPath`.

#### func  UserHomeDirPath

```go
func UserHomeDirPath() string
```
UserHomeDirPath returns the path to the current user's Home directory. It
remains cached in-memory after the first call.
