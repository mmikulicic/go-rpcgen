// Code generated by protoc-gen-go.
// source: examples/echo/echoservice/echoservice.proto
// DO NOT EDIT!

package echoservice

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

import "net"
import "net/rpc"
import "github.com/kylelemons/go-rpcgen/codec"
import "net/url"
import "net/http"
import "github.com/kylelemons/go-rpcgen/webrpc"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Payload struct {
	Message          *string `protobuf:"bytes,1,req,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *Payload) Reset()         { *this = Payload{} }
func (this *Payload) String() string { return proto.CompactTextString(this) }
func (*Payload) ProtoMessage()       {}

func (this *Payload) GetMessage() string {
	if this != nil && this.Message != nil {
		return *this.Message
	}
	return ""
}

func init() {
}

// EchoService is an interface satisfied by the generated client and
// which must be implemented by the object wrapped by the server.
type EchoService interface {
	Echo(in *Payload, out *Payload) error
}

// internal wrapper for type-safe RPC calling
type rpcEchoServiceClient struct {
	*rpc.Client
}

func (this rpcEchoServiceClient) Echo(in *Payload, out *Payload) error {
	return this.Call("EchoService.Echo", in, out)
}

// NewEchoServiceClient returns an *rpc.Client wrapper for calling the methods of
// EchoService remotely.
func NewEchoServiceClient(conn net.Conn) EchoService {
	return rpcEchoServiceClient{rpc.NewClientWithCodec(codec.NewClientCodec(conn))}
}

// ServeEchoService serves the given EchoService backend implementation on conn.
func ServeEchoService(conn net.Conn, backend EchoService) error {
	srv := rpc.NewServer()
	if err := srv.RegisterName("EchoService", backend); err != nil {
		return err
	}
	srv.ServeCodec(codec.NewServerCodec(conn))
	return nil
}

// DialEchoService returns a EchoService for calling the EchoService servince at addr (TCP).
func DialEchoService(addr string) (EchoService, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return NewEchoServiceClient(conn), nil
}

// ListenAndServeEchoService serves the given EchoService backend implementation
// on all connections accepted as a result of listening on addr (TCP).
func ListenAndServeEchoService(addr string, backend EchoService) error {
	clients, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	srv := rpc.NewServer()
	if err := srv.RegisterName("EchoService", backend); err != nil {
		return err
	}
	for {
		conn, err := clients.Accept()
		if err != nil {
			return err
		}
		go srv.ServeCodec(codec.NewServerCodec(conn))
	}
	panic("unreachable")
}

// EchoServiceWeb is the web-based RPC version of the interface which
// must be implemented by the object wrapped by the webrpc server.
type EchoServiceWeb interface {
	Echo(r *http.Request, in *Payload, out *Payload) error
}

// internal wrapper for type-safe webrpc calling
type rpcEchoServiceWebClient struct {
	remote   *url.URL
	protocol webrpc.Protocol
}

func (this rpcEchoServiceWebClient) Echo(in *Payload, out *Payload) error {
	return webrpc.Post(this.protocol, this.remote, "/EchoService/Echo", in, out)
}

// Register a EchoServiceWeb implementation with the given webrpc ServeMux.
// If mux is nil, the default webrpc.ServeMux is used.
func RegisterEchoServiceWeb(this EchoServiceWeb, mux webrpc.ServeMux) error {
	if mux == nil {
		mux = webrpc.DefaultServeMux
	}
	if err := mux.Handle("/EchoService/Echo", func(c *webrpc.Call) error {
		in, out := new(Payload), new(Payload)
		if err := c.ReadRequest(in); err != nil {
			return err
		}
		if err := this.Echo(c.Request, in, out); err != nil {
			return err
		}
		return c.WriteResponse(out)
	}); err != nil {
		return err
	}
	return nil
}

// NewEchoServiceWebClient returns a webrpc wrapper for calling the methods of EchoService
// remotely via the web.  The remote URL is the base URL of the webrpc server.
func NewEchoServiceWebClient(pro webrpc.Protocol, remote *url.URL) EchoService {
	return rpcEchoServiceWebClient{remote, pro}
}
