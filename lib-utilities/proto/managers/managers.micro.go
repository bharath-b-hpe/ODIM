// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/managers/managers.proto

package managers

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Managers service

type ManagersService interface {
	GetManagersCollection(ctx context.Context, in *ManagerRequest, opts ...client.CallOption) (*ManagerResponse, error)
	GetManager(ctx context.Context, in *ManagerRequest, opts ...client.CallOption) (*ManagerResponse, error)
	GetManagersResource(ctx context.Context, in *ManagerRequest, opts ...client.CallOption) (*ManagerResponse, error)
	VirtualMediaInsert(ctx context.Context, in *ManagerRequest, opts ...client.CallOption) (*ManagerResponse, error)
	VirtualMediaEject(ctx context.Context, in *ManagerRequest, opts ...client.CallOption) (*ManagerResponse, error)
}

type managersService struct {
	c    client.Client
	name string
}

func NewManagersService(name string, c client.Client) ManagersService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "managers"
	}
	return &managersService{
		c:    c,
		name: name,
	}
}

func (c *managersService) GetManagersCollection(ctx context.Context, in *ManagerRequest, opts ...client.CallOption) (*ManagerResponse, error) {
	req := c.c.NewRequest(c.name, "Managers.GetManagersCollection", in)
	out := new(ManagerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managersService) GetManager(ctx context.Context, in *ManagerRequest, opts ...client.CallOption) (*ManagerResponse, error) {
	req := c.c.NewRequest(c.name, "Managers.GetManager", in)
	out := new(ManagerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managersService) GetManagersResource(ctx context.Context, in *ManagerRequest, opts ...client.CallOption) (*ManagerResponse, error) {
	req := c.c.NewRequest(c.name, "Managers.GetManagersResource", in)
	out := new(ManagerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managersService) VirtualMediaInsert(ctx context.Context, in *ManagerRequest, opts ...client.CallOption) (*ManagerResponse, error) {
	req := c.c.NewRequest(c.name, "Managers.VirtualMediaInsert", in)
	out := new(ManagerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managersService) VirtualMediaEject(ctx context.Context, in *ManagerRequest, opts ...client.CallOption) (*ManagerResponse, error) {
	req := c.c.NewRequest(c.name, "Managers.VirtualMediaEject", in)
	out := new(ManagerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Managers service

type ManagersHandler interface {
	GetManagersCollection(context.Context, *ManagerRequest, *ManagerResponse) error
	GetManager(context.Context, *ManagerRequest, *ManagerResponse) error
	GetManagersResource(context.Context, *ManagerRequest, *ManagerResponse) error
	VirtualMediaInsert(context.Context, *ManagerRequest, *ManagerResponse) error
	VirtualMediaEject(context.Context, *ManagerRequest, *ManagerResponse) error
}

func RegisterManagersHandler(s server.Server, hdlr ManagersHandler, opts ...server.HandlerOption) error {
	type managers interface {
		GetManagersCollection(ctx context.Context, in *ManagerRequest, out *ManagerResponse) error
		GetManager(ctx context.Context, in *ManagerRequest, out *ManagerResponse) error
		GetManagersResource(ctx context.Context, in *ManagerRequest, out *ManagerResponse) error
		VirtualMediaInsert(ctx context.Context, in *ManagerRequest, out *ManagerResponse) error
		VirtualMediaEject(ctx context.Context, in *ManagerRequest, out *ManagerResponse) error
	}
	type Managers struct {
		managers
	}
	h := &managersHandler{hdlr}
	return s.Handle(s.NewHandler(&Managers{h}, opts...))
}

type managersHandler struct {
	ManagersHandler
}

func (h *managersHandler) GetManagersCollection(ctx context.Context, in *ManagerRequest, out *ManagerResponse) error {
	return h.ManagersHandler.GetManagersCollection(ctx, in, out)
}

func (h *managersHandler) GetManager(ctx context.Context, in *ManagerRequest, out *ManagerResponse) error {
	return h.ManagersHandler.GetManager(ctx, in, out)
}

func (h *managersHandler) GetManagersResource(ctx context.Context, in *ManagerRequest, out *ManagerResponse) error {
	return h.ManagersHandler.GetManagersResource(ctx, in, out)
}

func (h *managersHandler) VirtualMediaInsert(ctx context.Context, in *ManagerRequest, out *ManagerResponse) error {
	return h.ManagersHandler.VirtualMediaInsert(ctx, in, out)
}

func (h *managersHandler) VirtualMediaEject(ctx context.Context, in *ManagerRequest, out *ManagerResponse) error {
	return h.ManagersHandler.VirtualMediaEject(ctx, in, out)
}