// Package objdeliv is the client library
// for objdeliv
package objdeliv

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/satori/uuid"
)

type Client struct {
	addr string
}

type NewObjectOptions struct {
	ID uuid.UUID
	// Expire can only be as precise as second
	Expire time.Duration
}

type newObjectMeta struct {
	Status  string `json:"status"`
	ID      string `json:"id"`
	Message string `json:"message"`
}

type generalResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (c *Client) NewObject(opt *NewObjectOptions) (io.WriteCloser, uuid.UUID, error) {
	endpoint := "/new-object?"
	if !uuid.Equal(opt.ID, uuid.Nil) {
		endpoint += "id=" + opt.ID.String()
		if opt.Expire != 0 {
			endpoint += "&"
		}
	}
	if opt.Expire != 0 {
		endpoint += "expire=" + strconv.Itoa(int(opt.Expire/time.Second))
	}
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return nil, uuid.Nil, err
	}
	_, err = conn.Write([]byte("CONNECT " + endpoint + " HTTP/1.1\r\n\r\n"))
	if err != nil {
		conn.Close()
		return nil, uuid.Nil, err
	}
	buf := bufio.NewReader(conn)
	meta, err := buf.ReadBytes('\n')
	if err != nil {
		conn.Close()
		return nil, uuid.Nil, err
	}
	if string(meta[:8]) == "HTTP/1.1" {
		defer conn.Close()
		return nil, uuid.Nil, fmt.Errorf("connection upgrade error")
	}
	m := &newObjectMeta{}
	err = json.Unmarshal(meta, m)
	if err != nil {
		conn.Close()
		return nil, uuid.Nil, err
	}
	if m.Status == "success" {
		id, err := uuid.FromString(m.ID)
		if err != nil {
			conn.Close()
			return nil, uuid.Nil, err
		}
		return conn, id, nil
	}
	return nil, uuid.Nil, fmt.Errorf(m.Message)
}

func (c *Client) GetObject(id uuid.UUID, autoRelease bool) (io.ReadCloser, error) {
	autoReleaseStr := "false"
	if autoRelease {
		autoReleaseStr = "true"
	}
	resp, err := http.Get("http://" + c.addr + "/get-object?id=" + id.String() + "&auto-release=" + autoReleaseStr)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		byt, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		m := &generalResponse{}
		err = json.Unmarshal(byt, m)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf(m.Message)
	}
	return resp.Body, nil
}

func (c *Client) ReleaseObject(id uuid.UUID) error {
	resp, err := http.Get("http://" + c.addr + "/release-object?id=" + id.String())
	if err != nil {
		return err
	}
	byt, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	m := &generalResponse{}
	err = json.Unmarshal(byt, m)
	if err != nil {
		return err
	}
	if m.Status != "success" {
		return fmt.Errorf(m.Message)
	}
	return nil
}

func (c *Client) SetExpire(id uuid.UUID, expire time.Duration) error {
	durationStr := strconv.Itoa(int(expire / time.Second))
	resp, err := http.Get("http://" + c.addr + "/set-expire?id=" + id.String() + "&expire=" + durationStr)
	if err != nil {
		return err
	}
	byt, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	m := &generalResponse{}
	err = json.Unmarshal(byt, m)
	if err != nil {
		return err
	}
	if m.Status != "success" {
		return fmt.Errorf(m.Message)
	}
	return nil
}

func NewClient(addr string) *Client {
	return &Client{
		addr: addr,
	}
}
