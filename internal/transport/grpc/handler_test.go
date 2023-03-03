package grpc

import (
	"context"
	"io"
	"testing"

	"github.com/rs/zerolog"
	"github.com/zhas-off/test-music-player/internal/models"
	"github.com/zhas-off/test-music-player/internal/service"
	"github.com/zhas-off/test-music-player/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGrpcEndpoints_PlaySong_PauseSong_StatusSong(t *testing.T) {
	logger := zerolog.New(io.Discard)
	pl := service.Init()
	pl.Logger = &logger
	go pl.Run()
	grpc := GrpcService{
		Pl: pl,
	}
	pl.AddNewSong(models.Song{
		Name:     "random",
		Duration: 25,
	})
	t.Run("play test", func(t *testing.T) {
		resp, err := grpc.PlaySong(context.Background(), &pb.Empty{})
		if err != status.Errorf(codes.OK, "OK") {
			t.Error(err)
			t.Fail()
		}
		if resp.Name != "random" || resp.Time != "00:00:25" {
			t.Error(err)
			t.Fail()
		}
	})

	t.Run("pause test", func(t *testing.T) {
		resp, err := grpc.PauseSong(context.Background(), &pb.Empty{})
		if err != status.Errorf(codes.OK, "OK") {
			t.Error(err)
			t.Fail()
		}
		if resp.Name != "random" || resp.Time != "00:00:25" {
			t.Error(err)
			t.Fail()
		}
	})

	t.Run("status test", func(t *testing.T) {
		resp, err := grpc.Status(context.Background(), &pb.Empty{})
		if err != status.Errorf(codes.OK, "OK") {
			t.Error(err)
			t.Fail()
		}
		if resp.Name != "random" || resp.Time != "00:00:25" {
			t.Error(err)
			t.Fail()
		}
	})
}

func TestGrpcEndpoints_Next_Prev(t *testing.T) {
	logger := zerolog.New(io.Discard)
	pl := service.Init()
	pl.Logger = &logger
	go pl.Run()
	grpc := GrpcService{
		Pl: pl,
	}
	pl.AddNewSong(models.Song{
		Name:     "random",
		Duration: 25,
	})
	pl.AddNewSong(models.Song{
		Name:     "random2",
		Duration: 25,
	})

	t.Run("play test", func(t *testing.T) {
		resp, err := grpc.PlaySong(context.Background(), &pb.Empty{})
		if err != status.Errorf(codes.OK, "OK") {
			t.Error(err)
			t.Fail()
		}

		resp, err = grpc.PauseSong(context.Background(), &pb.Empty{})
		if err != status.Errorf(codes.OK, "OK") {
			t.Error(err)
			t.Fail()
		}
		if resp.Name != "random" || resp.Time != "00:00:25" {
			t.Error(err)
			t.Fail()
		}
	})
	t.Run("next test #1", func(t *testing.T) {
		resp, err := grpc.Next(context.Background(), &pb.Empty{})
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		if resp.Name != "random2" || resp.Time != "00:00:25" {
			t.Error(err)
			t.Fail()
		}
	})
	t.Run("next test #2", func(t *testing.T) {
		_, err := grpc.Next(context.Background(), &pb.Empty{})
		statusString := "The next song does not exist."
		if err.Error() != status.Errorf(codes.NotFound, statusString).Error() {
			t.Error(err)
			t.Fail()
		}
	})

	t.Run("prev test #1", func(t *testing.T) {
		resp, err := grpc.Prev(context.Background(), &pb.Empty{})
		if err != status.Errorf(codes.OK, "OK") {
			t.Error(err)
			t.Fail()
		}
		if resp.Name != "random" || resp.Time != "00:00:25" {
			t.Error(err)
			t.Fail()
		}
	})

	t.Run("next test #2", func(t *testing.T) {
		_, err := grpc.Prev(context.Background(), &pb.Empty{})
		statusString := "You are at the beginning of the playlist."
		if err.Error() != status.Errorf(codes.NotFound, statusString).Error() {
			t.Error(err)
			t.Fail()
		}
	})
}
