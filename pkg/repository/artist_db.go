package repository

import (
	"database/sql"
	"errors"
	"fmt"

	model "github.com/vnSasa/music-market-api/model"
)

type ArtistDB struct {
	db *sql.DB
}

func NewArtistDB(db *sql.DB) *ArtistDB {
	return &ArtistDB{db: db}
}

func (r *ArtistDB) CreateArtist(artist model.ArtistList) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	
	query := fmt.Sprintf("INSERT INTO %s (name_artist, date_of_birth, about_artist)"+
		"VALUES (?, ?, ?)", artistTable)
	
	_, err = r.db.Exec(query, artist.Name, artist.Birth, artist.About)
	if err != nil {
		return errors.New(err.Error())
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
