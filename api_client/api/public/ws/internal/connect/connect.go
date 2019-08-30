package connect

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
)

const host = "https://api.coin.z.com/ws/public/v1"

type connectionState string

const (
	connectionStateConnected = "Connected"
	connectionStateClosed    = "Closed"
)

// Connection ...
type Connection struct {
	sync.Mutex
	conn  *websocket.Conn
	state *atomic.Value
	ctx   context.Context
	stop  context.CancelFunc
}

// New is...
func New() *Connection {
	conn := &Connection{
		state: &atomic.Value{},
	}
	conn.state.Store(connectionStateClosed)
	ctx, cancelFunc := context.WithCancel(context.Background())
	conn.ctx = ctx
	conn.stop = cancelFunc

	go conn.run()
	return conn
}

func (c *Connection) run() {
	defer func() {
		if c.isConnected() {
			_ = c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			_ = c.conn.Close()
		}
	}()

	for {
		if !c.isConnected() {
			if err := c.dial(); err != nil {
				continue
			}
		}

		select {
		case <-c.ctx.Done():
			return
		}
	}
}

func (c *Connection) isConnected() bool {
	v, ok := c.state.Load().(connectionState)

	if !ok {
		c.state.Store(connectionStateClosed)
		return false
	}

	return v == connectionStateConnected
}

// Send is...
func (c *Connection) Send(msg interface{}) error {
	if !c.isConnected() {
		err := c.dial()
		if err != nil {
			return err
		}
	}

	return c.conn.WriteJSON(msg)
}

// SendByte is....
func (c *Connection) SendByte(msg []byte) error {
	if !c.isConnected() {
		err := c.dial()
		if err != nil {
			return err
		}
	}

	return c.conn.WriteMessage(websocket.TextMessage, msg)
}

func (c *Connection) dial() error {
	c.Lock()
	defer c.Unlock()

	if c.conn != nil {
		_ = c.conn.Close()
		c.state.Store(connectionStateClosed)
	}

	conn, res, err := websocket.DefaultDialer.Dial(host, nil)
	if err != nil {
		return fmt.Errorf("dial error:%v, response:%v", err, res)
	}

	conn.SetReadLimit(1024)
	_ = conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	conn.SetPongHandler(func(appData string) error {
		_ = conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		return nil
	})
	c.conn = conn
	c.state.Store(connectionStateConnected)

	return nil
}

// Close is ...
func (c *Connection) Close() {
	c.stop()
}
