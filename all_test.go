package socketio

import (
	"fmt"
	"net/url"
	"testing"
	"time"
)

func TestURL(t *testing.T) {
	raw := "wss://streamer.cryptocompare.com/"
	u, err := url.Parse(raw)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u.Path)
	t.Log(u.RawPath)
	t.Log(u.Query())
	t.Log(u.Scheme)

	u.Path = "/socket.io/"
	q := u.Query()
	q.Add("EIO", "3")
	q.Add("transport", "websocket")
	u.RawQuery = q.Encode()
	t.Log(u.String())
}

func TestSocket(t *testing.T) {
	s, err := Socket("wss://streamer.cryptocompare.com/")
	if err != nil {
		t.Fatal(err)
	}
	s.On(EventError, func(args ...interface{}) {
		err := args[0].(error)
		t.Error(err)
	})
	s.On("m", func(args ...interface{}) {
		fmt.Println(args)
	})
	err = s.Connect()
	defer s.Disconnect()
	if err != nil {
		t.Fatal(err)
	}
	s.Emit("SubAdd", map[string]interface{}{
		"subs": []string{"0~Binance~BTC~USDT"},
	})
	time.Sleep(21 * time.Second)
}
