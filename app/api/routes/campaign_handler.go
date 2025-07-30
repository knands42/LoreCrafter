package routes

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/knands42/lorecrafter/app/api/middleware"
	"github.com/knands42/lorecrafter/app/api/utils"
	"github.com/knands42/lorecrafter/internal/domain"
	"github.com/knands42/lorecrafter/internal/usecases"
)

// CampaignHandler handles campaign-related HTTP requests
type CampaignHandler struct {
	campaignUseCase           *usecases.CampaignUseCase
	campaignInvitationUseCase *usecases.CampaignInvitationUseCase
	campaignMembersUseCase    *usecases.CampaignMembersUseCase
}

// NewCampaignHandler creates a new CampaignHandler
func NewCampaignHandler(
	campaignUseCase *usecases.CampaignUseCase,
	campaignInvitationUseCase *usecases.CampaignInvitationUseCase,
	campaignMembersUseCase *usecases.CampaignMembersUseCase,
) *CampaignHandler {
	return &CampaignHandler{
		campaignUseCase:           campaignUseCase,
		campaignInvitationUseCase: campaignInvitationUseCase,
		campaignMembersUseCase:    campaignMembersUseCase,
	}
}

// RegisterRoutes registers the campaign routes
func (h *CampaignHandler) RegisterRoutes(r chi.Router) {
	r.Route("/campaigns", func(r chi.Router) {
		r.Post("/", middleware.ErrorHandlerMiddleware(h.CreateCampaign))
		r.Get("/", middleware.ErrorHandlerMiddleware(h.ListUserCampaigns))
		r.Post("/{campaignID}/invitations", middleware.ErrorHandlerMiddleware(h.CreateCampaignInvitation))
		r.Patch("/invitations/{token}", middleware.ErrorHandlerMiddleware(h.ReceiveCampaignInvitation))
		r.Get("/invitations", middleware.ErrorHandlerMiddleware(h.GetCampaignInvitations))

		r.Route("/{campaignID}", func(r chi.Router) {
			r.Get("/", middleware.ErrorHandlerMiddleware(h.GetCampaign))
			r.Put("/", middleware.ErrorHandlerMiddleware(h.UpdateCampaign))
			r.Delete("/", middleware.ErrorHandlerMiddleware(h.DeleteCampaign))

			r.Route("/members", func(r chi.Router) {
				r.Get("/", middleware.ErrorHandlerMiddleware(h.GetCampaignMembers))
				r.Get("/{memberID}", middleware.ErrorHandlerMiddleware(h.GetCampaignMember))
				r.Delete("/{userID}", middleware.ErrorHandlerMiddleware(h.RemoveCampaignMember))
			})
		})

		r.Delete("/leave/{campaignID}", middleware.ErrorHandlerMiddleware(h.LeaveCampaign))
	})
}

// CreateCampaign handles campaign creation
// @Summary Create a new campaign
// @Description Create a new campaign with the provided details. If use_ai query parameter is true, AI will generate the campaign settings.
// @Tags campaigns
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body domain.CampaignCreationInput true "Campaign creation details"
// @Param use_gen_ai query bool false "Whether to use AI to generate campaign settings" default(false)
// @Success 201 {object} domain.Campaign "Campaign created successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid request body or parameters"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /api/campaigns [post]
func (h *CampaignHandler) CreateCampaign(w http.ResponseWriter, r *http.Request) error {
	// Parse query parameters
	useAI := r.URL.Query().Get("use_gen_ai") == "true"

	// Parse request body
	var input domain.CampaignCreationInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
	}

	// Get user ID from context
	userIDStr, ok := r.Context().Value(middleware.UserIDContextKey).(string)
	if !ok {
		return utils.WriteJSONError(w, http.StatusUnauthorized, "User ID not found")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid user ID")
	}

	// Create the campaign with either original or AI-enhanced input
	campaign, err := h.campaignUseCase.CreateCampaign(userID, useAI, input)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(campaign)
}

