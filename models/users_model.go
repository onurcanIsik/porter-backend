package models

import "time"

type UserModel struct {
	ID               int       `db:"id" json:"id"`
	UserMail         string    `db:"user_mail" json:"user_mail"`
	UserName         string    `db:"user_name" json:"user_name"`
	UserProfileUrl   string    `db:"user_profile_url" json:"user_profile_url"`
	IsPremium        bool      `db:"is_premium" json:"is_premium"`
	UserTokenCount   int       `db:"user_token_count" json:"user_token_count"`
	UserJobTitle     string    `db:"user_job_title" json:"user_job_title"`
	UserDeviceId     string    `db:"user_device_id" json:"user_device_id"`
	UserCreatedAt    time.Time `db:"user_created_at" json:"user_created_at"`
	UserUpdatedAt    time.Time `db:"user_updated_at" json:"user_updated_at"`
	UserAccessToken  string    `db:"user_access_token" json:"user_access_token"`
	UserRefreshToken string    `db:"user_refresh_token" json:"user_refresh_token"`
}

type UserModelRepository interface {
	CreateUser(user *UserModel) error
	UpdateUser(user *UserModel) error
	GetUserByMail(mail string) (*UserModel, error)
	GetUserById(id int) (*UserModel, error)
}

type UserModelUpdate struct {
	UserName         *string    `db:"user_name" json:"user_name"`
	UserJobTitle     *string    `db:"user_job_title" json:"user_job_title"`
	UserUpdatedAt    *time.Time `db:"user_updated_at" json:"user_updated_at"`
	UserAccessToken  *string    `db:"user_access_token" json:"user_access_token"`
	UserRefreshToken *string    `db:"user_refresh_token" json:"user_refresh_token"`
}

type UserModelCreate struct {
	UserMail         string    `db:"user_mail" json:"user_mail"`
	UserName         string    `db:"user_name" json:"user_name"`
	UserProfileUrl   string    `db:"user_profile_url" json:"user_profile_url"`
	IsPremium        bool      `db:"is_premium" json:"is_premium"`
	UserTokenCount   int       `db:"user_token_count" json:"user_token_count"`
	UserJobTitle     string    `db:"user_job_title" json:"user_job_title"`
	UserDeviceId     string    `db:"user_device_id" json:"user_device_id"`
	UserCreatedAt    time.Time `db:"user_created_at" json:"user_created_at"`
	UserUpdatedAt    time.Time `db:"user_updated_at" json:"user_updated_at"`
	UserAccessToken  string    `db:"user_access_token" json:"user_access_token"`
	UserRefreshToken string    `db:"user_refresh_token" json:"user_refresh_token"`
}

type UserSocialLoginModel struct {
	//todo Arastirilicak
}
