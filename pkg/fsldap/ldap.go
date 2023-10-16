package fsldap

import (
	"fmt"
	"time"

	"github.com/go-ldap/ldap"
)

const (
	BindUsername = "voip@rayanehkomak.com"
	BindPassword = "Voipqaz22"
	BaseDN       = "dc=rayanehkomak, dc=com"
)

type LDAPConfig struct {
	Username string
	Password string
	BaseDN   string
	URL      string
}

func SetLDAPDefaultTimeout(d time.Duration) {
	ldap.DefaultTimeout = d
}

type LDAPConnection struct {
	config LDAPConfig
	conn   *ldap.Conn
}

func NewLDAPConnection(config LDAPConfig) LDAPConnection {
	return LDAPConnection{config: config}
}

func (con *LDAPConnection) Connect() error {
	conn, err := ldap.DialURL(con.config.URL)
	if err != nil {
		return err
	}
	con.conn = conn
	return nil
}

func (con *LDAPConnection) GetUserByUsername(username string, attributes []string) (*ldap.SearchResult, error) {
	filter := fmt.Sprintf("(sAMAccountName=%s)", username)

	return con.BindAndSearch(filter, attributes)
}

func (con *LDAPConnection) GetUsers(attributes []string) (*ldap.SearchResult, error) {
	filter := "(&(objectClass=user)(!(objectClass=computer)))"

	return con.BindAndSearch(filter, attributes)
}

func (con *LDAPConnection) GetActiveUsers(attributes []string) (*ldap.SearchResult, error) {
	filter := "(&(objectClass=user)(!(objectClass=computer))(!(userAccountControl:1.2.840.113556.1.4.803:=2)))"

	return con.BindAndSearch(filter, attributes)
}

func (con *LDAPConnection) GetGroups() (*ldap.SearchResult, error) {
	filter := "(&(objectClass=group)(!(objectClass=computer)))"
	return con.BindAndSearch(filter, []string{})
}

func (con *LDAPConnection) BindAndSearch(filter string, attributes []string) (*ldap.SearchResult, error) {
	err := con.conn.Bind(BindUsername, BindPassword)
	if err != nil {
		return nil, fmt.Errorf("bind error: %s", err)
	}
	searchReq := ldap.NewSearchRequest(
		BaseDN,
		ldap.ScopeWholeSubtree, // you can also use ldap.ScopeWholeSubtree
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		filter,
		attributes,
		nil,
	)
	return con.conn.Search(searchReq)

}
