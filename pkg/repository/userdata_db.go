package repository

import (
	model "github.com/vnSasa/music-market-api/model"
	"database/sql"
	"errors"
	"fmt"
)

type DataFromUserDB struct {
	db *sql.DB
}

func NewDataFromUserDB(db *sql.DB) *DataFromUserDB {
	return &DataFromUserDB{db: db}
}

func (r *DataFromUserDB) CreateNewData(data model.DataFromUserList) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := fmt.Sprintf("INSERT INTO %s (name_artist, date_of_birth, about_artist, name_song, genre, second_genre, year_of_release)"+
		"VALUES (?, ?, ?, ?, ?, ?, ?)", dataFromUserTable)

	_, err = r.db.Exec(query, data.NameArtist, data.Birth, data.About, data.NameSong, data.Genre, data.Genre2, data.Year)
	if err != nil {
		return errors.New(err.Error())
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *DataFromUserDB) CreateNewSong(song model.SongFromUserList) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var count int
	checkQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE artist_id = ? AND name_song = ?", songFromUserTable)
	err = r.db.QueryRow(checkQuery, song.ArtistID, song.Name).Scan(&count)
	if err != nil {
		return errors.New(err.Error())
	}
	if count > 0 {
		return errors.New("Song already exists")
	}

	query := fmt.Sprintf("INSERT INTO %s (artist_id, name_song, genre, second_genre, year_of_release)"+
		"VALUES (?, ?, ?, ?, ?)", songFromUserTable)

	_, err = r.db.Exec(query, song.ArtistID, song.Name, song.Genre, song.Genre2, song.Year)
	if err != nil {
		return errors.New(err.Error())
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	
	return nil
}