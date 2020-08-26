package controllers

import (
	"freelancertest/models"
	"freelancertest/services/auth"
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"freelancertest/restapi/operations/auth_data"
)

// QueryAuthSubsDataHandlerFunc handles GET /subscription-data/{ueId}/authentication-data/authentication-subscription
func QueryAuthSubsDataHandlerFunc(params auth_data.QueryAuthSubsDataParams) middleware.Responder {

	service := auth.NewAuthenticationSubscriptionService()
	if service == nil {
		return auth_data.NewQueryAuthSubsDataDefault(http.StatusInternalServerError).
			WithPayload(models.ProblemDetails{
				Code:   "Server not responding",
				Detail: "failed to connect",
				Status: http.StatusInternalServerError,
			})
	}

	payload, problem := service.AuthenticationSubscriptionSearch(params.UeID)
	if problem != nil {
		return auth_data.NewQueryAuthSubsDataDefault(int(problem.Status)).WithPayload(problem)
	}

	return auth_data.NewQueryAuthSubsDataOK().WithPayload(payload)
}
