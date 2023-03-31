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
}

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
		Songs:         NewSongDB(db),
		UsersLibrary:  NewLibraryDB(db),
	}
}
