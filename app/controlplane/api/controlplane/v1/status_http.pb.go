// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.3
// - protoc             (unknown)
// source: controlplane/v1/status.proto

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

const OperationStatusServiceInfoz = "/controlplane.v1.StatusService/Infoz"
const OperationStatusServiceStatusz = "/controlplane.v1.StatusService/Statusz"

type StatusServiceHTTPServer interface {
	Infoz(context.Context, *InfozRequest) (*InfozResponse, error)
	Statusz(context.Context, *StatuszRequest) (*StatuszResponse, error)
}

func RegisterStatusServiceHTTPServer(s *http.Server, srv StatusServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/infoz", _StatusService_Infoz0_HTTP_Handler(srv))
	r.GET("/statusz", _StatusService_Statusz0_HTTP_Handler(srv))
}

func _StatusService_Infoz0_HTTP_Handler(srv StatusServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in InfozRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStatusServiceInfoz)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Infoz(ctx, req.(*InfozRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*InfozResponse)
		return ctx.Result(200, reply)
	}
}

func _StatusService_Statusz0_HTTP_Handler(srv StatusServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StatuszRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStatusServiceStatusz)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Statusz(ctx, req.(*StatuszRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*StatuszResponse)
		return ctx.Result(200, reply)
	}
}

type StatusServiceHTTPClient interface {
	Infoz(ctx context.Context, req *InfozRequest, opts ...http.CallOption) (rsp *InfozResponse, err error)
	Statusz(ctx context.Context, req *StatuszRequest, opts ...http.CallOption) (rsp *StatuszResponse, err error)
}

type StatusServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewStatusServiceHTTPClient(client *http.Client) StatusServiceHTTPClient {
	return &StatusServiceHTTPClientImpl{client}
}

func (c *StatusServiceHTTPClientImpl) Infoz(ctx context.Context, in *InfozRequest, opts ...http.CallOption) (*InfozResponse, error) {
	var out InfozResponse
	pattern := "/infoz"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationStatusServiceInfoz))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StatusServiceHTTPClientImpl) Statusz(ctx context.Context, in *StatuszRequest, opts ...http.CallOption) (*StatuszResponse, error) {
	var out StatuszResponse
	pattern := "/statusz"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationStatusServiceStatusz))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
