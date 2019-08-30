package connect

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"

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
	stop  context.CancelFunc
}

// New is...
func New() *Connection {
	conn := &Connection{
		state: &atomic.Value{},
	}
	conn.state.Store(connectionStateClosed)
	ctx, cancelFunc := context.WithCancel(context.Background())
	conn.stop = cancelFunc

	go conn.run(ctx)
	return conn
}

func (c *Connection) run(ctx context.Context) {
	defer func() {
	}()

	for {
		if !c.isConnected() {
			if err := c.dial(); err != nil {
				continue
			}
		}

		select {
		case <-ctx.Done():
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
		c.conn = nil
		c.state.Store(connectionStateClosed)
	}

	conn, res, err := websocket.DefaultDialer.Dial(host, nil)
	if err != nil {
		return fmt.Errorf("dial error:%v, response:%v", err, res)
	}
	c.conn = conn
	c.state.Store(connectionStateConnected)

	return nil
}

// Close is ...
func (c *Connection) Close() {
	c.stop()
}
