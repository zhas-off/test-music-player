package app

import (
	"flag"
	"net"

	"github.com/rs/zerolog"
	config "github.com/zhas-off/test-music-player/internal/config"
	repo "github.com/zhas-off/test-music-player/internal/repository"
	"github.com/zhas-off/test-music-player/internal/service"
	transport "github.com/zhas-off/test-music-player/internal/transport/grpc"
	"github.com/zhas-off/test-music-player/proto/pb"
	"google.golang.org/grpc"
)

type PlaylistServer struct {
	Logger      *zerolog.Logger
	Config      *config.Config
	GrpcService *transport.GrpcService
}

func New(pl *service.Playlist, db *repo.PlaylistDatabase) *PlaylistServer {
	return &PlaylistServer{GrpcService: &transport.GrpcService{Pl: pl, Db: db}}
}

// Запуск программы
func (s PlaylistServer) Run() {

	grpcServerAddr := flag.String("gRPC server", s.Config.Addr, "gRPC server address")

	lis, err := net.Listen("tcp", *grpcServerAddr)
	if err != nil {
		s.Logger.Fatal().Err(err).Msg("failed to listen")
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPlaylistServer(grpcServer, s.GrpcService)

	grpcErr := grpcServer.Serve(lis)

	s.Logger.Info().Msg("gRPC server has started")

	if grpcErr != nil {
		s.Logger.Fatal().Err(grpcErr).Msg("failed on starting gRPC")
	}
}
