package models

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	ID             uuid.UUID `db:"id" json:"id"`
	UserMail       string    `db:"user_mail" json:"user_mail"`
	UserName       string    `db:"user_name" json:"user_name"`
	UserProfileUrl string    `db:"user_profile_url" json:"user_profile_url"`
	IsPremium      bool      `db:"is_premium" json:"is_premium"`
	UserJobTitle   string    `db:"user_job_title" json:"user_job_title"`
	UserDeviceId   string    `db:"user_device_id" json:"user_device_id"`
	UserCreatedAt  time.Time `db:"user_created_at" json:"user_created_at"`
	UserUpdatedAt  time.Time `db:"user_updated_at" json:"user_updated_at"`
	Provider       string    `db:"provider" json:"provider"`       // -- 'google' | 'github'
	ProviderId     string    `db:"provider_id" json:"provider_id"` // -- provider'ın verdiği sub
}

type UserModelRepository interface {
	CreateUser(user *UserModel) error
	//UpdateUser(user *UserModel) error
	GetUserByMail(mail string) (*UserModel, error)
	//GetUserById(id uuid.UUID) (*UserModel, error)
}

type UserModelUpdate struct {
	UserName      *string    `db:"user_name" json:"user_name"`
	UserJobTitle  *string    `db:"user_job_title" json:"user_job_title"`
	UserUpdatedAt *time.Time `db:"user_updated_at" json:"user_updated_at"`
}

type UserModelCreate struct {
	ID             uuid.UUID `db:"id" json:"id"`
	UserMail       string    `db:"user_mail" json:"user_mail"`
	UserName       string    `db:"user_name" json:"user_name"`
	UserProfileUrl string    `db:"user_profile_url" json:"user_profile_url"`
	IsPremium      bool      `db:"is_premium" json:"is_premium"`
	UserJobTitle   string    `db:"user_job_title" json:"user_job_title"`
	UserDeviceId   string    `db:"user_device_id" json:"user_device_id"`
	UserCreatedAt  time.Time `db:"user_created_at" json:"user_created_at"`
	UserUpdatedAt  time.Time `db:"user_updated_at" json:"user_updated_at"`
	Provider       string    `db:"provider" json:"provider"`
	ProviderId     string    `db:"provider_id" json:"provider_id"`
}
