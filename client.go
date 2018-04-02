package socketio

import (
	"math"
	"net/url"
	"sync/atomic"
	"time"

	"github.com/LightKool/socketio-client-go/protocol"
)

const (
	stateOpen uint32 = iota
	stateConnecting
	stateReady
	stateReconnecting
	stateClose
)

type option struct {
	AutoReconnect    bool
	MaxReconnections int32
}

var defaultOption = &option{
	AutoReconnect:    true,
	MaxReconnections: math.MaxInt32,
}

type socketClient struct {
	emitter
	state     uint32
	url       *url.URL
	option    *option
	transprot protocol.Transport
	conn      protocol.Conn
	errChan   chan error
	outChan   chan *protocol.Packet
}

func Socket(urlstring string) (*socketClient, error) {
	u, err := url.Parse(urlstring)
	if err != nil {
		return nil, err
	}
	u.Path = "/socket.io/"
	q := u.Query()
	q.Add("EIO", "3")
	q.Add("transport", "websocket")
	u.RawQuery = q.Encode()
	return &socketClient{
		emitter:   emitter{listeners: make(map[string][]Listener)},
		url:       u,
		option:    defaultOption,
		transprot: protocol.NewWebSocketTransport(),
	}, nil
}

func (s *socketClient) Connect() (err error) {
	if atomic.CompareAndSwapUint32(&s.state, stateOpen, stateConnecting) {
		s.conn, err = s.transprot.Dial(s.url.String())
		if err != nil {
			atomic.StoreUint32(&s.state, stateClose)
			return
		}
		s.errChan = make(chan error, 1)
		s.outChan = make(chan *protocol.Packet, 64)
		if atomic.CompareAndSwapUint32(&s.state, stateConnecting, stateReady) {
			go s.start()
		}
	}
	return
}

func (s *socketClient) Disconnect() {
	atomic.StoreUint32(&s.state, stateClose)
	err := s.conn.Close()
	if err != nil {
		s.emit(EventError, err)
	}
	close(s.errChan)
	close(s.outChan)
}

func (s *socketClient) Emit(event string, args ...interface{}) {
	if atomic.LoadUint32(&s.state) == stateReady && !s.emit(event, args) {
		m := &protocol.Message{
			Type:      protocol.MessageTypeEvent,
			Namespace: "/",
			ID:        -1,
			Event:     event,
			Payloads:  args,
		}
		p, err := m.Encode()
		if err != nil {
			s.emit(EventError, err)
		} else {
			s.outChan <- p
		}
	}
}

func (s *socketClient) reconnect() {
	if atomic.CompareAndSwapUint32(&s.state, stateReady, stateReconnecting) {
		conn, err := s.transprot.Dial(s.url.String())
		for err != nil {
			s.emit(EventError, err)
			time.Sleep(time.Second)
			conn, err = s.transprot.Dial(s.url.String())
		}
		old := s.conn
		s.conn = conn
		go old.Close()
		if atomic.CompareAndSwapUint32(&s.state, stateReconnecting, stateReady) {
			go s.start()
		}
	}
}

func (s *socketClient) start() {
	go s.startRead()
	go s.startWrite()
	for err := range s.errChan {
		s.emit(EventError, err)
		go s.reconnect()
	}
}

func (s *socketClient) startRead() {
	defer func() {
		recover()
	}()
	for atomic.LoadUint32(&s.state) == stateReady {
		p, err := s.conn.Read()
		if err != nil {
			s.errChan <- err
			return
		}
		switch p.Type {
		case protocol.PacketTypeOpen:
			h, err := p.DecodeHandshake()
			if err != nil {
				s.emit(EventError, err)
			} else {
				go s.startPing(h)
			}
		case protocol.PacketTypePing:
			s.outChan <- protocol.NewPongPacket()
		case protocol.PacketTypeMessage:
			m, err := p.DecodeMessage()
			if err != nil {
				s.emit(EventError, err)
			} else {
				s.emit(m.Event, m.Payloads...)
			}
		}
	}
}

func (s *socketClient) startWrite() {
	defer func() {
		recover()
	}()
	for atomic.LoadUint32(&s.state) == stateReady {
		p, ok := <-s.outChan
		if !ok {
			return
		}
		err := s.conn.Write(p)
		if err != nil {
			s.errChan <- err
			return
		}
	}
}

func (s *socketClient) startPing(h *protocol.Handshake) {
	defer func() {
		recover()
	}()
	for {
		time.Sleep(time.Duration(h.PingInterval) * time.Millisecond)
		if atomic.LoadUint32(&s.state) != stateReady {
			return
		}
		s.outChan <- protocol.NewPingPacket()
	}
}
