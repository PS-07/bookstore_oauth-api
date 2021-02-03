package http

import (
	"net/http"

	"github.com/PS-07/bookstore_oauth-api/src/domain/accesstoken"
	"github.com/PS-07/bookstore_oauth-api/src/utils/errors/errors"
	"github.com/gin-gonic/gin"
)

// AccessTokenHandler interface
type AccessTokenHandler interface {
	GetByID(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	service accesstoken.Service
}

// NewHandler func
func NewHandler(service accesstoken.Service) AccessTokenHandler {
	return &accessTokenHandler{service: service}
}

// GetByID func
func (handler *accessTokenHandler) GetByID(c *gin.Context) {
	accessTokenID := c.Param("access_token_id")
	accessToken, err := handler.service.GetByID(accessTokenID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

// Create func
func (handler *accessTokenHandler) Create(c *gin.Context) {
	var at accesstoken.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	if err := handler.service.Create(at); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, at)
}
