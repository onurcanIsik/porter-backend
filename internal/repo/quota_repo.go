package repo

import (
	"context"
	"porter/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type QuotaRepo struct {
	db *pgxpool.Pool
}

func NewQuotaRepo(db *pgxpool.Pool) *QuotaRepo {
	return &QuotaRepo{db: db}
}

func (r *QuotaRepo) GetQuotaByUserID(userID uuid.UUID) (*models.QuotaModel, error) {
	query := `SELECT quota_request,quota_endpoint,quota_bandwidth FROM quotas WHERE user_id = $1`
	var quotaRequest, quotaEndpoint, quotaBandwidth int
	err := r.db.QueryRow(context.Background(), query, userID).Scan(&quotaRequest, &quotaEndpoint, &quotaBandwidth)
	if err != nil {
		return nil, err
	}
	quota := &models.QuotaModel{
		UserID:         userID,
		QuotaRequest:   quotaRequest,
		QuotaEndpoint:  quotaEndpoint,
		QuotaBandwidth: quotaBandwidth,
	}
	return quota, nil
}

func (r *QuotaRepo) UpdateQuota(userID uuid.UUID, quotaRequest, quotaEndpoint, quotaBandwidth int) error {
	query := `UPDATE quotas SET quota_request = $1, quota_endpoint = $2, quota_bandwidth = $3 WHERE user_id = $4`
	_, err := r.db.Exec(context.Background(), query, quotaRequest, quotaEndpoint, quotaBandwidth, userID)
	return err
}
