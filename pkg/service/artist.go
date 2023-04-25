package service

import (
	model "github.com/vnSasa/music-market-api/model"
	"github.com/vnSasa/music-market-api/pkg/repository"
)

type ArtistService struct {
	repo repository.Artists
}

func NewArtistService(repo repository.Artists) *ArtistService {
	return &ArtistService{repo: repo}
}

func (s *ArtistService) CreateArtist(artist model.ArtistList) error {
	return s.repo.CreateArtist(artist)
}

func (s *ArtistService) GetAllArtists() ([]model.ArtistList, error) {
	return s.repo.GetAllArtists()
}

func (s *ArtistService) GetArtistByID(artistID int) (*model.ArtistList, error) {
	return s.repo.GetArtistByID(artistID)
}

func (s *ArtistService) UpdateArtist(id int, artist model.ArtistList) error {
	return s.repo.UpdateArtist(id, artist)
}

func (s *ArtistService) DeleteArtist(id int) error {
	return s.repo.DeleteArtist(id)
}