// GetCampaign handles retrieving a campaign by ID
// @Summary Get a campaign by ID
// @Description Get a campaign by ID if the user has access
// @Tags campaigns
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param campaignID path string true "Campaign ID"
// @Success 200 {object} domain.Campaign "Campaign retrieved successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid campaign ID"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 403 {object} utils.ErrorResponse "Insufficient permissions"
// @Failure 404 {object} utils.ErrorResponse "Campaign not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /api/campaigns/{campaignID} [get]
func (h *CampaignHandler) GetCampaign(w http.ResponseWriter, r *http.Request) error {
	campaignID, err := uuid.Parse(chi.URLParam(r, "campaignID"))
	if err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid campaign ID")
	}

	userIDStr, ok := r.Context().Value(middleware.UserIDContextKey).(string)
	if !ok {
		return utils.WriteJSONError(w, http.StatusUnauthorized, "User ID not found in context")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid user ID")
	}

	input := domain.GetCampaignInput{
		UserId:     userID,
		CampaignId: campaignID,
	}

	campaign, err := h.campaignUseCase.GetCampaign(input)
	if err != nil {
		switch {
		case errors.Is(err, usecases.ErrCampaignNotFound):
			return utils.WriteJSONError(w, http.StatusNotFound, "Campaign not found")
		default:
			return err
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "private, max-age=300")
	return json.NewEncoder(w).Encode(campaign)
}

// UpdateCampaign handles updating a campaign
// @Summary Update a campaign
// @Description Update a campaign if the user has GM permissions
// @Tags campaigns
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param campaignID path string true "Campaign ID"
// @Param input body domain.UpdateCampaignInput true "Campaign update details"
// @Success 200 {object} domain.Campaign "Campaign updated successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid request body or campaign ID"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 403 {object} utils.ErrorResponse "Insufficient permissions"
// @Failure 404 {object} utils.ErrorResponse "Campaign not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /api/campaigns/{campaignID} [put]
func (h *CampaignHandler) UpdateCampaign(w http.ResponseWriter, r *http.Request) error {
	var campaign domain.UpdateCampaignInput
	if err := json.NewDecoder(r.Body).Decode(&campaign); err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
	}

	campaignID, err := uuid.Parse(chi.URLParam(r, "campaignID"))
	if err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid campaign ID")
	}
	campaign.ID = campaignID

	userIDStr, ok := r.Context().Value(middleware.UserIDContextKey).(string)
	if !ok {
		return utils.WriteJSONError(w, http.StatusUnauthorized, "User ID not found in context")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid user ID")
	}

	err = h.campaignUseCase.UpdateCampaign(userID, campaign)
	if err != nil {
		switch {
		case errors.Is(err, usecases.ErrCampaignNotFound):
			return utils.WriteJSONError(w, http.StatusNotFound, "Campaign not found")
		default:
			return err
		}
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(campaign)
}

// DeleteCampaign handles deleting a campaign
// @Summary Delete a campaign
// @Description Delete a campaign if the user has GM permissions
// @Tags campaigns
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param campaignID path string true "Campaign ID"
// @Success 204 "Campaign deleted successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid campaign ID"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 403 {object} utils.ErrorResponse "Insufficient permissions"
// @Failure 404 {object} utils.ErrorResponse "Campaign not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /api/campaigns/{campaignID} [delete]
func (h *CampaignHandler) DeleteCampaign(w http.ResponseWriter, r *http.Request) error {
	var input domain.DeleteCampaignInput
	campaignID, err := uuid.Parse(chi.URLParam(r, "campaignID"))
	if err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid campaign ID")
	}
	input.ID = campaignID

	userIDStr, ok := r.Context().Value(middleware.UserIDContextKey).(string)
	if !ok {
		return utils.WriteJSONError(w, http.StatusUnauthorized, "User ID not found in context")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid user ID")
	}
	input.UserID = userID

	err = h.campaignUseCase.DeleteCampaign(input)
	if err != nil {
		switch {
		case errors.Is(err, usecases.ErrCampaignNotFound):
			return utils.WriteJSONError(w, http.StatusNotFound, "Campaign not found")
		default:
			return err
		}
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

// ListUserCampaigns handles listing all campaigns a user is a member of
// @Summary List user campaigns
// @Description List all campaigns a user is a member of
// @Tags campaigns
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} domain.Campaign "Campaigns retrieved successfully"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /api/campaigns [get]
func (h *CampaignHandler) ListUserCampaigns(w http.ResponseWriter, r *http.Request) error {
	userIDStr, ok := r.Context().Value(middleware.UserIDContextKey).(string)
	if !ok {
		return utils.WriteJSONError(w, http.StatusUnauthorized, "User ID not found in context")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid user ID")
	}

	campaigns, err := h.campaignUseCase.ListUserCampaigns(userID)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "private, max-age=60")
	return json.NewEncoder(w).Encode(campaigns)
}

