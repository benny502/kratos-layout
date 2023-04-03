// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.1
// - protoc             v3.21.12
// source: login/v1/login.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationLoginCheck = "/login.v1.Login/Check"
const OperationLoginLogout = "/login.v1.Login/Logout"
const OperationLoginVcode = "/login.v1.Login/Vcode"

type LoginHTTPServer interface {
	Check(context.Context, *CheckRequest) (*CheckReply, error)
	Logout(context.Context, *LogoutRequest) (*LogoutReply, error)
	Vcode(context.Context, *VcodeRequest) (*VcodeReply, error)
}

func RegisterLoginHTTPServer(s *http.Server, srv LoginHTTPServer) {
	r := s.Route("/")
	r.GET("/login/vcode", _Login_Vcode0_HTTP_Handler(srv))
	r.POST("/login/check", _Login_Check0_HTTP_Handler(srv))
	r.POST("/login/logout", _Login_Logout0_HTTP_Handler(srv))
}

func _Login_Vcode0_HTTP_Handler(srv LoginHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in VcodeRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLoginVcode)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Vcode(ctx, req.(*VcodeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*VcodeReply)
		return ctx.Result(200, reply)
	}
}

func _Login_Check0_HTTP_Handler(srv LoginHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CheckRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLoginCheck)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Check(ctx, req.(*CheckRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CheckReply)
		return ctx.Result(200, reply)
	}
}

func _Login_Logout0_HTTP_Handler(srv LoginHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LogoutRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLoginLogout)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Logout(ctx, req.(*LogoutRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LogoutReply)
		return ctx.Result(200, reply)
	}
}

type LoginHTTPClient interface {
	Check(ctx context.Context, req *CheckRequest, opts ...http.CallOption) (rsp *CheckReply, err error)
	Logout(ctx context.Context, req *LogoutRequest, opts ...http.CallOption) (rsp *LogoutReply, err error)
	Vcode(ctx context.Context, req *VcodeRequest, opts ...http.CallOption) (rsp *VcodeReply, err error)
}

type LoginHTTPClientImpl struct {
	cc *http.Client
}

func NewLoginHTTPClient(client *http.Client) LoginHTTPClient {
	return &LoginHTTPClientImpl{client}
}

func (c *LoginHTTPClientImpl) Check(ctx context.Context, in *CheckRequest, opts ...http.CallOption) (*CheckReply, error) {
	var out CheckReply
	pattern := "/login/check"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLoginCheck))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LoginHTTPClientImpl) Logout(ctx context.Context, in *LogoutRequest, opts ...http.CallOption) (*LogoutReply, error) {
	var out LogoutReply
	pattern := "/login/logout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLoginLogout))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LoginHTTPClientImpl) Vcode(ctx context.Context, in *VcodeRequest, opts ...http.CallOption) (*VcodeReply, error) {
	var out VcodeReply
	pattern := "/login/vcode"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationLoginVcode))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}