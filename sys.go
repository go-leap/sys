package usys

import (
	"os"
	"os/user"
	"path/filepath"
	"strconv"

	"github.com/go-leap/fs"
)

var (
	// Env aliases `os.Getenv` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `usys` to not have to repeatedly pull in and out the extra `os` import.
	Env = os.Getenv

	_userDataDirPaths = make(map[bool]string, 2)
	_userHomeDirPath  string
)

func EnvBool(name string, defaultValue bool) bool {
	if envstr := os.Getenv(name); envstr != "" {
		if envbool, e := strconv.ParseBool(envstr); e == nil {
			defaultValue = envbool
		}
	}
	return defaultValue
}

// UserDataDirPath looks for the user's local configuration or cache directory, probing for a variety of common folder-name idioms and environment variables.
// If no match, it returns the result of `UserHomeDirPath`.
func UserDataDirPath(preferCacheOverConfig bool) string {
	dirpath := _userDataDirPaths[preferCacheOverConfig]
	if dirpath == "" {
		probeenvvars := []string{"XDG_CONFIG_HOME", "XDG_CACHE_HOME", "LOCALAPPDATA", "APPDATA"}
		if preferCacheOverConfig {
			probeenvvars[0], probeenvvars[1] = probeenvvars[1], probeenvvars[0]
		}
		for _, envvar := range probeenvvars {
			if maybedirpath := os.Getenv(envvar); maybedirpath != "" && ufs.IsDir(maybedirpath) {
				dirpath = maybedirpath
				break
			}
		}

		if homedirpath := UserHomeDirPath(); dirpath == "" && homedirpath != "" {
			probehomesubdirs := []string{".config", ".cache", "Library/Application Support", "Library/Caches"}
			if preferCacheOverConfig {
				probehomesubdirs[0], probehomesubdirs[1] = probehomesubdirs[1], probehomesubdirs[0]
				probehomesubdirs[2], probehomesubdirs[3] = probehomesubdirs[3], probehomesubdirs[2]
			}
			for _, homesubdir := range probehomesubdirs {
				if maybedirpath := filepath.Join(homedirpath, homesubdir); ufs.IsDir(maybedirpath) {
					dirpath = maybedirpath
					break
				}
			}
			if dirpath == "" {
				dirpath = homedirpath
			}
		}
		_userDataDirPaths[preferCacheOverConfig] = dirpath
	}
	return dirpath
}

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
