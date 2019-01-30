package minit

import (
	"io"
	"os"
)

// Service service unit file
type Service struct {
	Type                 string
	ExecStart            string
	ExecStop             string
	User                 string
	Group                string
	RuntimeDirectory     string
	RuntimeDirectoryMode os.FileMode
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
