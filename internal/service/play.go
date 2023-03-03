package service

import (
	"errors"

	"github.com/zhas-off/test-music-player/internal/models"
)

type SongPlay struct {
	Name        string
	CurrentTime int
	Duration    int

	Playing bool
	Exist   bool
}

func (pl *Playlist) Play() SongPlay {
	var data SongPlay
	pl.mutex.RLock()
	pl.PlayChan <- struct{}{}
	select {
	case data = <-pl.RequestChan:
		break
	}
	pl.playing = true
	pl.mutex.RUnlock()
	return data
}

func (pl *Playlist) Pause() SongPlay {
	var data SongPlay
	pl.mutex.RLock()
	pl.StopChan <- struct{}{}
	select {
	case data = <-pl.RequestChan:
		break
	}
	pl.playing = false
	pl.mutex.RUnlock()
	return data
}

func (pl *Playlist) Next() SongPlay {
	var data SongPlay
	pl.mutex.RLock()
	if pl.current.currentElem == nil {
		pl.NextChan <- false
	} else {
		if pl.current.currentElem.Next() == nil {
			pl.NextChan <- false
		} else {
			pl.current.currentElem = pl.current.currentElem.Next()
			pl.NextChan <- true
		}
	}

	select {
	case data = <-pl.RequestChan:
		break
	}
	pl.mutex.RUnlock()
	return data
}

func (pl *Playlist) Prev() SongPlay {
	var data SongPlay
	pl.mutex.RLock()

	if pl.current.currentElem == nil {
		pl.NextChan <- false
	} else {
		if pl.current.currentElem.Prev() == nil {
			pl.PrevChan <- false
		} else {
			pl.current.currentElem = pl.current.currentElem.Prev()
			pl.PrevChan <- true
		}
	}

	select {
	case data = <-pl.RequestChan:
		break
	}
	pl.mutex.RUnlock()
	return data
}

func (pl *Playlist) Status() SongPlay {
	var data SongPlay
	pl.mutex.RLock()
	pl.StatusChan <- struct{}{}
	select {
	case data = <-pl.RequestChan:
		break
	}
	pl.mutex.RUnlock()
	return data
}

// Изменение плейлиста

func (pl *Playlist) AddNewSong(song models.Song) bool {
	pl.mutex.Lock()
	defer pl.mutex.Unlock()

	for e := pl.list.Front(); e != nil; e = e.Next() {
		tmp, ok := e.Value.(models.Song)
		if !ok {
			return false
		}
		if tmp.Name == song.Name {
			return false
		}
	}

	el := pl.list.PushBack(song)
	if el == nil {
		return false
	}
	return true
}

func (pl *Playlist) GetList() ([]models.Song, error) {
	var res []models.Song
	pl.mutex.Lock()
	defer pl.mutex.Unlock()
	for e := pl.list.Front(); e != nil; e = e.Next() {
		tmp, ok := e.Value.(models.Song)
		if !ok {
			return res, errors.New("element to Song converting error")
		}
		res = append(res, tmp)
	}
	return res, nil
}

func (pl *Playlist) DeleteSong(name string) error {
	pl.mutex.Lock()
	defer pl.mutex.Unlock()

	pl.StatusChan <- struct{}{}
	var data SongPlay
	select {
	case data = <-pl.RequestChan:
		break
	}
	if pl.current.currentElem == nil {
		return errors.New("playlist is empty")
	}

	el, ok := pl.current.currentElem.Value.(models.Song)
	if !ok {
		return errors.New("element to Song converting error")
	}
	e := pl.list.Len()
	if e == 0 {
		return errors.New("song does not exist in playlist")
	}
	for e := pl.list.Front(); e != nil; e = e.Next() {
		tmp, ok := e.Value.(models.Song)
		if !ok {
			return errors.New("element to Song converting error")
		}
		if tmp.Name == name {
			if name == el.Name {
				if data.Playing {
					return errors.New("can't delete song while playing")
				}
				pl.list.Remove(e)
				break
			} else {
				pl.list.Remove(e)
				break
			}
		}
		if e.Next() == nil && tmp.Name != name {
			return errors.New("song does not exist in playlist")
		}
	}

	return nil
}

func (pl *Playlist) LoadListToPlaylistFromDatabase(databaseList []models.Song) {
	for _, s := range databaseList {
		pl.list.PushBack(s)
	}
}

// nextChannelsProc - отправка данных в каналы, повторяющийся код
func (pl *Playlist) nextChannelsProc() string {
	if pl.current.currentElem != nil {
		el, _ := pl.current.currentElem.Value.(models.Song)
		pl.RequestChan <- SongPlay{Name: el.Name, CurrentTime: 0, Duration: el.Duration, Exist: true}

		return "next"
	}
	return ""
}

// prevChannelsProc - отправка данных в каналы, повторяющийся код
func (pl *Playlist) prevChannelsProc() string {
	if pl.current.currentElem != nil {
		el, _ := pl.current.currentElem.Value.(models.Song)
		pl.RequestChan <- SongPlay{Name: el.Name, CurrentTime: 0, Duration: el.Duration, Exist: true}

		return "prev"
	}
	return ""
}
