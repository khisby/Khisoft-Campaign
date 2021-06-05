package campaign

type Service interface {
	GetCampaigns(UserId int) ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(UserId int) ([]Campaign, error) {
	if UserId != 0 {
		campaigns, err := s.repository.FindByUserID(UserId)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}
	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil

}
