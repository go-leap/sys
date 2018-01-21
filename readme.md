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

#### func  UserHomeDirPath

```go
func UserHomeDirPath() string
```
UserHomeDirPath returns the path to the current user's Home directory. It
remains cached in-memory after the first call.
