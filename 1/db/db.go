package db

import (
	"1/model"
	"database/sql"
)

// albums slice to seed record album data.
var Albums = []*model.Album{}

/*
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
*/

type DB interface {
	//GetTechnologies() ([]*model.Technology, error)
	GetAlbums() ([]*model.Album, error)
}

type MySQLDB struct {
	db *sql.DB
}

func NewDB(db *sql.DB) DB {
	return MySQLDB{db: db}
}

func (d MySQLDB) GetAlbums() ([]*model.Album, error) {
	rows, err := d.db.Query("SELECT id, title FROM album")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var albums []*model.Album
	for rows.Next() {
		a := new(model.Album)
		err = rows.Scan(&a.ID, &a.Title)
		if err != nil {
			return nil, err
		}
		albums = append(albums, a)
	}
	return albums, nil
}
