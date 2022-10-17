// Package server implements the objdeliv server
package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/lcpu-club/objdeliv/storage"
	"github.com/satori/uuid"
)

type Server struct {
	listen        string
	storageDriver storage.Driver
}

func (s *Server) Serve() error {
	if s.storageDriver == nil {
		return fmt.Errorf("storageDriver should be initialized")
	}
	return http.ListenAndServe(s.listen, s)
}

func (s *Server) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println("request", req.Method, req.RequestURI, "remote", req.RemoteAddr)
	if strings.ToLower(req.Method) == "connect" && req.URL.Path == "/connect" {
		s.ServeConnect(resp, req)
		return
	}
	resp.Write([]byte("Hello World\r\n"))
}

func (s *Server) ServeConnect(resp http.ResponseWriter, req *http.Request) {
	var id uuid.UUID
	if req.URL.Query().Has("id") {
		var err error
		id, err = uuid.FromString(req.URL.Query().Get("id"))
		if err != nil {
			resp.WriteHeader(400)
			resp.Write([]byte("400 Bad Request\r\n\r\nInvalid ID given."))
			return
		}
	} else {
		i := 1
	CREATE_UUID:
		id = uuid.NewV4()
		ok, err := s.storageDriver.IsExist(id)
		if err != nil {
			resp.WriteHeader(500)
			resp.Write([]byte("500 Internal Server Error\r\n\r\nAn internal error occurred."))
			log.Println("[Error] StorageDriver returns error " + err.Error())
			return
		}
		if ok {
			i++
			if i > 10 {
				log.Println("The driver keeps telling us the key exists. It may be broken. Stops trying.")
				resp.WriteHeader(500)
				resp.Write([]byte("500 Internal Server Error\r\n\r\nAn internal error occurred."))
				return
			}
			goto CREATE_UUID
		}
	}
	hijacker, ok := resp.(http.Hijacker)
	if !ok {
		resp.WriteHeader(400)
		resp.Write([]byte("400 Bad Request\r\n\r\nThis type of connection cannot be hijacked."))
		return
	}
	conn, buf, err := hijacker.Hijack()
	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte("500 Internal Server Error\r\n\r\nAn error occurred when establishing connection."))
		return
	}
	defer conn.Close()
	_, err = conn.Write([]byte("{\"status\":\"success\",\"id\":\"" + id.String() + "\"}\r\n"))
	if err == net.ErrClosed {
		log.Println("Client unexpectedly unconnected")
		return
	}
	if err != nil {
		log.Println(err)
		return
	}
	obj, err := s.storageDriver.NewObject(id)
	if err != nil {
		log.Println(err)
		conn.Write([]byte("{\"status\":\"error\",\"message\":\"object creation failed\"}\r\n"))
		return
	}
	defer func() {
		err := obj.Close()
		if err != nil {
			log.Println(err)
		}
		log.Println("Object", id.String(), "written")
	}()
	bufContent, err := io.ReadAll(buf)
	if err != nil {
		log.Println(err)
		conn.Write([]byte("{\"status\":\"error\",\"message\":\"buffer flush failed\"}\r\n"))
		return
	}
	_, err = obj.Write(bufContent)
	if err != nil {
		log.Println(err)
		conn.Write([]byte("{\"status\":\"error\",\"message\":\"write object failed\"}\r\n"))
		return
	}
	_, err = io.Copy(obj, conn)
	if err != nil {
		log.Println(err)
		conn.Write([]byte("{\"status\":\"error\",\"message\":\"write object failed\"}\r\n"))
		return
	}
}

func New(listen string, storageDriver storage.Driver) *Server {
	return &Server{
		listen:        listen,
		storageDriver: storageDriver,
	}
}
