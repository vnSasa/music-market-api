package model

type ArtistList struct {
	ID    int
	Name  string `json:"name_artist" db:"name_artist" form:"name_artist" binding:"required"`
	Birth string `json:"date_of_birth" db:"date_of_birth" form:"date_of_birth" binding:"required"`
	About string `json:"about_artist" db:"about_artist" form:"about_artist"`
}

type SongList struct {
	ID       int
	ArtistID int    `json:"artist_id" db:"artist_id" form:"artist_id" binding:"required"`
	Name     string `json:"name_song" db:"name_song" form:"name_song" binding:"required"`
	Genre    string `json:"genre" db:"genre" form:"genre" binding:"required"`
	Genre2   string `json:"second_genre" db:"second_genre" form:"second_genre"`
	Year     int    `json:"year_of_release" db:"year_of_release" form:"year_of_release" binding:"required"`
}

type Library struct {
	ID     int
	UserID int
	SongID int
}

type SongData struct {
	ID         int
	ArtistID   int
	ArtistData string
	Name       string
	Genre      string
	Genre2     string
	Year       int
}
