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
