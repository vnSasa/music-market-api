package service

import (
	model "github.com/vnSasa/music-market-api/model"
	"github.com/vnSasa/music-market-api/pkg/repository"
)

type DataFromUserService struct {
	repo repository.DataFromUser
}

func NewDataFromUser(repo repository.DataFromUser) *DataFromUserService {
	return &DataFromUserService{repo: repo}
}

func (s *DataFromUserService) CreateNewData(data model.DataFromUserList) error {
	return s.repo.CreateNewData(data)
}

func (s *DataFromUserService) CreateNewSong(song model.SongFromUserList) error {
	return s.repo.CreateNewSong(song)
}

func (s *DataFromUserService) GetAllData() ([]model.DataFromUserList, error) {
	return s.repo.GetAllData()
}

func (s *DataFromUserService) GetSongsFromUsers() ([]model.SongFromUserList, error) {
	return s.repo.GetSongsFromUsers()
}