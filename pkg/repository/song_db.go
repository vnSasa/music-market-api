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

	var count int
	checkQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE artist_id = ? AND name_song = ?", songTable)
	err = r.db.QueryRow(checkQuery, song.ArtistID, song.Name).Scan(&count)
	if err != nil {
		return errors.New(err.Error())
	}
	if count > 0 {
		return errors.New("Song already exists")
	}

	startRate := 0
	query := fmt.Sprintf("INSERT INTO %s (artist_id, name_song, genre, second_genre, year_of_release, rating)"+
		"VALUES (?, ?, ?, ?, ?, ?)", songTable)

	_, err = r.db.Exec(query, song.ArtistID, song.Name, song.Genre, song.Genre2, song.Year, startRate)
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
	query := fmt.Sprintf("SELECT id, artist_id, name_song, genre, second_genre, year_of_release, rating FROM %s", songTable)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var song model.SongList
		err = rows.Scan(&song.ID, &song.ArtistID, &song.Name, &song.Genre, &song.Genre2, &song.Year, &song.Rating)
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

func (r *SongDB) GetSongByID(songID int) (*model.SongList, error) {
	var songData model.SongList
	confirmSong := fmt.Sprintf("SELECT artist_id, name_song, genre, second_genre, year_of_release, rating FROM %s WHERE id = ?", songTable)
	row := r.db.QueryRow(confirmSong, songID)
	err := row.Scan(&songData.ArtistID, &songData.Name, &songData.Genre, &songData.Genre2, &songData.Year, &songData.Rating)
	if err != nil {
		return nil, errors.New("song not found")
	}

	return &songData, nil
}

func (r *SongDB) UpdateSong(id int, song model.SongList) error {
	query := fmt.Sprintf("UPDATE %s SET artist_id=?, name_song=?, genre=?, second_genre=?, year_of_release=? WHERE id=?", songTable)

	_, err := r.db.Exec(query, song.ArtistID, song.Name, song.Genre, song.Genre2, song.Year, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *SongDB) UpdateRating(songID int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var countLibrary int
	err = r.db.QueryRow("SELECT COUNT(*) FROM "+libraryTable+" WHERE song_id = ?", songID).Scan(&countLibrary)
	if err != nil {
		return err
	}
	var countTop int
	err = r.db.QueryRow("SELECT COUNT(*) FROM "+topTable+" WHERE song_id = ?", songID).Scan(&countTop)
	if err != nil {
		return err
	}
	count := countLibrary + countTop
	query := fmt.Sprintf("UPDATE %s SET rating=? WHERE id=?", songTable)
	_, err = r.db.Exec(query, count, songID)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
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
	query := fmt.Sprintf("SELECT id, name_song, genre, second_genre, year_of_release, rating FROM %s WHERE artist_id=?", songTable)

	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var song model.SongList
		err = rows.Scan(&song.ID, &song.Name, &song.Genre, &song.Genre2, &song.Year, &song.Rating)
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
