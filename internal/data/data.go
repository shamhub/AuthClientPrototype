package data

//	input authentication data for clients using username & password based authentication

//	DBClientAuthenticationData
//	&
//	LDAPClientAuthenticationData
//	&
//	ClientAuthenData are open for modification

// InputAuthenData provides interface for an input to multiple authentication implementations
type InputAuthenData interface {
	IsValidData() bool
	// whateverbehavior(string)string
}

// DBInputAuthenData resembles the input sent to DB server
type DBInputAuthenData struct {
	username string
	password string
	port     int
	tenancy  int
	whatever int
}

// IsValidData validates the data before sending data to DB server
func (c *DBInputAuthenData) IsValidData() bool {
	return false
	// parse the username & password string for standard compliance
}

// LDAPInputAuthenData resembles the input sent to LDAP server
type LDAPInputAuthenData struct {
	username string
	password string
	whatever int
}

// IsValidData validates the data before sending data to LDAP server
func (c *LDAPInputAuthenData) IsValidData() bool {

	// parse the username & password string for standard compliance
	return true
}

// TokenBasedInputAuthenData resembles the input sent to DB server
type TokenBasedInputAuthenData struct {
	apiToken string
	duration int
	whatever int
}

// IsValidData validates the data before sending data to DB server
func (c *TokenBasedInputAuthenData) IsValidData() bool {
	return false
	// parse the username & password string for standard compliance
}

//	Response from LDAP or Database after validating credentials

//	ServerResponseData
//	&
//	DBServerResponseData
//	&
//	LDAPServerResponseData are open for modification

// ServerResponseData provides interface to ouput given by multiple authentication impleemntations
type ServerResponseData interface {
	ReadResponse() string
}

// DBServerResponseData holds the response received from DB server
type DBServerResponseData struct {
	responseCode int
	responseData string
	whatever     int
}

// ReadResponse interprets and provide readable response
func (c *DBServerResponseData) ReadResponse() string {

	// parse the response
	return "this is the response"
}

// LDAPServerResponseData holds the response received from LDAP server
type LDAPServerResponseData struct {
	responseData string
	whatever     int
}

// ReadResponse interprets and provide readable response
func (c *LDAPServerResponseData) ReadResponse() string {

	// parse the response
	return "this is the response"
}

// TokenBasedServerResponseData holds the response received from LDAP server
type TokenBasedServerResponseData struct {
	responseData string
	whatever     int
}

// ReadResponse interprets and provide readable response
func (c *TokenBasedServerResponseData) ReadResponse() string {

	// parse the response
	return "this is the response"
}
