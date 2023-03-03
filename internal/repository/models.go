package repository

import (
	"context"

	"github.com/zhas-off/test-music-player/internal/models"
	"github.com/zhas-off/test-music-player/internal/utils"
)

// Загружаем случайные песни в БД
func (s *PlaylistDatabase) LoadToDatabaseIfNotExistBaseSongsSet() error {
	songs := []models.Song{
		{
			Name:     "Single Ladies - Beyoncé",
			Duration: 25,
		},
		{
			Name:     "Shake it Off - Taylor Swift",
			Duration: 32,
		},
		{
			Name:     "Rolling in the Deep - Adele",
			Duration: 36,
		},
		{
			Name:     "Blinding Lights - The Weeknd",
			Duration: 28,
		},
		{
			Name:     "Old Town Road - Lil Nas X",
			Duration: 31,
		},
		{
			Name:     "Poker Face - Lady Gaga",
			Duration: 30,
		},
		{
			Name:     "Starships - Nicki Minaj",
			Duration: 28,
		},
	}
	for _, song := range songs {
		_, err := s.Conn.Exec(context.Background(), `
		INSERT INTO playlist (song, duration)
		SELECT $1, $2
		WHERE
			NOT EXISTS (
				SELECT song FROM playlist WHERE song = $3
			);
		`, song.Name, utils.ConvertFromSecondsToString(song.Duration),
			song.Name)
		if err != nil {
			return err
		}
	}
	return nil
}
