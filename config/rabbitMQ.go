package config

import (
	"fmt"
	"os"
)

type RMQ struct {
	user  string
	pass  string
	addr  string
	port  string
	vhost string

	queueFrom string
	queueTo   []string

	queueToSQL          string
	sessionControlQueue string
}

func newRMQ() *RMQ {
	return &RMQ{
		user:      os.Getenv("RMQ_USER"),
		pass:      os.Getenv("RMQ_PASS"),
		addr:      os.Getenv("RMQ_ADDRESS"),
		port:      os.Getenv("RMQ_PORT"),
		vhost:     os.Getenv("RMQ_VHOST"),
		queueFrom: os.Getenv("RMQ_QUEUE_FROM"),
		queueTo: []string{
			os.Getenv("RMQ_QUEUE_TO"),
		},
		queueToSQL:          os.Getenv("RMQ_QUEUE_TO_SQL"),
		sessionControlQueue: os.Getenv("RMQ_SESSION_CONTROL_QUEUE"),
	}
}

func (c *RMQ) URL() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/%s", c.user, c.pass, c.addr, c.port, c.vhost)
}

func (c *RMQ) QueueFrom() string {
	return c.queueFrom
}

func (c *RMQ) QueueTo() []string {
	return c.queueTo
}

func (c *RMQ) QueueToSQL() string {
	return c.queueToSQL
}

func (c *RMQ) SessionControlQueue() string {
	return c.sessionControlQueue
}
