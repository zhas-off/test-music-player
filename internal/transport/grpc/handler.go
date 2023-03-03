package grpc

import (
	"context"
	"errors"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/zhas-off/test-music-player/internal/models"
	"github.com/zhas-off/test-music-player/internal/utils"
	"github.com/zhas-off/test-music-player/proto/pb"
	"google.golang.org/grpc/codes"
	st "google.golang.org/grpc/status"
)

func (s *GrpcService) PlaySong(ctx context.Context, req *pb.Empty) (*pb.SongStats, error) {
	songStats := s.Pl.Play()
	timeString := utils.ConvertSongStatsToString(songStats)
	s.Pl.Logger.Info().Msg(fmt.Sprintf("playing %v at %s", songStats, timeString))
	return &pb.SongStats{
		Name:   songStats.Name,
		Time:   utils.ConvertFromSecondsToString(songStats.Duration),
		Status: fmt.Sprintf("%s plays at %s", songStats.Name, timeString),
	}, st.Errorf(codes.OK, "OK")
}

func (s *GrpcService) PauseSong(ctx context.Context, req *pb.Empty) (*pb.SongStats, error) {
	songStats := s.Pl.Pause()
	timeString := utils.ConvertSongStatsToString(songStats)
	s.Pl.Logger.Info().Msg(fmt.Sprintf("paused [%v] at %s", songStats, timeString))
	return &pb.SongStats{
		Name:   songStats.Name,
		Time:   utils.ConvertFromSecondsToString(songStats.Duration),
		Status: fmt.Sprintf("%s is paused at %s", songStats.Name, timeString),
	}, st.Errorf(codes.OK, "OK")
}

func (s *GrpcService) Next(ctx context.Context, req *pb.Empty) (*pb.SongStats, error) {
	songStats := s.Pl.Next()
	timeString := utils.ConvertSongStatsToString(songStats)
	s.Pl.Logger.Info().Msg(fmt.Sprintf("next song: [%v] at %s (exist: %v)", songStats, timeString, songStats.Exist))

	var status string
	if songStats.Exist {
		status = fmt.Sprintf("Switched to next song: %s", songStats.Name)
		return &pb.SongStats{
			Name:   songStats.Name,
			Time:   utils.ConvertFromSecondsToString(songStats.Duration),
			Status: status,
		}, st.Errorf(codes.OK, "OK")
	} else {
		status = "The next song does not exist."
		return &pb.SongStats{}, st.Errorf(codes.NotFound, status)
	}

}

func (s *GrpcService) Prev(ctx context.Context, req *pb.Empty) (*pb.SongStats, error) {
	songStats := s.Pl.Prev()
	timeString := utils.ConvertSongStatsToString(songStats)
	s.Pl.Logger.Info().Msg(fmt.Sprintf("prev song: [%v] at %s (exist: %v)", songStats, timeString, songStats.Exist))
	var status string
	if songStats.Exist {
		status = fmt.Sprintf("Switched to previous song: %s", songStats.Name)
		return &pb.SongStats{
			Name:   songStats.Name,
			Time:   utils.ConvertFromSecondsToString(songStats.Duration),
			Status: status,
		}, st.Errorf(codes.OK, "OK")
	} else {
		status = "You are at the beginning of the playlist."
		return &pb.SongStats{}, st.Errorf(codes.NotFound, status)

	}

}

func (s *GrpcService) Status(ctx context.Context, req *pb.Empty) (*pb.SongStats, error) {
	songStats := s.Pl.Status()
	timeString := utils.ConvertSongStatsToString(songStats)
	s.Pl.Logger.Info().Msg(fmt.Sprintf("status song: [%v] at %s(playing: %v)", songStats, timeString, songStats.Playing))
	var status string
	if songStats.Playing {
		status = fmt.Sprintf("Playback status: %s playing on %s", songStats.Name, timeString)
	} else {
		status = fmt.Sprintf("Playback status: %s paused on %s", songStats.Name, timeString)
	}
	return &pb.SongStats{
		Name:   songStats.Name,
		Time:   utils.ConvertFromSecondsToString(songStats.Duration),
		Status: status,
	}, st.Errorf(codes.OK, "OK")
}

func (s *GrpcService) AddSong(ctx context.Context, req *pb.AddRequest) (*pb.PlaylistResponse, error) {
	time, err := utils.ParseTimeToSeconds(req.Time)
	if err != nil {
		s.Pl.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("parse time to seconds error")
		return nil, st.Error(codes.FailedPrecondition, "incorrect duration format")
	}
	if req.Name == "" {
		s.Pl.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("empty name")
		return nil, st.Error(codes.FailedPrecondition, "empty name")
	}
	ok := s.Pl.AddNewSong(models.Song{Name: req.Name, Duration: time})
	if !ok {
		s.Pl.Logger.WithLevel(zerolog.WarnLevel).Err(errors.New("new song adding error")).Msg("song already exist")
		return nil, st.Error(codes.FailedPrecondition, "failed on adding, song already exist or incorrect data")
	}
	list, err := s.Pl.GetList()
	if err != nil {
		s.Pl.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("playlist getting error in AddSong")
		return nil, st.Error(codes.Internal, err.Error())
	}
	var res pb.PlaylistResponse
	for _, song := range list {
		dur := utils.ConvertFromSecondsToString(song.Duration)
		songRes := pb.Song{
			Name:     song.Name,
			Duration: dur,
		}
		res.Playlist = append(res.Playlist, &songRes)
	}
	err = s.Db.Add(req.Name, req.Time)
	if err != nil {
		s.Pl.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("adding to database error")
		return nil, st.Error(codes.Internal, err.Error())
	}
	s.Pl.Logger.Info().Msg(fmt.Sprintf("[%v] added into playlist", models.Song{Name: req.Name, Duration: time}))
	return &res, st.Error(codes.OK, "OK")
}

func (s *GrpcService) DeleteSong(ctx context.Context, req *pb.DeleteSongRequest) (*pb.PlaylistResponse, error) {
	var res pb.PlaylistResponse
	if req.Name == "" {
		s.Pl.Logger.WithLevel(zerolog.WarnLevel).Err(errors.New("empty name field")).Msg("song deleting error")
		return &res, st.Error(codes.FailedPrecondition, errors.New("empty name field").Error())
	}
	err := s.Pl.DeleteSong(req.Name)
	if err != nil {
		s.Pl.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("song deleting error")
		return &res, st.Error(codes.FailedPrecondition, err.Error())
	}

	list, err := s.Pl.GetList()
	if err != nil {
		s.Pl.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("playlist getting error in DeleteSong")
		return nil, err
	}

	for _, song := range list {
		dur := utils.ConvertFromSecondsToString(song.Duration)
		songRes := pb.Song{
			Name:     song.Name,
			Duration: dur,
		}
		res.Playlist = append(res.Playlist, &songRes)
	}
	err = s.Db.Delete(req.Name)
	if err != nil {
		s.Pl.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("deleting from database error")
		return nil, st.Error(codes.Internal, err.Error())
	}
	s.Pl.Logger.Info().Msg(fmt.Sprintf("[%s] deleted from playlist", req.Name))
	return &res, nil
}
