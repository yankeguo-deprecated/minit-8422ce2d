package minit

import (
	"net"
	"testing"
	"time"
)

func TestNotifyServer(t *testing.T) {
	s := NewNotifyServer()
	var data map[string]string
	s.SetHandler(func(l map[string]string) {
		data = l
	})
	go s.ListenAndServe()
	time.Sleep(time.Second)
	addr, _ := net.ResolveUnixAddr("unixgram", s.NotifySocket())
	conn, _ := net.DialUnix("unixgram", nil, addr)
	conn.Write([]byte(" Key1 = Val1 \r\n Key2 = Val2  "))
	time.Sleep(time.Second)
	t.Log(data)
	if data["Key1"] != "Val1" || data["Key2"] != "Val2" {
		t.Fatal(data)
	}
	s.Close()
	time.Sleep(time.Second)
}
