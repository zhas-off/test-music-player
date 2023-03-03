package grpc

import (
	db "github.com/zhas-off/test-music-player/internal/repository"
	"github.com/zhas-off/test-music-player/internal/service"
	pb "github.com/zhas-off/test-music-player/proto/pb"
)

type GrpcService struct {
	pb.PlaylistServer
	Pl *service.Playlist
	Db *db.PlaylistDatabase
}
