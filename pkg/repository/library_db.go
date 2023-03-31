package repository

import (
	"database/sql"
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
