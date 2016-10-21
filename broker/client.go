package broker

import (
	"net"
	"github.com/buptmiao/msgo/msg"
	"sync"
	"fmt"
	"io"
)

type Client struct {
	sync.Mutex
	broker *Broker
	conn net.Conn
	stop chan struct{}
	subscribes map[string]*subscribe
}

func newClient(broker *Broker, conn net.Conn) *Client {
	res := new(Client)

	res.broker = broker
	res.conn = conn
	res.stop = make(chan struct{})
	res.subscribes = make(map[string]*subscribe)

	return res
}

func (c *Client) Run() {
	//
	count := 0
	for {
		// block here
		msgs, err := c.recvMsg()
		if err != nil {
			if err != io.EOF {
				Error.Println(err)
			}
			break
		}
		count++
		err = c.handle(msgs)
		if err != nil {
			Error.Println(err)
			c.Close()
			return
		}
	}
	c.Close()
}

// return err if it is needed to close client
func (c *Client) handle(m *msg.MessageList) error {
	var err error
	var needAck bool
	for _, v := range m.Msgs {
		switch v.GetType() {
		case msg.MessageType_Subscribe:
			c.handleSubscribe(v)
			needAck = true
		case msg.MessageType_Publish:
			c.handlePublish(v)
			needAck = true
		case msg.MessageType_UnSubscribe:
			c.handleUnSubscribe(v)
		case msg.MessageType_Ack:
			c.handleAck(v)
		case msg.MessageType_Heartbeat:
			c.handleHeartbeat(v)
		default:
			panic(fmt.Errorf("unsupported msg type %v", v.GetType()))
		}
	}
	if needAck {
		err = c.Ack()
	}
	return err
}

func (c *Client) Ack() error {
	ack := msg.NewAckMsg()
	return c.sendMsg(ack)
}

func (c *Client) handleSubscribe(m *msg.Message) {
	sub, ok := c.subscribes[m.GetTopic()]
	// new subscribe
	if !ok {
		Log.Printf("%v subscribe topic %s\n", c.conn.RemoteAddr(), m.GetTopic())
		topic := c.broker.Get(m.GetTopic())
		sub = newsubscribe(topic, c, m.GetFilter(), m.GetCount(), m.GetNeedAck())
		c.subscribes[m.GetTopic()] = sub
	} else {
		sub.update(m.GetFilter(), m.GetCount())
	}
}

func (c *Client) handlePublish(m *msg.Message) {
	if m.GetPersist() {
		err := c.broker.stable.Save(m)
		if err != nil {
			Error.Println(err)
			return
		}
	} else {
		topic := c.broker.Get(m.GetTopic())
		topic.Push(m)
	}
}

func (c *Client) handleUnSubscribe(m *msg.Message) {
	if sub, ok := c.subscribes[m.GetTopic()]; ok {
		Log.Printf("%v unsubscribe topic %s\n", c.conn.RemoteAddr(), m.GetTopic())
		delete(c.subscribes, m.GetTopic())
		sub.close()
	}
	// if a client subscribe no topics, then close it
	if len(c.subscribes) == 0 {
		c.Close()
	}
}

func (c *Client) handleAck(m *msg.Message) {
	if sub, ok := c.subscribes[m.GetTopic()]; ok {
		sub.pushAck()
	}
}

func (c *Client) handleHeartbeat(m *msg.Message) {
	return
}

func (c *Client) recvMsg() (*msg.MessageList, error) {
	res, err := msg.BatchUnmarshal(c.conn)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) sendMsg(m ...*msg.Message) error {
	err := msg.BatchMarshal(&msg.MessageList{m}, c.conn)
	if err != nil {
		Error.Println(err)
	}
	return err
}

func (c *Client) Close() {
	close(c.stop)
	for _, s := range c.subscribes {
		s.close()
	}
	c.subscribes = make(map[string]*subscribe)
	err := c.conn.Close()
	if err != nil {
		Error.Println(err)
	}
	return
}