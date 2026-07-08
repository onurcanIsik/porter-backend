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

func (r *userRepo) CreateUser(user *models.UserModel) error {
	query := `INSERT INTO users (
	user_mail, 
	user_name, 
	user_profile_url, 
	is_premium, 
	user_job_title, 
	user_device_id, 
	user_created_at, 
	user_updated_at, 
	provider, 
	provider_id) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err := r.db.Exec(context.Background(), query,
		user.UserMail,
		user.UserName,
		user.UserProfileUrl,
		user.IsPremium,
		user.UserJobTitle,
		user.UserDeviceId,
		user.UserCreatedAt,
		user.UserUpdatedAt,
		user.Provider,
		user.ProviderId,
	)

	return err
}

func (r *userRepo) GetUserByMail(mail string) (*models.UserModel, error) {
	query := `SELECT id, user_mail, user_name, user_profile_url, is_premium, user_job_title, user_device_id, user_created_at, user_updated_at, provider, provider_id FROM users WHERE user_mail = $1`
	var user models.UserModel
	err := r.db.QueryRow(context.Background(), query, mail).Scan(
		&user.ID,
		&user.UserMail,
		&user.UserName,
		&user.UserProfileUrl,
		&user.IsPremium,
		&user.UserJobTitle,
		&user.UserDeviceId,
		&user.UserCreatedAt,
		&user.UserUpdatedAt,
		&user.Provider,
		&user.ProviderId,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
