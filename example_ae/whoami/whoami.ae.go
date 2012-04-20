// +build appengine

// Code generated by protoc-gen-go from "whoami/whoami.proto"
// DO NOT EDIT!

package whoami

import "math"

import "net/url"
import "net/http"
import "github.com/kylelemons/go-rpcgen/webrpc"

// Reference proto and math imports to suppress error if they are not otherwise used.
var _ = math.Inf

type Empty struct {
	XXX_unrecognized []byte `json:"-"`
}

func (this *Empty) Reset()         { *this = Empty{} }

type YouAre struct {
	IpAddr           *string `protobuf:"bytes,1,req,name=ip_addr" json:"ip_addr,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *YouAre) Reset()         { *this = YouAre{} }

func init() {
}

// WhoamiService is an interface satisfied by the generated client and
// which must be implemented by the object wrapped by the server.
type WhoamiService interface {
	Whoami(in *Empty, out *YouAre) error
}

// WhoamiServiceWeb is the web-based RPC version of the interface which
// must be implemented by the object wrapped by the webrpc server.
type WhoamiServiceWeb interface {
	Whoami(r *http.Request, in *Empty, out *YouAre) error
}

// internal wrapper for type-safe webrpc calling
type rpcWhoamiServiceWebClient struct {
	remote   *url.URL
	protocol webrpc.Protocol
}

func (this rpcWhoamiServiceWebClient) Whoami(in *Empty, out *YouAre) error {
	return webrpc.Post(this.protocol, this.remote, "/WhoamiService/Whoami", in, out)
}

// Register a WhoamiServiceWeb implementation with the given webrpc ServeMux.
// If mux is nil, the default webrpc.ServeMux is used.
func RegisterWhoamiServiceWeb(this WhoamiServiceWeb, mux webrpc.ServeMux) error {
	if mux == nil {
		mux = webrpc.DefaultServeMux
	}
	if err := mux.Handle("/WhoamiService/Whoami", func(c *webrpc.Call) error {
		in, out := new(Empty), new(YouAre)
		if err := c.ReadRequest(in); err != nil {
			return err
		}
		if err := this.Whoami(c.Request, in, out); err != nil {
			return err
		}
		return c.WriteResponse(out)
	}); err != nil {
		return err
	}
	return nil
}

// NewWhoamiServiceWebClient returns a webrpc wrapper for calling the methods of WhoamiService
// remotely via the web.  The remote URL is the base URL of the webrpc server.
func NewWhoamiServiceWebClient(pro webrpc.Protocol, remote *url.URL) WhoamiService {
	return rpcWhoamiServiceWebClient{remote, pro}
}
