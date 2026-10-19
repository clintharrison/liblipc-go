//go:build go1.27

package lipc

import (
	"context"

	"github.com/godbus/dbus/v5"
)

type Client struct {
	conn *dbus.Conn
}

func NewClient(dbusConn *dbus.Conn) *Client {
	return &Client{conn: dbusConn}
}

func (c *Client) Get[T string | int32](ctx context.Context, service, property string) (T, error) {
	return LipcGetProperty[T](ctx, c.conn, service, property)
}

func (c *Client) Set[T string | int32](ctx context.Context, service, property string, value T) (error) {
	return LipcSetProperty[T](ctx, c.conn, service, property, value)
}
