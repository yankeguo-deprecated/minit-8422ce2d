package minit

import (
	"io"
)

// Service service unit file
type Service struct {
	Type string

	PIDFile string

	ExecStartPre  string
	ExecStart     string
	ExecStartPost string

	ExecStop     string
	ExecStopPost string

	Restart string

	// Paths
	RootDirectory    string
	WorkingDirectory string

	// Credentials
	User  string
	Group string

	// Sandboxing
	RuntimeDirectory           string
	RuntimeDirectoryMode       string
	StateDirectory             string
	StateDirectoryMode         string
	CacheDirectory             string
	CacheDirectoryMode         string
	LogsDirectory              string
	LogsDirectoryMode          string
	ConfigurationDirectory     string
	ConfigurationDirectoryMode string

	// Environment
	Environments []string
}

func (s *Service) assign(sec, key, val string) {
}

func (s *Service) validate() (err error) {
	return
}

// ReadFrom read a service definition from a unit file reader
func (s *Service) ReadFrom(r Reader) (err error) {
	var sec, key, val string

	// iterate k-v
	for {
		if sec, key, val, err = r.Next(); err != nil {
			if err != io.EOF {
				return
			}
			err = nil
			break
		}
		s.assign(sec, key, val)
	}

	// validate values
	err = s.validate()

	return
}
