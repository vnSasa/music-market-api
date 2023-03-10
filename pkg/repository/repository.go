package repository

import (
	"database/sql"

	model "github.com/vnSasa/music-market-api/model"
)

type Authorization interface {
	CreateUser(user model.User) error
	GetUser(login, password string) (int, error)
}

type Artists interface {
	CreateArtist(artist model.ArtistList) error
	GetAllArtists() ([]model.ArtistList, error)
}

type Songs interface{
	CreateSong(song model.SongList) error
	GetAllSongs() ([]model.SongList, error)
}

type UsersLibrary interface{}

type Repository struct {
	Authorization
	Artists
	Songs
	UsersLibrary
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthDB(db),
		Artists:       NewArtistDB(db),
		Songs:	NewSongDB(db),
	}
}
