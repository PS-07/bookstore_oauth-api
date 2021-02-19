package rest

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/PS-07/bookstore_oauth-api/src/domain/users"
	"github.com/PS-07/bookstore_utils-go/resterrors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8082",
		Timeout: 100 * time.Millisecond,
	}
)

// RestUsersRepository interface
type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, resterrors.RestErr)
}

type usersRepository struct{}

// NewRestUsersRepository func
func NewRestUsersRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, resterrors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	response := usersRestClient.Post("/users/login", request)

	if response == nil || response.Response == nil {
		return nil, resterrors.NewInternalServerError("invalid restclient response when trying to login user", errors.New("restclient error"))
	}

	if response.StatusCode > 299 {
		apiErr, err := resterrors.NewRestErrorFromBytes(response.Bytes())
		if err != nil {
			return nil, resterrors.NewInternalServerError("invalid error interface when trying to login user", err)
		}
		return nil, apiErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, resterrors.NewInternalServerError("error when trying to unmarshal users login response", errors.New("json parsing error"))
	}
	return &user, nil
}
