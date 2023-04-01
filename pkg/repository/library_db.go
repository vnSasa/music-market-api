package repository

import (
	"database/sql"
	"errors"
	"fmt"

	model "github.com/vnSasa/music-market-api/model"
)

type LibraryDB struct {
	db *sql.DB
}

func NewLibraryDB(db *sql.DB) *LibraryDB {
	return &LibraryDB{db: db}
}

func (r *LibraryDB) GetUserPlaylist(id int) ([]model.SongList, error) {
	var songs []model.SongList
	query := fmt.Sprintf("SELECT s.id, s.artist_id, s.name_song, s.genre, s.second_genre, s.year_of_release "+
		"FROM %s s JOIN %s ul ON s.id = ul.song_id WHERE ul.user_id = ?", songTable, libraryTable)

	rows, err := r.db.Query(query, id)
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

func (r *LibraryDB) AddToPlaylist(userID, songID int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var count int
	err = r.db.QueryRow("SELECT COUNT(*) FROM "+libraryTable+" WHERE user_id = ? AND song_id = ?", userID, songID).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("song already in playlist")
	}

	query := fmt.Sprintf("INSERT INTO %s (user_id, song_id) VALUES (?, ?)", libraryTable)

	_, err = r.db.Exec(query, userID, songID)
	if err != nil {
		return errors.New(err.Error())
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *LibraryDB) DeleteSongFromPlaylist(songID int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := fmt.Sprintf("DELETE FROM %s WHERE song_id=?", libraryTable)
	_, err = r.db.Exec(query, songID)
	if err != nil {
		return errors.New(err.Error())
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
