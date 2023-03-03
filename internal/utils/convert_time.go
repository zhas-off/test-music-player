package utils

import (
	"fmt"
	"time"

	"github.com/zhas-off/test-music-player/internal/service"
)

func ConvertSongStatsToString(songStats service.SongPlay) string {
	s1 := ConvertFromSecondsToString(songStats.Duration)
	s2 := ConvertFromSecondsToString(songStats.CurrentTime)
	return fmt.Sprintf("%s of %s", s2, s1)
}

func ConvertFromSecondsToString(seconds int) string {
	if seconds < 0 {
		return ""
	}
	duration := time.Duration(seconds) * time.Second
	Hours := int(duration.Hours())
	Minutes := int(duration.Minutes()) % 60
	Seconds := int(duration.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", Hours, Minutes, Seconds)
}

func ParseTimeToSeconds(timeStr string) (int, error) {
	t, err := time.Parse("15:04:05", timeStr)
	if err != nil {
		return 0, err
	}
	return t.Hour()*3600 + t.Minute()*60 + t.Second(), nil
}
