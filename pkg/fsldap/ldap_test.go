package fsldap_test

import (
	"encoding/xml"
	"fmt"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/babakyakhchali/go-fsxml/pkg/fsldap"
	"github.com/babakyakhchali/go-fsxml/pkg/fsxml"
	"github.com/spf13/viper"
)

func TestLdap(t *testing.T) {
	wd, _ := os.Getwd()
	fmt.Println("working directory:", wd)
	viper.SetConfigFile("../../.env")
	err := viper.ReadInConfig()
	if err != nil {
		t.Errorf("error reading config file for test")
		return
	}
	fsldap.SetLDAPDefaultTimeout(3 * time.Second)
	ldapConn := fsldap.NewLDAPConnection(fsldap.LDAPConfig{
		Username: viper.GetString("LDAP_USERNAME"),
		Password: viper.GetString("LDAP_PASSWORD"),
		BaseDN:   viper.GetString("LDAP_BASE_DN"),
		URL:      viper.GetString("LDAP_URL"),
	})

	err = ldapConn.Connect()
	if err != nil {
		t.Errorf("error connecting, err:%s", err)
		return
	}
	userAttributes := []string{"distinguishedName", "sAMAccountName", "userPrincipalName", "memberOf", "displayName"}

	users, err := ldapConn.GetActiveUsers(userAttributes)
	if err != nil {
		t.Errorf("error searching users, err:%s", err)
		return
	}
	for _, entry := range users.Entries {
		entry.PrettyPrint(2)
	}

	groups, err := ldapConn.GetGroups()
	if err != nil {
		t.Errorf("error searching users, err:%s", err)
		return
	}
	//users.Entries, groups.Entries, "default", "1234"
	domain := fsldap.CreateDomain(users.Entries, groups.Entries, fsldap.CreateDomainOptions{
		DomainName:      "default",
		DefaultPassword: "1234",
	})
	//fmt.Printf("domain is %v", domain)

	section := fsxml.Section{Name: "directory", Domain: []fsxml.Domain{domain}}

	out, _ := xml.MarshalIndent(section, " ", "  ")
	r := regexp.MustCompile("></[a-zA-Z0-9]*>")
	ns := r.ReplaceAllString(string(out), "/>")
	fmt.Println(ns)
}
