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