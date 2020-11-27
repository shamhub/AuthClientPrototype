package platform

import (
	"errors"
	"math/rand"
	"time"

	"github.com/shamhub/codingtest/internal/data"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // mock logic to call auth source
}

//
//
//
// Authenticator represents a contract
// Authenticator performs series of exchanges with the server until the serer says "permitted" or "denied"
type Authenticator interface {
	AuthenticationStep(in data.InputAuthenData) (out data.ServerResponseData, complete bool, err error)
}

//
//
//
// DBBasedAuthentication, a concrete type that implements Authenticator
// DBBasedAuthentication is a DB authentication client for DB based authentication.
type DBBasedAuthentication struct {
	Host         string
	DatabaseName string
	port         int
}

// AuthenticationStep Validate the credentials for DB based authentication
// Implementation of Authenticator interface
func (authenticator *DBBasedAuthentication) AuthenticationStep(d data.InputAuthenData) (out data.ServerResponseData, complete bool, err error) {
	switch rand.Intn(10) { // mock logic
	case 0, 1, 2, 3, 4:
		// Exchange with server and provide response
		return &data.DBServerResponseData{}, true, nil

	case 5, 6, 7, 8, 9:
		// Exchange with server and provide response
		return &data.DBServerResponseData{}, false, errors.New("something is wrong")

	default:
		// Exchange with server and provide response
		return &data.DBServerResponseData{}, false, errors.New("unknown error")
	}
}

//
//
//
//
//
// LDAPBasedAuthentication is a LDAP authentication client for LDAP based authentication.
type LDAPBasedAuthentication struct {
	Host     string
	port     int
	acs      string
	whatever int
}

// AuthenticationStep Validate the credentials for LDAP based authentication
// Implementation of Authenticator interface
func (authenticator *LDAPBasedAuthentication) AuthenticationStep(d data.InputAuthenData) (out data.ServerResponseData, complete bool, err error) {
	switch rand.Intn(10) { // mock logic
	case 0, 1, 2, 3, 4:
		// Exchange with server and provide response
		return &data.LDAPServerResponseData{}, true, nil

	case 5, 6, 7, 8, 9:
		// Exchange with server and provide response
		return &data.LDAPServerResponseData{}, false, errors.New("something is wrong")

	default:
		// Exchange with server and provide response
		return &data.LDAPServerResponseData{}, false, errors.New("unknown error")
	}
}

//
//
//
//
//
//
//
// TokenBasedAuthentication is a token based authentication client for KeyStore server
type TokenBasedAuthentication struct {
	apiToken string
	expiry   int
	host     string
	whatever int
}

// AuthenticationStep Validate the credentials for LDAP based authentication
// Implementation of Authenticator interface
func (authenticator *TokenBasedAuthentication) AuthenticationStep(d data.InputAuthenData) (out data.ServerResponseData, complete bool, err error) {
	switch rand.Intn(10) { // mock logic
	case 0, 1, 2, 3, 4: // mock logic - just a filler code
		// Exchange with server and provide some response status
		return &data.TokenBasedServerResponseData{}, true, nil

	case 5, 6, 7, 8, 9: // mock logic - just a filler code
		// Exchange with server and provide some response status
		return &data.TokenBasedServerResponseData{}, false, errors.New("something is wrong")

	default: // mock logic - just a filler code
		// Exchange with server and provide some response status
		return &data.TokenBasedServerResponseData{}, false, errors.New("unknown error")
	}

}
