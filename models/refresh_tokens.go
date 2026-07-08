package models

type RefreshTokenModel struct {
	ID           string `db:"id" json:"id"`
	UserID       string `db:"user_id" json:"user_id"`
	RefreshToken string `db:"refresh_token" json:"refresh_token"`
	AccessToken  string `db:"access_token" json:"access_token"`
}

type RefreshTokenModelRepository interface {
	CreateRefreshToken(token *RefreshTokenModel) error
	GetRefreshTokenByUserID(userID string) (*RefreshTokenModel, error)
	DeleteRefreshTokenByUserID(userID string) error
}
