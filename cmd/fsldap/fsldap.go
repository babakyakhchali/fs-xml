package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/babakyakhchali/go-fsxml/pkg/fsldap"
)

func main() {
	section := flag.String("section", "directory", "freeswitch xml configuration section type")
	url := flag.String("url", "ldap://127.0.0.1:389", "ldap server url")
	username := flag.String("username", "ldap://127.0.0.1:389", "username for ldap login")
	password := flag.String("password", "", "password for ldap login")
	defaultPassword := flag.String("defaultPassword", "", "password for ldap login")
	usernameAttr := flag.String("usernameAttr", "", "ldap attribute name used to get sip username from ldap entity")
	passwordAttr := flag.String("passwordAttr", "", "ldap attribute name used to get sip password from ldap entity")
	domainName := flag.String("domainName", "default", "freeswitch directory domain name")
	baseDN := flag.String("dn", "example.com", "base dn")
	outfile := flag.String("o", "", "save output to this file")
	timeout := flag.Duration("t", 3, "ldap default timeout")

	flag.Parse()

	if timeout != nil {
		fsldap.SetLDAPDefaultTimeout(*timeout * time.Second)
	}

	switch *section {
	case "directory":
		s, e := fsldap.CreateDomainFromLDAP(fsldap.CreateDomainFromLDAPOpts{
			LdapConfig: fsldap.LDAPConfig{
				Username: *username,
				Password: *password,
				BaseDN:   *baseDN,
				URL:      *url,
			},
			CreateDomainOptions: fsldap.CreateDomainOptions{
				DomainName:         *domainName,
				DefaultPassword:    *defaultPassword,
				SipUsernameField:   *usernameAttr,
				SipPasswordField:   *passwordAttr,
				ExtraFieldMappings: map[string]string{}, //TODO:
			},
		})
		if e != nil {
			fmt.Println("error:", e)
		}
		if *outfile != "" {
			os.WriteFile(*outfile, []byte(s), 0644)
		} else {
			fmt.Println("result:", s)
		}

		return
	default:
		fmt.Println("not suppoert section:", *section)

	}

}
