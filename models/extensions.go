package models

import "time"

type Extension struct {
	Extension   int    `db:"extension" ,json:"extension"`
	Secret      string `db:"secret" ,json:"secret"`
	ServerID    int    `db:"server_id" ,json:"server_id"`
	DisplayName string `db:"display_name" ,json:"display_name"`

	ID        int       `db:"id" ,json:"id"`
	CreatedAt time.Time `db:"created_at" ,json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" ,json:"updated_at"`
}

func GetExtensions() (*[]Extension, error) {
	var extensions []Extension

	err := client.Select(&extensions,
		"SELECT id, extension, secret, server_id, display_name, created_at, updated_at " +
		"FROM public.extensions " +
		"ORDER BY display_name ASC;")
	if err != nil {
		return nil, err
	}

	return &extensions, nil
}