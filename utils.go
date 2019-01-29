package minit

import (
	"errors"
	"os"
	"path/filepath"
)

var (
	// SearchPaths search paths of unit files
	// we support only a subset of standard search paths
	// see https://www.freedesktop.org/software/systemd/man/systemd.unit.html
	SearchPaths = []string{
		"/etc/systemd/system",
		"/usr/lib/systemd/system",
		"/usr/local/lib/systemd/system",
	}
)

var (
	// ErrNotFound unit file not found
	ErrNotFound = errors.New("unit file not found")
)

// SearchUnitFile search the unit file
func SearchUnitFile(name string) (ret string, err error) {
	for _, dir := range SearchPaths {
		// combile the full path
		ret = filepath.Join(dir, name)

		// stat the file, check accessability
		var info os.FileInfo
		if info, err = os.Stat(ret); err != nil {
			// ignore not exist error
			if os.IsNotExist(err) {
				continue
			}
			return
		}
		// ignore directory
		if info.IsDir() {
			continue
		}

		// found
		return
	}

	// all tried, not found
	ret = ""
	err = ErrNotFound
	return
}
