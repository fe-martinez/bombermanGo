package client

type ClientState interface {
	Handle(c *Client)
}
