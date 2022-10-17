// Package server implements the objdeliv server
package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/lcpu-club/objdeliv/storage"
	"github.com/satori/uuid"
)

type Server struct {
	listen        string
	storageDriver storage.Driver
	defaultExpire int
}

func (s *Server) Serve() error {
	if s.storageDriver == nil {
		return fmt.Errorf("storageDriver should be initialized")
	}
	return http.ListenAndServe(s.listen, s)
}

func (s *Server) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println("request", req.Method, req.RequestURI, "remote", req.RemoteAddr)
	if strings.ToLower(req.Method) == "connect" && req.URL.Path == "/new-object" {
		s.serveConnect(resp, req)
		return
	}
	if strings.ToLower(req.Method) == "get" && req.URL.Path == "/get-object" {
		s.serveGet(resp, req)
		return
	}
	if req.URL.Path == "/release-object" {
		s.serveRelease(resp, req)
		return
	}
	if req.URL.Path == "/set-expire" {
		s.serveSetExpire(resp, req)
		return
	}
	resp.WriteHeader(404)
	resp.Write([]byte("404 Not Found\r\n\r\nobjdeliv server\r\n"))
}

func (s *Server) respError(resp http.ResponseWriter, msg string) {
	resp.WriteHeader(200)
	resp.Write([]byte("{\"status\":\"error\",\"message\":\"" + msg + "\"}"))
}

func (s *Server) respSuccess(resp http.ResponseWriter, msg string) {
	resp.WriteHeader(200)
	resp.Write([]byte("{\"status\":\"success\",\"message\":\"" + msg + "\"}"))
}

func (s *Server) serveSetExpire(resp http.ResponseWriter, req *http.Request) {
	var id uuid.UUID
	if !req.URL.Query().Has("id") {
		s.respError(resp, "`id` is needed for this request")
		return
	}
	id, err := uuid.FromString(req.URL.Query().Get("id"))
	if err != nil {
		s.respError(resp, "Invalid ID given")
		return
	}
	ok, err := s.storageDriver.IsExist(id)
	if !ok || err != nil {
		s.respError(resp, "No such ID")
		return
	}
	if !req.URL.Query().Has("expire") {
		s.respError(resp, "`expire` is needed for this request")
	}
	exp := req.URL.Query().Get("expire")
	e, err := strconv.Atoi(exp)
	if err != nil {
		s.respError(resp, "Invalid expire duration given")
		return
	}
	expire := time.Duration(e) * time.Second
	s.storageDriver.SetExpire(id, expire)
	s.respSuccess(resp, "Set expire time for "+id.String())
}

func (s *Server) serveRelease(resp http.ResponseWriter, req *http.Request) {
	var id uuid.UUID
	if !req.URL.Query().Has("id") {
		s.respError(resp, "ID needed for this request")
		return
	}
	id, err := uuid.FromString(req.URL.Query().Get("id"))
	if err != nil {
		s.respError(resp, "Invalid ID given")
		return
	}
	ok, err := s.storageDriver.IsExist(id)
	if !ok || err != nil {
		s.respError(resp, "No such ID")
		return
	}
	err = s.storageDriver.ReleaseObject(id)
	if err != nil {
		s.respError(resp, "Cannot release object "+id.String())
		return
	}
	s.respSuccess(resp, "Object "+id.String()+" released")
}

func (s *Server) serveGet(resp http.ResponseWriter, req *http.Request) {
	var id uuid.UUID
	if !req.URL.Query().Has("id") {
		s.respError(resp, "ID needed for this request")
		return
	}
	id, err := uuid.FromString(req.URL.Query().Get("id"))
	if err != nil {
		s.respError(resp, "Invalid ID given")
		return
	}
	ok, err := s.storageDriver.IsExist(id)
	if !ok || err != nil {
		s.respError(resp, "No such ID")
		return
	}
	autoRelease := req.URL.Query().Get("auto-release") == "true"
	// resp.Header().Add("Content-Length", )
	// TODO: add Content-Length
	if req.URL.Query().Has("content-type") {
		resp.Header().Add("Content-Type", req.URL.Query().Get("content-type"))
	}
	if req.URL.Query().Has("content-disposition") {
		resp.Header().Add("Content-Disposition", req.URL.Query().Get("content-disposition"))
	}
	if req.URL.Query().Has("pragma") {
		resp.Header().Add("Pragma", req.URL.Query().Get("pragma"))
	}
	if req.URL.Query().Has("cache-control") {
		resp.Header().Add("Cache-Control", req.URL.Query().Get("cache-control"))
	}
	if req.URL.Query().Has("expires") {
		resp.Header().Add("Expires", req.URL.Query().Get("expires"))
	}
	resp.WriteHeader(200)
	r, err := s.storageDriver.GetObject(id)
	if err != nil {
		s.respError(resp, "Cannot get object from storage")
		log.Println(err)
		return
	}
	defer r.Close()
	io.Copy(resp, r)
	if autoRelease {
		err = s.storageDriver.ReleaseObject(id)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func (s *Server) serveConnect(resp http.ResponseWriter, req *http.Request) {
	var id uuid.UUID
	var expire time.Duration = time.Duration(s.defaultExpire) * time.Second
	if req.URL.Query().Has("id") {
		var err error
		id, err = uuid.FromString(req.URL.Query().Get("id"))
		if err != nil {
			resp.WriteHeader(400)
			resp.Write([]byte("400 Bad Request\r\n\r\nInvalid ID given.\r\n"))
			return
		}
		ok, err := s.storageDriver.IsExist(id)
		if ok || err != nil {
			resp.WriteHeader(400)
			resp.Write([]byte("400 Bad Request\r\n\r\nID given already exists.\r\n"))
			return
		}
	} else {
		i := 1
	CREATE_UUID:
		id = uuid.NewV4()
		ok, err := s.storageDriver.IsExist(id)
		if err != nil {
			resp.WriteHeader(500)
			resp.Write([]byte("500 Internal Server Error\r\n\r\nAn internal error occurred.\r\n"))
			log.Println("[Error] StorageDriver returns error " + err.Error())
			return
		}
		if ok {
			i++
			if i > 10 {
				log.Println("The driver keeps telling us the key exists. It may be broken. Stops trying.")
				resp.WriteHeader(500)
				resp.Write([]byte("500 Internal Server Error\r\n\r\nAn internal error occurred.\r\n"))
				return
			}
			goto CREATE_UUID
		}
	}
	if req.URL.Query().Has("expire") {
		exp := req.URL.Query().Get("expire")
		e, err := strconv.Atoi(exp)
		if err != nil {
			resp.WriteHeader(400)
			resp.Write([]byte("400 Bad Request\r\n\r\nInvalid expire duration.\r\n"))
			return
		}
		expire = time.Duration(e) * time.Second
	}
	hijacker, ok := resp.(http.Hijacker)
	if !ok {
		resp.WriteHeader(400)
		resp.Write([]byte("400 Bad Request\r\n\r\nThis type of connection cannot be hijacked.\r\n"))
		return
	}
	conn, buf, err := hijacker.Hijack()
	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte("500 Internal Server Error\r\n\r\nAn error occurred when establishing connection.\r\n"))
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
		if expire > 0 {
			s.storageDriver.SetExpire(id, expire)
		}
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

func New(listen string, storageDriver storage.Driver, defaultExpire int) *Server {
	return &Server{
		listen:        listen,
		storageDriver: storageDriver,
		defaultExpire: defaultExpire,
	}
}
