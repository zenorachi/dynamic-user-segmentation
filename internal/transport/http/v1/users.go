package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zenorachi/dynamic-user-segmentation/internal/entity"
	"net/http"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.POST("/sign-up", h.signUp)
		users.POST("/sign-in", h.signIn)
		users.GET("/refresh", h.refresh)
	}
	{
		users.GET("/active_segments/", h.userIdentity, h.getActiveSegments)
	}
}

type signUpInput struct {
	Login    string `json:"login"    binding:"required,min=2,max=64"`
	Email    string `json:"email"    binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

type signUpResponse struct {
	ID int `json:"id"`
}

// @Summary User SignUp
// @Description create user account
// @Tags auth
// @Accept json
// @Produce json
// @Param input body signUpInput true "input"
// @Success 201 {object} signUpResponse
// @Failure 400 {object} errorResponse
// @Failure 409 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v1/users/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input signUpInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidInput.Error())
		return
	}

	id, err := h.services.Users.SignUp(c, input.Login, input.Email, input.Password)
	if err != nil {
		if errors.Is(err, entity.ErrUserAlreadyExists) {
			newErrorResponse(c, http.StatusConflict, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	newResponse(c, http.StatusCreated, signUpResponse{ID: id})
}

type signInInput struct {
	Login    string `json:"login"    binding:"required,min=2,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

type tokenResponse struct {
	Token string `json:"token"`
}

// @Summary User SignIn
// @Description user sign in
// @Tags auth
// @Accept json
// @Produce json
// @Param input body signInInput true "input"
// @Success 200 {object} tokenResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v1/users/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidInput.Error())
		return
	}

	tokens, err := h.services.Users.SignIn(c, input.Login, input.Password)
	if err != nil {
		if errors.Is(err, entity.ErrUserDoesNotExist) || errors.Is(err, entity.ErrIncorrectPassword) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.Header("Set-Cookie", fmt.Sprintf("refresh-token=%s; HttpOnly", tokens.RefreshToken))
	newResponse(c, http.StatusOK, tokenResponse{Token: tokens.AccessToken})
}

// @Summary User Refresh Token
// @Description refresh user's access token
// @Tags auth
// @Produce json
// @HeaderParam Set-Cookie string true "RefreshToken"
// @Success 200 {object} tokenResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v1/users/refresh [get]
func (h *Handler) refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh-token")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "refresh-token not found")
		return
	}

	tokens, err := h.services.RefreshTokens(c, refreshToken)
	if err != nil {
		if errors.Is(err, entity.ErrSessionDoesNotExist) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.Header("Set-Cookie", fmt.Sprintf("refresh-token=%s; HttpOnly", tokens.RefreshToken))
	newResponse(c, http.StatusOK, tokenResponse{Token: tokens.AccessToken})
}

type getActiveSegmentsInput struct {
	UserID int `json:"user_id"`
}

type getActiveSegmentsResponse struct {
	Segments []entity.Segment `json:"segments"`
}

// @Summary Get active segments for a user
// @Security Bearer
// @Description get active segments for a specific user by user_id
// @Tags user-segments
// @Accept json
// @Produce json
// @Param input body getActiveSegmentsInput true "input"
// @Success 200 {object} getActiveSegmentsResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v1/users/active_segments/ [get]
func (h *Handler) getActiveSegments(c *gin.Context) {
	var input getActiveSegmentsInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, entity.ErrInvalidInput.Error())
		return
	}

	segments, err := h.services.Users.GetActiveSegmentsByUserID(c, input.UserID)
	if err != nil {
		if errors.Is(err, entity.ErrUserDoesNotExist) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	newResponse(c, http.StatusOK, getActiveSegmentsResponse{Segments: segments})
}
