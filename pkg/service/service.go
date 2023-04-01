package service

import (
	model "github.com/vnSasa/music-market-api/model"
	"github.com/vnSasa/music-market-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) error
	GenerateToken(login, password string) (*model.TokenDetails, error)
	ParseToken(accessToken string) (*model.AccessTokenClaims, error)
	RefreshToken(refreshToken string) (string, error)
}

type Artists interface {
	CreateArtist(artist model.ArtistList) error
	GetAllArtists() ([]model.ArtistList, error)
	UpdateArtist(id int, artist model.ArtistList) error
	DeleteArtist(id int) error
}

type Songs interface {
	CreateSong(song model.SongList) error
	GetAllSongs() ([]model.SongList, error)
	UpdateSong(id int, song model.SongList) error
	DeleteSong(id int) error
	GetPlaylist(id int) ([]model.SongList, error)
}

type UsersLibrary interface {
	GetUserPlaylist(id int) ([]model.SongList, error)
	AddToPlaylist(userID, songID int) error
}

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
		UsersLibrary:  NewUsersLibrary(repos.UsersLibrary),
	}
}
