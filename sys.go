package usys

import (
	"os"
	"os/user"

	"github.com/go-leap/fs"
)

var (
	// Env aliases `os.Getenv` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `usys` to not have to repeatedly pull in and out the extra `os` import.
	Env = os.Getenv

	_userHomeDirPath string
)

// UserHomeDirPath returns the path to the current user's Home directory. It remains cached in-memory after the first call.
func UserHomeDirPath() string {
	dirpath := _userHomeDirPath
	if dirpath == "" {
		if userinfo, err := user.Current(); err == nil && userinfo.HomeDir != "" && ufs.IsDir(userinfo.HomeDir) {
			dirpath = userinfo.HomeDir
		} else if dirpath = os.Getenv("HOME"); dirpath == "" || !ufs.IsDir(dirpath) {
			dirpath = os.Getenv("USERPROFILE")
		}
		if dirpath != "" {
			_userHomeDirPath = dirpath
		}
	}
	return dirpath
}
