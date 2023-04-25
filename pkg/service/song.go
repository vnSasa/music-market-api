package service

import (
	model "github.com/vnSasa/music-market-api/model"
	"github.com/vnSasa/music-market-api/pkg/repository"
)

type SongService struct {
	repo repository.Songs
}

func NewSongService(repo repository.Songs) *SongService {
	return &SongService{repo: repo}
}

func (s *SongService) CreateSong(song model.SongList) error {
	return s.repo.CreateSong(song)
}

func (s *SongService) GetAllSongs() ([]model.SongList, error) {
	return s.repo.GetAllSongs()
}

func (s *SongService) GetSongByID(songID int) (*model.SongList, error) {
	return s.repo.GetSongByID(songID)
}

func (s *SongService) UpdateSong(id int, song model.SongList) error {
	return s.repo.UpdateSong(id, song)
}

func (s *SongService) AddRating(songID, ratingPlus int) error {
	return s.repo.AddRating(songID, ratingPlus)
}

func (s *SongService) DeleteSong(id int) error {
	return s.repo.DeleteSong(id)
}

func (s *SongService) GetPlaylist(id int) ([]model.SongList, error) {
	return s.repo.GetPlaylist(id)
}
