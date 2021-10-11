package access_token

import (
	"bookstoreapi/oauth/utils/errors"
	"bookstoreapi/users/utils/crypto_utils"
	"fmt"
	"strings"
	"time"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id,omitempty"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)
	return expirationTime.Before(now)
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}

/*
0: Every check
1: Just token
*/

func validateInts(fieldVal int64, fieldName string) *errors.RestErr {
	if fieldVal <= 0 {
		return errors.NewBadRequestError(fmt.Sprintf("invalid data %s", fieldName))
	}
	return nil
}

func (at *AccessToken) Validate(instructs map[int]bool) *errors.RestErr {
	switch {
	case instructs[0]:
		err := validateInts(at.ClientId, "clientId")
		if err != nil {
			return err
		}
		err = validateInts(at.UserId, "clientId")
		if err != nil {
			return err
		}
		err = validateInts(at.Expires, "expires data")
		if err != nil {
			return err
		}
		fallthrough
	case instructs[1]:
		at.AccessToken = strings.TrimSpace(at.AccessToken)
		if len(at.AccessToken) == 0 {
			return errors.NewBadRequestError("invalid access token")
		}
	}

	return nil
}

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`
	// GrantType Password
	Username string `json:"email"`
	Password string `json:"password"`
	// GrantType client_Credentials
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (atr *AccessTokenRequest) Validate(instructs map[int]bool) *errors.RestErr {
	switch atr.GrantType {
	//TODO: Validate parameters for each grant_type
	case grantTypePassword:
		break
	case grantTypeClientCredentials:
		break
	default:
		return errors.NewBadRequestError("Token request invalid, invalid grant_type")
	}
	return nil
}

//Web frontend: Client-Id: 123
//Android APP: Client-Id: 234