// AddCampaignMember handles adding a user to a campaign
// @Summary Add a user to a campaign
// @Description Add a user to a campaign if the requester has GM permissions
// @Tags campaigns
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param campaignID path string true "Campaign ID"
// @Param input body domain.CreateCampaignMemberInput true "User ID and role"
// @Success 204 "User added to campaign successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid request body or campaign ID"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 403 {object} utils.ErrorResponse "Insufficient permissions"
// @Failure 404 {object} utils.ErrorResponse "Campaign not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /api/campaigns/{campaignID}/members [post]
func (h *CampaignHandler) AddCampaignMember(w http.ResponseWriter, r *http.Request) error {
	// parse payload
	var input domain.CreateCampaignMemberInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
	}

	// parse url params
	campaignID, err := uuid.Parse(chi.URLParam(r, "campaignID"))
	if err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid campaign ID")
	}

	userID, err := uuid.Parse(input.UserID.String())
	if err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid user ID")
	}

	// get requester id
	requesterIDStr, ok := r.Context().Value(middleware.UserIDContextKey).(string)
	if !ok {
		return utils.WriteJSONError(w, http.StatusUnauthorized, "User ID not found in context")
	}

	requesterID, err := uuid.Parse(requesterIDStr)
	if err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid requester ID")
	}

	// create the member campaign
	input.UserID = userID
	input.CampaignID = campaignID

	createdCampaignMember, err := h.campaignMembersUseCase.AddCampaignPlayerMember(requesterID, input)
	if err != nil {
		switch {
		case errors.Is(err, usecases.ErrCampaignNotFound):
			return utils.WriteJSONError(w, http.StatusNotFound, "Campaign not found")
		default:
			return err
		}
	}

	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(createdCampaignMember)
}

