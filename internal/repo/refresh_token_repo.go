package repo

import (
	"context"
	"porter/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type refreshTokenRepo struct {
	db *pgxpool.Pool
}

func NewRefreshTokenRepo(db *pgxpool.Pool) *refreshTokenRepo {
	return &refreshTokenRepo{db: db}
}

func (r *refreshTokenRepo) SetRefreshToken(token *models.RefreshTokenModel) error {
	query := `INSERT INTO refresh_tokens (id, user_id, refresh_token, updated_at, expired_at, created_at) 
	VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(context.Background(), query,
		token.ID,
		token.UserID,
		token.RefreshToken,
		token.UpdatedAt,
		token.ExpiredAt,
		token.CreatedAt,
	)
	return err
}

func (r *refreshTokenRepo) GetByTokenHash(tokenHash string) (*models.RefreshTokenModel, error) {
	query := `SELECT id, user_id, refresh_token, updated_at, expired_at, created_at FROM refresh_tokens WHERE refresh_token = $1`
	var token models.RefreshTokenModel
	err := r.db.QueryRow(context.Background(), query, tokenHash).Scan(
		&token.ID,
		&token.UserID,
		&token.RefreshToken,
		&token.UpdatedAt,
		&token.ExpiredAt,
		&token.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *refreshTokenRepo) DeleteByTokenHash(tokenHash string) error {
	query := `DELETE FROM refresh_tokens WHERE refresh_token = $1`
	_, err := r.db.Exec(context.Background(), query, tokenHash)
	return err
}
