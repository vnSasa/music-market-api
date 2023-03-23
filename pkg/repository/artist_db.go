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

func (r *ArtistDB) GetAllArtists() ([]model.ArtistList, error) {
	var artists []model.ArtistList
	query := fmt.Sprintf("SELECT id, name_artist, date_of_birth, about_artist FROM %s", artistTable)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var artist model.ArtistList
		err = rows.Scan(&artist.ID, &artist.Name, &artist.Birth, &artist.About)
		if err != nil {
			return nil, err
		}
		artists = append(artists, artist)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return artists, nil
}

func (r *ArtistDB) UpdateArtist(id int, artist model.ArtistList) error {
	query := fmt.Sprintf("UPDATE %s SET name_artist=?, date_of_birth=?, about_artist=? WHERE id=?", artistTable)

	_, err := r.db.Exec(query, artist.Name, artist.Birth, artist.About, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *ArtistDB) DeleteArtist(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=?", artistTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