// RemoveCampaignMember handles removing a user from a campaign
// @Summary Remove a user from a campaign
// @Description Remove a user from a campaign if the requester has GM permissions
// @Tags campaigns
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param campaignID path string true "Campaign ID"
// @Param userID path string true "User ID"
// @Success 204 "User removed from campaign successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid campaign ID or user ID"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 403 {object} utils.ErrorResponse "Insufficient permissions"
// @Failure 404 {object} utils.ErrorResponse "Campaign not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /api/campaigns/{campaignID}/members/{userID} [delete]
func (h *CampaignHandler) RemoveCampaignMember(w http.ResponseWriter, r *http.Request) error {
	campaignID, err := uuid.Parse(chi.URLParam(r, "campaignID"))
	if err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid campaign ID")
	}

	userID, err := uuid.Parse(chi.URLParam(r, "userID"))
	if err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid user ID")
	}

	requesterIDStr, ok := r.Context().Value(middleware.UserIDContextKey).(string)
	if !ok {
		return utils.WriteJSONError(w, http.StatusUnauthorized, "User ID not found in context")
	}

	requesterID, err := uuid.Parse(requesterIDStr)
	if err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid requester ID")
	}

	err = h.campaignMembersUseCase.RemoveCampaignMember(campaignID, userID, requesterID)
	if err != nil {
		switch {
		case errors.Is(err, usecases.ErrCampaignNotFound):
			return utils.WriteJSONError(w, http.StatusNotFound, "Campaign not found")
		default:
			return err
		}
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

// LeaveCampaign handles a user leaving a campaign
// @Summary Leave a campaign
// @Description Allow a user to leave a campaign
// @Tags campaigns
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param campaignID path string true "Campaign ID"
// @Success 204 "Left campaign successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid campaign ID"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 404 {object} utils.ErrorResponse "Campaign not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /api/campaigns/leave/{campaignID} [delete]
func (h *CampaignHandler) LeaveCampaign(w http.ResponseWriter, r *http.Request) error {
	campaignID, err := uuid.Parse(chi.URLParam(r, "campaignID"))
	if err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid campaign ID")
	}

	userIDStr, ok := r.Context().Value(middleware.UserIDContextKey).(string)
	if !ok {
		return utils.WriteJSONError(w, http.StatusUnauthorized, "User ID not found in context")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid user ID")
	}

	err = h.campaignMembersUseCase.LeaveCampaign(campaignID, userID)
	if err != nil {
		if errors.Is(err, usecases.ErrCampaignNotFound) {
			return utils.WriteJSONError(w, http.StatusNotFound, "Campaign not found")
		}
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

// GetCampaignMembers handles listing all members of a campaign
// @Summary List campaign members
// @Description List all members of a campaign if the user has access
// @Tags campaigns
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param campaignID path string true "Campaign ID"
// @Success 200 {array} domain.CampaignMember "Campaign members retrieved successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid campaign ID"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 403 {object} utils.ErrorResponse "Insufficient permissions"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /api/campaigns/{campaignID}/members [get]
func (h *CampaignHandler) GetCampaignMembers(w http.ResponseWriter, r *http.Request) error {
	// parse url params
	campaignID, err := uuid.Parse(chi.URLParam(r, "campaignID"))
	if err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid campaign ID")
	}

	userIDStr, ok := r.Context().Value(middleware.UserIDContextKey).(string)
	if !ok {
		return utils.WriteJSONError(w, http.StatusUnauthorized, "User ID not found in context")
	}

	requesterID, err := uuid.Parse(userIDStr)
	if err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid user ID")
	}

	// get members
	input := domain.GetCampaignMembersInput{
		CampaignID:  campaignID,
		RequesterID: requesterID,
	}
	members, err := h.campaignMembersUseCase.GetCampaignMembers(input)
	if err != nil {
		if errors.Is(err, usecases.ErrFailedToRetrieveAMember) {
			return utils.WriteJSONError(w, http.StatusNotFound, "Insufficient permissions")
		}
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(members)
}

// GetCampaignMember handles listing all members of a campaign
// @Summary List campaign members
// @Description List all members of a campaign if the user has access
// @Tags campaigns
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param campaignID path string true "Campaign ID"
// @Success 200 {array} domain.CampaignMember "Campaign members retrieved successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid campaign ID"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 403 {object} utils.ErrorResponse "Insufficient permissions"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /api/campaigns/{campaignID}/members/{memberID} [get]
func (h *CampaignHandler) GetCampaignMember(w http.ResponseWriter, r *http.Request) error {
	// parse url params
	campaignID, err := uuid.Parse(chi.URLParam(r, "campaignID"))
	if err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid campaign ID")
	}

	memberID, err := uuid.Parse(chi.URLParam(r, "memberID"))
	if err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid member ID")
	}

	// get logged user
	userIDStr, ok := r.Context().Value(middleware.UserIDContextKey).(string)
	if !ok {
		return utils.WriteJSONError(w, http.StatusUnauthorized, "User ID not found in context")
	}

	requesterID, err := uuid.Parse(userIDStr)
	if err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid user ID")
	}

	// get members
	input := domain.NewGetCampaignMemberInput(campaignID, memberID, requesterID)
	member, err := h.campaignMembersUseCase.GetCampaignMember(*input)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(member)
}

