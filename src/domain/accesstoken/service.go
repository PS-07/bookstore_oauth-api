package accesstoken

import (
	"strings"

	"github.com/PS-07/bookstore_oauth-api/src/utils/errors/errors"
)

// Service interface
type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

// Repository interface
type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

// NewService func
func NewService(repo Repository) Service {
	return &service{repository: repo}
}

// GetByID func
func (srv *service) GetByID(accessTokenID string) (*AccessToken, *errors.RestErr) {
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}

	accessToken, err := srv.repository.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
