package service

import (
	model "github.com/vnSasa/music-market-api/model"
	"github.com/vnSasa/music-market-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) error
	GenerateToken(login, password string) (*model.TokenDetails, error)
	ParseToken(accessToken string) (*model.AccessTokenClaims, error)
}

type Artists interface {
	CreateArtist(artist model.ArtistList) error
	GetAllArtists() ([]model.ArtistList, error)
	UpdateArtist(id int, artist model.ArtistList) error
}

type Songs interface {
	CreateSong(song model.SongList) error
	GetAllSongs() ([]model.SongList, error)
}

type UsersLibrary interface{}

type Service struct {
	Authorization
	Artists
	Songs
	UsersLibrary
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Artists:       NewArtistService(repos.Artists),
		Songs:         NewSongService(repos.Songs),
	}
}
