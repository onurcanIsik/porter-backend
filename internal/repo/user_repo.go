package repo

import (
	"context"
	"porter/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(user *models.UserModelCreate) error {
	query := `INSERT INTO users (user_mail, 
	user_name, 
	user_profile_url, 
	is_premium, 
	user_token_count, 
	user_job_title, 
	user_device_id, 
	user_created_at, 
	user_updated_at, 
	provider, 
	provider_id) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := r.db.Exec(context.Background(), query,
		user.UserMail,
		user.UserName,
		user.UserProfileUrl,
		user.IsPremium,
		user.UserTokenCount,
		user.UserJobTitle,
		user.UserDeviceId,
		user.UserCreatedAt,
		user.UserUpdatedAt,
		user.Provider,
		user.ProviderId,
	)

	return err
}
