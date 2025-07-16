package usecases

type CampaignInvitationUseCase struct {
}

func NewCampaignInvitationUseCase() *CampaignInvitationUseCase {
	return &CampaignInvitationUseCase{}
}

func (c *CampaignInvitationUseCase) CreateAnInvite() {}

func (c *CampaignInvitationUseCase) ReceiveAnInvite() {}
