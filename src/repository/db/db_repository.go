package db

import (
	"github.com/PS-07/bookstore_oauth-api/src/clients/cassandra"
	"github.com/PS-07/bookstore_oauth-api/src/domain/accesstoken"
	"github.com/PS-07/bookstore_oauth-api/src/utils/errors/errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (? ,? ,? ,?);"
	queryUpdateExpires = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

// DatabaseRepository interface
type DatabaseRepository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestErr)
	Create(accesstoken.AccessToken) *errors.RestErr
	UpdateExpirationTime(accesstoken.AccessToken) *errors.RestErr
}

type databaseRepository struct{}

// NewRepository func
func NewRepository() DatabaseRepository {
	return &databaseRepository{}
}

func (repo *databaseRepository) GetByID(id string) (*accesstoken.AccessToken, *errors.RestErr) {var result accesstoken.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(&result.AccessToken, &result.UserID, &result.ClientID, &result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("no access token found with given id")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &result, nil
}

func (repo *databaseRepository) Create(at accesstoken.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(queryCreateAccessToken, at.AccessToken, at.UserID, at.ClientID, at.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (repo *databaseRepository) UpdateExpirationTime(at accesstoken.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(queryUpdateExpires, at.Expires, at.AccessToken).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
