package service

import (
	"porter/models"

	"github.com/google/uuid"
)

type QuotaService struct {
	quotaRepo models.QuotaModelRepository
}

func NewQuotaService(quotaRepo models.QuotaModelRepository) *QuotaService {
	return &QuotaService{quotaRepo: quotaRepo}
}

func (s *QuotaService) GetQuotaByUserID(userID uuid.UUID) (*models.QuotaModel, error) {
	return s.quotaRepo.GetQuotaByUserID(userID)

}

func (s *QuotaService) UpdateQuota(userID uuid.UUID, quota *models.QuotaModel) error {
	quota.UserID = userID
	return s.quotaRepo.UpdateQuota(quota.UserID, quota.QuotaRequest, quota.QuotaEndpoint, quota.QuotaBandwidth)
}
