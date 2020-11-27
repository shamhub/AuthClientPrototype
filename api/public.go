package api

import (
	"errors"

	"github.com/shamhub/codingtest/internal/data"
	"github.com/shamhub/codingtest/internal/platform"
)

// AuthenticationType - the client inevitably knows which "kind" of authentication takes place
type AuthenticationType string

const ( // open for modification
	// PasswordBasedAuthenType helps client to pick DB or LDAP based implementation
	PasswordBasedAuthenType AuthenticationType = "password-based"
	// TokenBasedAuthenType helps client to pick token based(temporary) authentication
	TokenBasedAuthenType AuthenticationType = "token-based"
)

//	Client authentication data for clients to provide credentials

//	PasswordBasedAuthenData type
//	&
//	TokenBasedClientAuthenticationData type
//	&
//	ClientAuthenData interface  - open for modification
//
//
//

// ClientAuthenData provides interface for an input to multiple authentication implementations
type ClientAuthenData interface { // open for modification
	IsClientDataValid() bool
}

// PasswordBasedClientAuthenData resembles the client input given for password based authentication(LDAP or DB)
type PasswordBasedClientAuthenData struct { // open for modification
	Username string
	Password string
	Whatever int
}

// IsClientDataValid api to validate client data
func (c *PasswordBasedClientAuthenData) IsClientDataValid() bool {
	return false
	// parse the username & password string for standard compliance
}

// TokenBasedClientAuthenData resembles the client input given for token based authentication
type TokenBasedClientAuthenData struct {
	ApiToken string
	Expiry   int
	Whatever int
}

// IsClientDataValid api to validate client data
func (c *TokenBasedClientAuthenData) IsClientDataValid() bool {
	return false
	// parse the username & password string for standard compliance
}

// Response provides unified response to user for any authentication type
type Response struct {
	ResponseCode int
	ResponseData string
	Whatever     int
}

// ReadResponse interprets response data and provide readable response
func (r *Response) ReadResponse() string {

	// parse the response
	return "this is the response"
}

// Authenticate - For AuthenticationType "password-based", Authenticate() receives `PasswordBasedClientAuthenData` type data as second argument
// Authenticate - For AuthenticationType "token-based", Authenticate() receives `TokenBasedClientAuthenData` type data as second argument
//
//
//
//
//
//
//
//
//
// Authenticate() provides api to client to provide authentication type & corresponding input data
func Authenticate(kindOfAuthentication AuthenticationType, authenInputData ClientAuthenData) (*Response, error) {

	if !authenInputData.IsClientDataValid() {
		return nil, errors.New("Invalid data")
	}

	var authenticator platform.Authenticator
	switch kindOfAuthentication { // code smell - because client inevitably must know which "kind" of authentication takes place
	case PasswordBasedAuthenType:
		//
		//
		// Firstly try DB authentication
		// DB authentication
		authenticator = &platform.DBBasedAuthentication{}
		serverResponse, permitted, err := authenticator.AuthenticationStep(&data.DBInputAuthenData{})
		if permitted {
			return &Response{1234, serverResponse.ReadResponse(), 01}, err
		}
		DBserverResponse := serverResponse.ReadResponse()
		//
		//
		//
		//
		// Secondly try LDAP authentication
		// LDAP authentication
		authenticator = &platform.LDAPBasedAuthentication{}
		serverResponse, permitted, err = authenticator.AuthenticationStep(&data.LDAPInputAuthenData{})
		if permitted {
			return &Response{1234, serverResponse.ReadResponse(), 01}, nil
		}
		LDAPServerResponse := serverResponse.ReadResponse()
		return &Response{5678, DBserverResponse + LDAPServerResponse, 02}, err
	case TokenBasedAuthenType:
		//
		//
		//
		// Token based authentication
		authenticator = &platform.TokenBasedAuthentication{}
		serverResponse, permitted, err := authenticator.AuthenticationStep(&data.TokenBasedInputAuthenData{})
		if permitted {
			return &Response{1234, serverResponse.ReadResponse(), 01}, nil
		}
		TokenServerResponse := serverResponse.ReadResponse()
		return &Response{5678, TokenServerResponse, 02}, err
	default:
		//
		//
		//
		return nil, errors.New("Unhandled authentication type")
	}
}
