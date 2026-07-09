package models

import (
	"time"

	"github.com/google/uuid"
)

type RefreshTokenModel struct {
	ID           uuid.UUID `db:"id" json:"id"`
	UserID       uuid.UUID `db:"user_id" json:"user_id"`
	RefreshToken string    `db:"refresh_token" json:"refresh_token"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	ExpiredAt    time.Time `db:"expired_at" json:"expired_at"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}

type RefreshTokenModelRepository interface {
	SetRefreshToken(token *RefreshTokenModel) error
	GetByTokenHash(tokenHash string) (*RefreshTokenModel, error)
	DeleteByTokenHash(tokenHash string) error
}
