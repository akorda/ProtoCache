package server

import (
	"context"
	"errors"
	"net"

	"github.com/akorda/protocache/caching"
	"github.com/akorda/protocache/proto"
	"google.golang.org/grpc"
)

type ProtoCacheOptions struct {
	ListenAddress string
}

type ProtoCacheServer struct {
	listenAddress string
	cache         caching.DistributedCache
	grpcServer    *grpc.Server
	proto.UnimplementedProtoCacheServer
}

func NewProtoCacheServer(cache caching.DistributedCache, options ProtoCacheOptions) (*ProtoCacheServer, error) {
	if cache == nil {
		return nil, errors.New("svc cannot be nil")
	}

	if len(options.ListenAddress) == 0 {
		options.ListenAddress = ":4000"
	}
	return &ProtoCacheServer{
		listenAddress: options.ListenAddress,
		cache:         cache,
	}, nil
}

func (s *ProtoCacheServer) Start() error {
	ln, err := net.Listen("tcp", s.listenAddress)
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	proto.RegisterProtoCacheServer(server, s)
	s.grpcServer = server

	return server.Serve(ln)
}

func (s *ProtoCacheServer) Stop() error {
	// TODO: check nil s.grpcServer

	s.grpcServer.GracefulStop()
	return nil
}

func (c *ProtoCacheServer) GetCacheItem(ctx context.Context, req *proto.GetCacheItemRequest) (*proto.GetCacheItemResponse, error) {
	value, err := c.cache.Get(req.Key)
	if err != nil {
		return nil, err
	}

	resp := &proto.GetCacheItemResponse{
		Value: value,
	}

	return resp, nil
}

func (c *ProtoCacheServer) SetCacheItem(ctx context.Context, req *proto.SetCacheItemRequest) (*proto.SetCacheItemResponse, error) {
	err := c.cache.Set(req.Key, req.Value)
	if err != nil {
		return nil, err
	}

	resp := &proto.SetCacheItemResponse{}

	return resp, nil
}

func (c *ProtoCacheServer) RemoveCacheItem(ctx context.Context, req *proto.RemoveCacheItemRequest) (*proto.RemoveCacheItemResponse, error) {
	err := c.cache.Remove(req.Key)
	if err != nil {
		return nil, err
	}

	resp := &proto.RemoveCacheItemResponse{}

	return resp, nil
}
