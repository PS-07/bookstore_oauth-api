package db

import (
	"github.com/PS-07/bookstore_oauth-api/src/domain/accesstoken"
	"github.com/PS-07/bookstore_oauth-api/src/utils/errors/errors"
)

// DatabaseRepository interface
type DatabaseRepository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestErr)
}

type databaseRepository struct{}

// NewRepository func
func NewRepository() DatabaseRepository {
	return &databaseRepository{}
}

func (repo *databaseRepository) GetByID(id string) (*accesstoken.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("database connection not yet implemented")
}
