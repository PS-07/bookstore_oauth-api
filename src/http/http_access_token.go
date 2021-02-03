package http

import (
	"net/http"

	"github.com/PS-07/bookstore_oauth-api/src/domain/accesstoken"
	"github.com/gin-gonic/gin"
)

// AccessTokenHandler interface
type AccessTokenHandler interface {
	GetByID(*gin.Context)
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
