package transaction

import (
	"errors"
	"khisoft_campign/campaign"
)

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

type Service interface {
	GetTransactionByCampaignID(campaignId GetCampaignTransactionInput) ([]Transaction, error)
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionByCampaignID(input GetCampaignTransactionInput) ([]Transaction, error) {

	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.ID != input.User.ID {
		return []Transaction{}, errors.New("Not an owner of the campaign")

	}
	transaction, err := s.repository.GetCampaignID(input.ID)
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
