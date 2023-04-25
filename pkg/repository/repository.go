package repository

import (
	"database/sql"

	model "github.com/vnSasa/music-market-api/model"
)

type Authorization interface {
	CreateUser(user model.User) error
	GetUserID(login, password string) (int, error)
	GetUserByID(id int) (*model.User, error)
	UpdateUserWithPassword(userID int, input model.User) error
	UpdateUserWithoutPassword(userID int, input model.User) error
}

type Artists interface {
	CreateArtist(artist model.ArtistList) error
	GetAllArtists() ([]model.ArtistList, error)
	GetArtistByID(artistID int) (*model.ArtistList, error)
	UpdateArtist(id int, artist model.ArtistList) error
	DeleteArtist(id int) error
}

type Songs interface {
	CreateSong(song model.SongList) error
	GetAllSongs() ([]model.SongList, error)
	GetSongByID(songID int) (*model.SongList, error)
	UpdateSong(id int, song model.SongList) error
	AddRating(songID, ratingPlus int) error
	DeleteSong(id int) error
	GetPlaylist(id int) ([]model.SongList, error)
}

type UsersLibrary interface {
	GetUserPlaylist(id int) ([]model.SongList, error)
	AddToPlaylist(userID, songID int) error
	DeleteSongFromPlaylist(songID int) error
}

type UsersTop interface{}

type Repository struct {
	Authorization
	Artists
	Songs
	UsersLibrary
	UsersTop
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthDB(db),
		Artists:       NewArtistDB(db),
		Songs:         NewSongDB(db),
		UsersLibrary:  NewLibraryDB(db),
	}
}
