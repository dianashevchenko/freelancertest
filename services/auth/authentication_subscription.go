package auth

import (
	"context"
	"fmt"
	"freelancertest/models"
	"freelancertest/repositories"
	"freelancertest/services"
	"net/http"
)

// AuthenticationSubscriptionService is an implementation of service working through LDAP
type AuthenticationSubscriptionService struct {
	repositories.LdapRepository
}

// NewAuthenticationSubscriptionService constructor
func NewAuthenticationSubscriptionService() *AuthenticationSubscriptionService {
	repo, err := services.CreateRepository(context.Background(), "subscriptionData.AuthData.AuthenticationSubscription")
	if err != nil {
		return nil
	}
	return &AuthenticationSubscriptionService{
		LdapRepository: *repo,
	}
}

// AuthenticationSubscriptionSearch returns object by it's UeID
func (s *AuthenticationSubscriptionService) AuthenticationSubscriptionSearch(ueID string) (*models.AuthenticationSubscription, *models.ProblemDetails) {

	subs, err := s.LdapRepository.Search(ueID)
	if err != nil {
		return nil, &models.ProblemDetails{
			Code:     "bad request",
			Detail:   err.Error(),
			Instance: fmt.Sprintf("/subscription-data/%s/authentication-data/authentication-subscription", ueID),
			Status:   http.StatusBadRequest,
			Title:    "authorization denied",
		}
	}
	return subs, nil
}
