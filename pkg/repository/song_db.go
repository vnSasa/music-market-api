package repository

import (
	"database/sql"
	"errors"
	"fmt"

	model "github.com/vnSasa/music-market-api/model"
)

type SongDB struct {
	db *sql.DB
}

func NewSongDB(db *sql.DB) *SongDB {
	return &SongDB{db: db}
}

func (r *SongDB) CreateSong(song model.SongList) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := fmt.Sprintf("INSERT INTO %s (artist_id, name_song, genre, second_genre, year_of_release)"+
		"VALUES (?, ?, ?, ?, ?)", songTable)

	_, err = r.db.Exec(query, song.ArtistID, song.Name, song.Genre, song.Genre2, song.Year)
	if err != nil {
		return errors.New(err.Error())
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *SongDB) GetAllSongs() ([]model.SongList, error) {
	var songs []model.SongList
	query := fmt.Sprintf("SELECT id, artist_id, name_song, genre, second_genre, year_of_release FROM %s", songTable)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var song model.SongList
		err = rows.Scan(&song.ID, &song.ArtistID, &song.Name, &song.Genre, &song.Genre2, &song.Year)
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return songs, nil
}

func (r *SongDB) UpdateSong(id int, song model.SongList) error {
	query := fmt.Sprintf("UPDATE %s SET artist_id=?, name_song=?, genre=?, second_genre=?, year_of_release=? WHERE id=?", songTable)
	 
	_, err := r.db.Exec(query, song.ArtistID, song.Name, song.Genre, song.Genre2, song.Year, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *SongDB) DeleteSong(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=?", songTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *SongDB) GetPlaylist(id int) ([]model.SongList, error) {
	var songs []model.SongList
	query := fmt.Sprintf("SELECT id, name_song, genre, second_genre, year_of_release FROM %s WHERE artist_id=?", songTable)

	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var song model.SongList
		err = rows.Scan(&song.ID, &song.Name, &song.Genre, &song.Genre2, &song.Year)
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return songs, nil
}
