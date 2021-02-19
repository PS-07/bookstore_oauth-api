package db

import (
	"errors"

	"github.com/PS-07/bookstore_oauth-api/src/clients/cassandra"
	"github.com/PS-07/bookstore_oauth-api/src/domain/accesstoken"
	"github.com/PS-07/bookstore_utils-go/resterrors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "SELECT accesstoken, user_id, client_id, expires FROM accesstokens WHERE accesstoken=?;"
	queryCreateAccessToken = "INSERT INTO accesstokens(accesstoken, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpires     = "UPDATE accesstokens SET expires=? WHERE accesstoken=?;"
)

// NewRepository func
func NewRepository() DbRepository {
	return &dbRepository{}
}

// DbRepository interface
type DbRepository interface {
	GetByID(string) (*accesstoken.AccessToken, resterrors.RestErr)
	Create(accesstoken.AccessToken) resterrors.RestErr
	UpdateExpirationTime(accesstoken.AccessToken) resterrors.RestErr
}

type dbRepository struct {
}

// GetByID func
func (r *dbRepository) GetByID(id string) (*accesstoken.AccessToken, resterrors.RestErr) {
	var result accesstoken.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserID,
		&result.ClientID,
		&result.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, resterrors.NewNotFoundError("no access token found with given id")
		}
		return nil, resterrors.NewInternalServerError("error when trying to get current id", errors.New("database error"))
	}
	return &result, nil
}

func (r *dbRepository) Create(at accesstoken.AccessToken) resterrors.RestErr {
	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		at.AccessToken,
		at.UserID,
		at.ClientID,
		at.Expires,
	).Exec(); err != nil {
		return resterrors.NewInternalServerError("error when trying to save access token in database", err)
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(at accesstoken.AccessToken) resterrors.RestErr {
	if err := cassandra.GetSession().Query(queryUpdateExpires,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		return resterrors.NewInternalServerError("error when trying to update current resource", errors.New("database error"))
	}
	return nil
}