// CreateCampaignInvitation handles creating a campaign invitation
// @Summary Create a campaign invitation
// @Description Create a campaign invitation if the user has GM permissions
// @Tags campaigns
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param campaignID path string true "Campaign ID"
// @Param input body domain.CreateCampaignInvitationInput true "Campaign invitation details"
// @Success 201 {object} domain.CampaignInvitation "Campaign invitation created successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid request body or campaign ID"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 403 {object} utils.ErrorResponse "Insufficient permissions"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /api/campaigns/{campaignID}/invitations [post]
func (h *CampaignHandler) CreateCampaignInvitation(w http.ResponseWriter, r *http.Request) error {
	// Parse request body
	var input domain.CreateCampaignInvitationInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
	}

	// Parse url param
	campaignID, err := uuid.Parse(chi.URLParam(r, "campaignID"))
	if err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid campaign ID")
	}

	// Get user ID from context
	userIDStr, ok := r.Context().Value(middleware.UserIDContextKey).(string)
	if !ok {
		return utils.WriteJSONError(w, http.StatusUnauthorized, "User ID not found")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid user ID")
	}

	// Create campaign invitation
	input.CampaignID = campaignID
	input.InvitedBy = userID

	createdInvitation, err := h.campaignInvitationUseCase.CreateAnInvite(input)
	if err != nil {
		switch {
		case errors.Is(err, usecases.ErrCreatingTheCampaignInvitation):
			return utils.WriteJSONError(w, http.StatusBadRequest, err.Error())
		case errors.Is(err, usecases.ErrGeneratingTheToken):
			return utils.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		case errors.Is(err, usecases.ErrUserIsAlreadyAMember):
			return utils.WriteJSONError(w, http.StatusConflict, err.Error())
		}

		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(createdInvitation)
}

// ReceiveCampaignInvitation handles receiving a campaign invitation
// @Summary Receive a campaign invitation
// @Description Receive a campaign invitation if the user has GM permissions
// @Tags campaigns
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param token path string true "Invite Token"
// @Param input body domain.ReceiveCampaignInvite true "Confirm or reject the campaign invite"
// @Success 204 "Campaign invitation received successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid campaign ID"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 403 {object} utils.ErrorResponse "Insufficient permissions"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /api/campaigns/invitations/{token} [patch]
func (h *CampaignHandler) ReceiveCampaignInvitation(w http.ResponseWriter, r *http.Request) error {
	// Parse url param
	token := chi.URLParam(r, "token")

	// Get user ID from context
	userIDStr, ok := r.Context().Value(middleware.UserIDContextKey).(string)
	if !ok {
		return utils.WriteJSONError(w, http.StatusUnauthorized, "User ID not found")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid user ID")
	}

	// Parse request body
	var input domain.ReceiveCampaignInvite
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
	}

	// Set the token and invitedBy
	input.Token = token

	err = h.campaignInvitationUseCase.ReceiveAnInvite(userID, input)
	if err != nil {
		switch {
		case errors.Is(err, usecases.ErrAddingTheCampaignMember):
			return utils.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		case errors.Is(err, usecases.ErrInviteIsNoLongerValid):
			return utils.WriteJSONError(w, http.StatusGone, err.Error())
		}
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

// GetCampaignInvitations handles listing all campaign invitations
// @Summary List campaign invitations
// @Description List all campaign invitations if the user has GM permissions
// @Tags campaigns
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param campaignID path string true "Campaign ID"
// @Success 200 {array} domain.CampaignInvitation "Campaign invitations retrieved successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid campaign ID"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 403 {object} utils.ErrorResponse "Insufficient permissions"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /api/campaigns/invitations [get]
func (h *CampaignHandler) GetCampaignInvitations(w http.ResponseWriter, r *http.Request) error {
	// Get user ID from context
	userIDStr, ok := r.Context().Value(middleware.UserIDContextKey).(string)
	if !ok {
		return utils.WriteJSONError(w, http.StatusUnauthorized, "User ID not found")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid user ID")
	}

	createdInvitation, err := h.campaignInvitationUseCase.ListCampaignInvitations(userID)
	if err != nil {
		switch {
		case errors.Is(err, usecases.ErrListingCampaignInvitations):
			return utils.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		}

		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(createdInvitation)
}
