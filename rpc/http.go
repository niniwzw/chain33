// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/rpc/jsonrpc"
	"strings"

	"github.com/rs/cors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pr "google.golang.org/grpc/peer"
)

// adapt HTTP connection to ReadWriteCloser
type HTTPConn struct {
	r   *http.Request
	in  io.Reader
	out io.Writer
}

func (c *HTTPConn) Read(p []byte) (n int, err error) { return c.in.Read(p) }

func (c *HTTPConn) Write(d []byte) (n int, err error) { //添加支持gzip 发送

	if strings.Contains(c.r.Header.Get("Accept-Encoding"), "gzip") {
		gw := gzip.NewWriter(c.out)
		defer gw.Close()
		return gw.Write(d)
	}
	return c.out.Write(d)
}

func (c *HTTPConn) Close() error { return nil }

func (j *JSONRPCServer) Listen() (int, error) {
	listener, err := net.Listen("tcp", rpcCfg.JrpcBindAddr)
	if err != nil {
		return 0, err
	}
	j.l = listener
	co := cors.New(cors.Options{})

	// Insert the middleware
	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debug("JSONRPCServer", "RemoteAddr", r.RemoteAddr)
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			writeError(w, r, 0, fmt.Sprintf(`The %s Address is not authorized!`, ip))
			return
		}

		if !checkIpWhitelist(ip) {
			writeError(w, r, 0, fmt.Sprintf(`The %s Address is not authorized!`, ip))
			return
		}
		if r.URL.Path == "/" {
			data, err := ioutil.ReadAll(r.Body)
			if err != nil {
				writeError(w, r, 0, "Can't get request body!")
				return
			}
			//格式做一个检查
			client, err := parseJsonRpcParams(data)
			errstr := "nil"
			if err != nil {
				errstr = err.Error()
			}
			log.Debug("JSONRPCServer", "request", string(data), "err", errstr)
			if err != nil {
				writeError(w, r, 0, fmt.Sprintf(`parse request err %s`, err.Error()))
				return
			}
			//Release local request
			ipaddr := net.ParseIP(ip)
			if !ipaddr.IsLoopback() {
				funcName := strings.Split(client.Method, ".")[len(strings.Split(client.Method, "."))-1]
				if checkJrpcFuncBlacklist(funcName) || !checkJrpcFuncWhitelist(funcName) {
					writeError(w, r, client.Id, fmt.Sprintf(`The %s method is not authorized!`, funcName))
					return
				}
			}
			serverCodec := jsonrpc.NewServerCodec(&HTTPConn{in: ioutil.NopCloser(bytes.NewReader(data)), out: w, r: r})
			w.Header().Set("Content-type", "application/json")
			if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
				w.Header().Set("Content-Encoding", "gzip")
			}
			w.WriteHeader(200)
			err = j.s.ServeRequest(serverCodec)
			if err != nil {
				log.Debug("Error while serving JSON request: %v", err)
				return
			}
		}
	})

	handler = co.Handler(handler)
	go http.Serve(listener, handler)
	return listener.Addr().(*net.TCPAddr).Port, nil
}

type serverResponse struct {
	Id     uint64      `json:"id"`
	Result interface{} `json:"result"`
	Error  interface{} `json:"error"`
}

func writeError(w http.ResponseWriter, r *http.Request, id uint64, errstr string) {
	w.Header().Set("Content-type", "application/json")
	//错误的请求也返回 200
	w.WriteHeader(200)
	resp, err := json.Marshal(&serverResponse{id, nil, errstr})
	if err != nil {
		log.Debug("json marshal error, nerver happen")
		return
	}
	w.Write(resp)
}

func (g *Grpcserver) Listen() (int, error) {
	listener, err := net.Listen("tcp", rpcCfg.GrpcBindAddr)
	if err != nil {
		return 0, err
	}
	g.l = listener
	go g.s.Serve(listener)
	return listener.Addr().(*net.TCPAddr).Port, nil
}

func isLoopBackAddr(addr net.Addr) bool {
	if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.IsLoopback() {
		return true
	}
	return false
}

func auth(ctx context.Context, info *grpc.UnaryServerInfo) error {
	getctx, ok := pr.FromContext(ctx)
	if ok {
		if isLoopBackAddr(getctx.Addr) {
			return nil
		}
		//remoteaddr := strings.Split(getctx.Addr.String(), ":")[0]
		ip, _, err := net.SplitHostPort(getctx.Addr.String())
		if err != nil {
			return fmt.Errorf("The %s Address is not authorized!", ip)
		}

		if !checkIpWhitelist(ip) {
			return fmt.Errorf("The %s Address is not authorized!", ip)
		}

		funcName := strings.Split(info.FullMethod, "/")[len(strings.Split(info.FullMethod, "/"))-1]
		if checkGrpcFuncBlacklist(funcName) || !checkGrpcFuncWhitelist(funcName) {
			return fmt.Errorf("The %s method is not authorized!", funcName)
		}
		return nil
	}
	return fmt.Errorf("Can't get remote ip!")
}

type clientRequest struct {
	Method string         `json:"method"`
	Params [1]interface{} `json:"params"`
	Id     uint64         `json:"id"`
}

func parseJsonRpcParams(data []byte) (*clientRequest, error) {
	var req clientRequest
	err := json.Unmarshal(data, &req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}
