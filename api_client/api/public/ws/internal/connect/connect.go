package connect

import (
	"context"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
)

const host = "wss://api.coin.z.com/ws/public/v1"

type connectionState string

const (
	connectionStateConnecting = connectionState("Connecting")
	connectionStateConnected  = connectionState("Connected")
	connectionStateClosed     = connectionState("Closed")
)

// Connection ...
type Connection struct {
	sync.Mutex
	conn   *websocket.Conn
	state  *atomic.Value
	ctx    context.Context
	stop   context.CancelFunc
	stream chan []byte
}

// New is...
func New() *Connection {
	ctx, cancelFunc := context.WithCancel(context.Background())
	conn := &Connection{
		state:  &atomic.Value{},
		ctx:    ctx,
		stop:   cancelFunc,
		stream: make(chan []byte, 100),
	}
	conn.state.Store(connectionStateClosed)

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
				log.Println(err)
				continue
			}
		}

		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			log.Println(fmt.Sprintf("[ReadMessage]error:%v", err))
			_ = c.conn.Close()
			c.state.Store(connectionStateClosed)
			continue // TODO:review
		}

		c.stream <- msg
		select {
		case <-c.ctx.Done():
			return
		default:
		}
	}
}

func (c *Connection) isConnected() bool {
	c.Lock()
	defer c.Unlock()
	v, ok := c.state.Load().(connectionState)

	if !ok {
		c.state.Store(connectionStateClosed)
		return false
	}

	return v == connectionStateConnected || v == connectionStateConnecting
}

// Send is...
func (c *Connection) Send(msg interface{}) error {
	err := c.waitForConnected()
	if err != nil {
		return err
	}

	err = c.conn.WriteJSON(msg)
	if err != nil {
		return fmt.Errorf("write error:%v", err)
	}
	log.Printf("[Send]msg:%+v", msg)
	return nil
}

// SendByte is....
func (c *Connection) SendByte(msg []byte) error {
	if !c.isConnected() {
		err := c.dial()
		if err != nil {
			return fmt.Errorf("dial error:%v", err)
		}
	}

	err := c.conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		return fmt.Errorf("write error:%v", err)
	}
	return nil
}

func (c *Connection) waitForConnected() error {
	if c.isConnected() {
		return nil
	}

	for i := 0; i < 10; i++ {
		if c.isConnected() {
			return nil
		}
		time.Sleep(time.Second)
	}
	return fmt.Errorf("connection timeout")
}

func (c *Connection) dial() error {
	c.Lock()
	defer c.Unlock()

	log.Println("dial start")
	if c.conn != nil {
		c.state.Store(connectionStateClosed)
		_ = c.conn.Close()
	}

	c.state.Store(connectionStateConnecting)
	conn, res, err := websocket.DefaultDialer.Dial(host, nil)
	if err != nil {
		c.state.Store(connectionStateClosed)
		return fmt.Errorf("dial error:%v, response:%v", err, res)
	}

	conn.SetReadLimit(1024)
	_ = conn.SetReadDeadline(time.Now().Add(120 * time.Second))
	//conn.SetPongHandler(func(appData string) error {
	//	_ = conn.SetReadDeadline(time.Now().Add(120 * time.Second))
	//	return nil
	//})
	c.conn = conn
	c.state.Store(connectionStateConnected)

	return nil
}

// Stream ...
func (c *Connection) Stream() <-chan []byte {
	return c.stream
}

// Close is ...
func (c *Connection) Close() {
	c.stop()
}
