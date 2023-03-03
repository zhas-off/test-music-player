// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: playlist.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PlaylistClient is the client API for Playlist service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlaylistClient interface {
	AddSong(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*PlaylistResponse, error)
	DeleteSong(ctx context.Context, in *DeleteSongRequest, opts ...grpc.CallOption) (*PlaylistResponse, error)
	PlaySong(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SongStats, error)
	PauseSong(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SongStats, error)
	Next(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SongStats, error)
	Prev(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SongStats, error)
	Status(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SongStats, error)
}

type playlistClient struct {
	cc grpc.ClientConnInterface
}

func NewPlaylistClient(cc grpc.ClientConnInterface) PlaylistClient {
	return &playlistClient{cc}
}

func (c *playlistClient) AddSong(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*PlaylistResponse, error) {
	out := new(PlaylistResponse)
	err := c.cc.Invoke(ctx, "/pb.Playlist/AddSong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistClient) DeleteSong(ctx context.Context, in *DeleteSongRequest, opts ...grpc.CallOption) (*PlaylistResponse, error) {
	out := new(PlaylistResponse)
	err := c.cc.Invoke(ctx, "/pb.Playlist/DeleteSong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistClient) PlaySong(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SongStats, error) {
	out := new(SongStats)
	err := c.cc.Invoke(ctx, "/pb.Playlist/PlaySong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistClient) PauseSong(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SongStats, error) {
	out := new(SongStats)
	err := c.cc.Invoke(ctx, "/pb.Playlist/PauseSong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistClient) Next(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SongStats, error) {
	out := new(SongStats)
	err := c.cc.Invoke(ctx, "/pb.Playlist/Next", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistClient) Prev(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SongStats, error) {
	out := new(SongStats)
	err := c.cc.Invoke(ctx, "/pb.Playlist/Prev", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistClient) Status(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SongStats, error) {
	out := new(SongStats)
	err := c.cc.Invoke(ctx, "/pb.Playlist/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlaylistServer is the server API for Playlist service.
// All implementations must embed UnimplementedPlaylistServer
// for forward compatibility
type PlaylistServer interface {
	AddSong(context.Context, *AddRequest) (*PlaylistResponse, error)
	DeleteSong(context.Context, *DeleteSongRequest) (*PlaylistResponse, error)
	PlaySong(context.Context, *Empty) (*SongStats, error)
	PauseSong(context.Context, *Empty) (*SongStats, error)
	Next(context.Context, *Empty) (*SongStats, error)
	Prev(context.Context, *Empty) (*SongStats, error)
	Status(context.Context, *Empty) (*SongStats, error)
	mustEmbedUnimplementedPlaylistServer()
}

// UnimplementedPlaylistServer must be embedded to have forward compatible implementations.
type UnimplementedPlaylistServer struct {
}

func (UnimplementedPlaylistServer) AddSong(context.Context, *AddRequest) (*PlaylistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSong not implemented")
}
func (UnimplementedPlaylistServer) DeleteSong(context.Context, *DeleteSongRequest) (*PlaylistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSong not implemented")
}
func (UnimplementedPlaylistServer) PlaySong(context.Context, *Empty) (*SongStats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaySong not implemented")
}
func (UnimplementedPlaylistServer) PauseSong(context.Context, *Empty) (*SongStats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PauseSong not implemented")
}
func (UnimplementedPlaylistServer) Next(context.Context, *Empty) (*SongStats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Next not implemented")
}
func (UnimplementedPlaylistServer) Prev(context.Context, *Empty) (*SongStats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prev not implemented")
}
func (UnimplementedPlaylistServer) Status(context.Context, *Empty) (*SongStats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedPlaylistServer) mustEmbedUnimplementedPlaylistServer() {}

// UnsafePlaylistServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlaylistServer will
// result in compilation errors.
type UnsafePlaylistServer interface {
	mustEmbedUnimplementedPlaylistServer()
}

func RegisterPlaylistServer(s grpc.ServiceRegistrar, srv PlaylistServer) {
	s.RegisterService(&Playlist_ServiceDesc, srv)
}

func _Playlist_AddSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServer).AddSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Playlist/AddSong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServer).AddSong(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playlist_DeleteSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServer).DeleteSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Playlist/DeleteSong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServer).DeleteSong(ctx, req.(*DeleteSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playlist_PlaySong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServer).PlaySong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Playlist/PlaySong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServer).PlaySong(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playlist_PauseSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServer).PauseSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Playlist/PauseSong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServer).PauseSong(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playlist_Next_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServer).Next(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Playlist/Next",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServer).Next(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playlist_Prev_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServer).Prev(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Playlist/Prev",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServer).Prev(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playlist_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Playlist/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServer).Status(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Playlist_ServiceDesc is the grpc.ServiceDesc for Playlist service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Playlist_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Playlist",
	HandlerType: (*PlaylistServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddSong",
			Handler:    _Playlist_AddSong_Handler,
		},
		{
			MethodName: "DeleteSong",
			Handler:    _Playlist_DeleteSong_Handler,
		},
		{
			MethodName: "PlaySong",
			Handler:    _Playlist_PlaySong_Handler,
		},
		{
			MethodName: "PauseSong",
			Handler:    _Playlist_PauseSong_Handler,
		},
		{
			MethodName: "Next",
			Handler:    _Playlist_Next_Handler,
		},
		{
			MethodName: "Prev",
			Handler:    _Playlist_Prev_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Playlist_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "playlist.proto",
}
