package repository

import (
	"context"

	"github.com/zhas-off/test-music-player/internal/models"
	utils "github.com/zhas-off/test-music-player/internal/utils"
)

func (s *PlaylistDatabase) Load() ([]models.Song, error) {
	var songs []models.Song

	rows, err := s.Conn.Query(context.Background(), "SELECT * FROM playlist")
	if err != nil {
		return songs, err
	}

	for rows.Next() {
		var name string
		var duration string
		if err := rows.Scan(&name, &duration); err != nil {
			return songs, err
		}
		durationValue, err := utils.ParseTimeToSeconds(duration)
		if err != nil {
			return songs, err
		}
		song := models.Song{
			Name:     name,
			Duration: durationValue,
		}
		songs = append(songs, song)
	}

	if err := rows.Err(); err != nil {
		return songs, err
	}
	return songs, err
}

func (s *PlaylistDatabase) Add(name, dur string) error {
	_, err := s.Conn.Exec(context.Background(), `INSERT INTO playlist (song, duration) VALUES ($1, $2)`, name, dur)
	if err != nil {
		return err
	}
	return nil
}

func (s PlaylistDatabase) Delete(name string) error {
	_, err := s.Conn.Exec(context.Background(), `DELETE FROM playlist WHERE song=$1;`, name)
	if err != nil {
		return err
	}
	return nil
}

func (s *PlaylistDatabase) Update(list []models.Song) error {
	_, err := s.Conn.Exec(context.Background(), `DELETE FROM playlist;`)
	if err != nil {
		return err
	}
	for _, song := range list {
		time := utils.ConvertFromSecondsToString(song.Duration)
		_, err := s.Conn.Exec(context.Background(), `INSERT INTO playlist (song, duration) VALUES ($1, $2);`, song.Name, time)
		if err != nil {
			return err
		}
	}
	return nil
}
