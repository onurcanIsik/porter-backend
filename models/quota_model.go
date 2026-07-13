package models

import (
	"time"

	"github.com/google/uuid"
)

type QuotaModel struct {
	ID             uuid.UUID `db:"id" json:"id"`
	UserID         uuid.UUID `db:"user_id" json:"user_id"`
	QuotaRequest   int       `db:"quota_request" json:"quota_request"`
	QuotaEndpoint  int       `db:"quota_endpoint" json:"quota_endpoint"`
	QuotaBandwidth int       `db:"quota_bandwidth" json:"quota_bandwidth"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
	ExpiredAt      time.Time `db:"expired_at" json:"expired_at"`
}

type QuotaModelRepository interface {
	GetQuotaByUserID(userID uuid.UUID) (*QuotaModel, error)
	UpdateQuota(userID uuid.UUID, quotaRequest, quotaEndpoint, quotaBandwidth int) error
}
