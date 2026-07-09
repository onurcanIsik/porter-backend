package models

type QuotaModel struct {
	ID     int    `db:"id" json:"id"`
	UserID string `db:"user_id" json:"user_id"`
	Quota  int    `db:"quota" json:"quota"`
}
