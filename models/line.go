package models

type Line struct {
	Enabled     bool
	Position    int    `db:"position" ,json:"position"`
	Extension   int    `db:"extension" ,json:"extension"`
	Secret      string `db:"secret" ,json:"secret"`
	Server      string `db:"server" ,json:"server"`
	DisplayName string `db:"display_name" ,json:"display_name"`
}
