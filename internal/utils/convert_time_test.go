package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhas-off/test-music-player/internal/service"
)

func TestConvertFromSecondsToString(t *testing.T) {
	testTable := []struct {
		seconds  int
		expected string
	}{
		{
			seconds:  25,
			expected: "00:00:25",
		},
		{
			seconds:  -1,
			expected: "",
		},
		{
			seconds:  0,
			expected: "00:00:00",
		},
		{
			seconds:  86399,
			expected: "23:59:59",
		},
	}

	for _, testCase := range testTable {
		result := ConvertFromSecondsToString(testCase.seconds)
		assert.Equal(t, testCase.expected, result, fmt.Sprintf("incorrect result, expected %s, got %s",
			testCase.expected, result))
	}
}

func TestParseTimeToSeconds(t *testing.T) {
	testTable := []struct {
		time     string
		expected int
	}{
		{
			expected: 17,
			time:     "00:00:17",
		},
		{
			expected: 0,
			time:     "00:00:00",
		},
		{
			expected: 81999,
			time:     "22:46:39",
		},
	}

	for _, testCase := range testTable {
		result, err := ParseTimeToSeconds(testCase.time)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, testCase.expected, result, fmt.Sprintf("incorrect result, expected %d, got %d",
			testCase.expected, result))
	}
	t.Run("empty string test", func(t *testing.T) {
		result, err := ParseTimeToSeconds("")
		if err == nil {
			t.Error(err)
		}
		assert.Equal(t, 0, result, fmt.Sprintf("incorrect result, expected %d, got %d",
			0, result))
	})
}

func TestConvertFromSongProcToString(t *testing.T) {
	testTable := []struct {
		songPlay service.SongPlay
		expected string
	}{
		{
			service.SongPlay{
				Name:        "string",
				Duration:    10,
				CurrentTime: 9,
				Playing:     false,
				Exist:       true,
			},
			"00:00:09 of 00:00:10",
		},
		{
			service.SongPlay{
				Name:        "string",
				Duration:    70,
				CurrentTime: 60,
				Playing:     false,
				Exist:       true,
			},
			"00:01:00 of 00:01:10",
		},
		{
			service.SongPlay{
				Name:        "string",
				Duration:    0,
				CurrentTime: 0,
				Playing:     false,
				Exist:       true,
			},
			"00:00:00 of 00:00:00",
		},
		{
			service.SongPlay{
				Name:        "string",
				Duration:    -1,
				CurrentTime: -1,
				Playing:     false,
				Exist:       true,
			},
			" of ",
		},
	}

	for _, testCase := range testTable {
		result := ConvertSongStatsToString(testCase.songPlay)
		assert.Equal(t, testCase.expected, result, fmt.Sprintf("incorrect result, expected %s, got %s",
			testCase.expected, result))
	}
}
