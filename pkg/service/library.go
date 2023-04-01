package service

import (
	model "github.com/vnSasa/music-market-api/model"
	"github.com/vnSasa/music-market-api/pkg/repository"
)

type LibraryService struct {
	repo repository.UsersLibrary
}

func NewUsersLibrary(repo repository.UsersLibrary) *LibraryService {
	return &LibraryService{repo: repo}
}

func (s *LibraryService) GetUserPlaylist(id int) ([]model.SongList, error) {
	return s.repo.GetUserPlaylist(id)
}

func (s *LibraryService) AddToPlaylist(userID, songID int) error {
	return s.repo.AddToPlaylist(userID, songID)
}

func (s *LibraryService) DeleteSongFromPlaylist(songID int) error {
	return s.repo.DeleteSongFromPlaylist(songID)
}
