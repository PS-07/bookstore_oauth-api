package accesstoken

import (
	"fmt"
	"strings"
	"time"

	"github.com/PS-07/bookstore_oauth-api/src/utils/cryptoutils"
	"github.com/PS-07/bookstore_utils-go/resterrors"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grandTypeClientCredentials = "client_credentials"
)

// AccessTokenRequest struct
type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// Used for password grant type
	Username string `json:"username"`
	Password string `json:"password"`

	// Used for client_credentials grant type
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// Validate func
func (at *AccessTokenRequest) Validate() resterrors.RestErr {
	switch at.GrantType {
	case grantTypePassword:
		break

	case grandTypeClientCredentials:
		break

	default:
		return resterrors.NewBadRequestError("invalid grant_type parameter")
	}

	//TODO: Validate parameters for each grant_type
	return nil
}

// AccessToken struct
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id,omitempty"`
	Expires     int64  `json:"expires"`
}

// Validate func
func (at *AccessToken) Validate() resterrors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return resterrors.NewBadRequestError("invalid access token id")
	}
	if at.UserID <= 0 {
		return resterrors.NewBadRequestError("invalid user id")
	}
	if at.ClientID <= 0 {
		return resterrors.NewBadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return resterrors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

// GetNewAccessToken func
func GetNewAccessToken(userID int64) AccessToken {
	return AccessToken{
		UserID:  userID,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

// IsExpired func
func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

// Generate func
func (at *AccessToken) Generate() {
	at.AccessToken = cryptoutils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserID, at.Expires))
}
