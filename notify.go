package minit

import (
	"net"
	"os"
	"path/filepath"
	"strings"
)

// NotifyServerHandler notify server handler
type NotifyServerHandler func(data map[string]string)

// NotifyServer notify server compatible with systemd
type NotifyServer interface {
	SetHandler(h NotifyServerHandler)
	NotifySocket() string
	ListenAndServe() error
	Close() error
}

type notifyServer struct {
	s string
	c *net.UnixConn
	h NotifyServerHandler
}

// NewNotifyServer create a new notify server
func NewNotifyServer() NotifyServer {
	return &notifyServer{
		s: filepath.Join("/tmp", RandomFilename("minit", "sock", 4)),
	}
}

func (n *notifyServer) SetHandler(h NotifyServerHandler) {
	n.h = h
}

func (n *notifyServer) NotifySocket() string {
	return n.s
}

func (n *notifyServer) Close() (err error) {
	if n.c != nil {
		err = n.c.Close()
	}
	return
}

func (n *notifyServer) ListenAndServe() (err error) {
	// remove socket file
	os.Remove(n.s)
	defer os.Remove(n.s)

	// resolve addr
	var addr *net.UnixAddr
	if addr, err = net.ResolveUnixAddr("unixgram", n.s); err != nil {
		return
	}
	// listen addr
	if n.c, err = net.ListenUnixgram("unixgram", addr); err != nil {
		return
	}
	defer n.c.Close()

	// chmod socket file
	if err = os.Chmod(n.s, 0777); err != nil {
		return
	}

	// the loop
	buf := make([]byte, 4*1024)
	oobbuf := make([]byte, 4*1024)
	for {
		// read message
		var l int
		if l, _, _, _, err = n.c.ReadMsgUnix(buf, oobbuf); err != nil {
			return
		}

		// decode message to map
		data := make(map[string]string)
		lines := strings.Split(string(buf[0:l]), "\n")
		for _, line := range lines {
			splits := strings.SplitN(line, "=", 2)
			if len(splits) == 2 {
				data[strings.TrimSpace(splits[0])] = strings.TrimSpace(splits[1])
			}
		}

		// handle the message
		go n.h(data)
	}
}
