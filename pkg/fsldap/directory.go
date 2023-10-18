package fsldap

import (
	"encoding/xml"
	"regexp"

	"github.com/babakyakhchali/go-fsxml/pkg/fsxml"
	"github.com/go-ldap/ldap"
)

func addUserToGroup(domain *fsxml.Domain, groupName string, user fsxml.User, isPointer bool) {
	foundGroupIDX := -1
	for i := 0; i < len(domain.Groups); i++ {
		if domain.Groups[i].Name == groupName {
			foundGroupIDX = i
			break
		}
	}

	if foundGroupIDX < 0 {
		if isPointer {
			domain.Groups = append(domain.Groups, fsxml.Group{Name: groupName, Users: []fsxml.User{{ID: user.ID, Type: "pointer"}}})
		} else {
			domain.Groups = append(domain.Groups, fsxml.Group{Name: groupName, Users: []fsxml.User{user}})
		}

	} else {
		if isPointer {
			domain.Groups[foundGroupIDX].Users = append(domain.Groups[foundGroupIDX].Users, fsxml.User{ID: user.ID, Type: "pointer"})
		} else {
			domain.Groups[foundGroupIDX].Users = append(domain.Groups[foundGroupIDX].Users, user)
		}
	}
}

type CreateDomainOptions struct {
	DomainName      string
	DefaultPassword string

	SipUsernameField   string
	SipPasswordField   string
	ExtraFieldMappings map[string]string
	DomainVars         map[string]string
	DomainParams       map[string]string
	UserVars           map[string]string
}

func CreateDomain(users []*ldap.Entry, groups []*ldap.Entry, opts CreateDomainOptions) fsxml.Domain {
	domain := fsxml.Domain{Name: opts.DomainName, Groups: []fsxml.Group{}}
	for name, value := range opts.DomainParams {
		domain.Variables = append(domain.Variables, fsxml.Variable{Name: name, Value: value})
	}
	for name, value := range opts.DomainParams {
		domain.Params = append(domain.Params, fsxml.Param{Name: name, Value: value})
	}
	sipPasswordAttr := ""
	if opts.SipPasswordField != "" {
		sipPasswordAttr = opts.SipPasswordField
	}

	sipUsernameAttr := "sAMAccountName"
	if opts.SipPasswordField != "" {
		sipUsernameAttr = opts.SipPasswordField
	}

	for _, ldapUser := range users {
		xmlUser := fsxml.User{}
		xmlUser.ID = ldapUser.GetAttributeValue(sipUsernameAttr)
		password := ""
		if sipPasswordAttr != "" {
			password = ldapUser.GetAttributeValue(sipPasswordAttr)
		}
		if password == "" {
			password = opts.DefaultPassword
		}
		xmlUser.Params = []fsxml.Param{
			{Name: "password", Value: password},
			{Name: "vm-password", Value: password}}
		for varName, varAttr := range opts.ExtraFieldMappings {
			xmlUser.Variables = append(xmlUser.Variables, fsxml.Variable{Name: varName, Value: ldapUser.GetAttributeValue(varAttr)})
		}
		xmlUser.Variables = append(xmlUser.Variables, fsxml.Variable{Name: "source", Value: "ldap"})
		addUserToGroup(&domain, "default", xmlUser, false)
		userLdapGroups := ldapUser.GetAttributeValues("memberOf")
		for _, gs := range userLdapGroups {
			for _, g := range groups {
				if g.GetAttributeValue("distinguishedName") == gs {
					addUserToGroup(&domain, g.GetAttributeValue("name"), xmlUser, true)
					break
				}
			}
		}
	}
	return domain
}

type CreateDomainFromLDAPOpts struct {
	LdapConfig          LDAPConfig
	CreateDomainOptions CreateDomainOptions
}

func CreateDomainFromLDAP(opts CreateDomainFromLDAPOpts) (xmlstring string, err error) {
	ldapConn := NewLDAPConnection(opts.LdapConfig)
	err = ldapConn.Connect()
	if err != nil {
		return
	}
	userAttributes := []string{"distinguishedName", "sAMAccountName", "userPrincipalName", "memberOf", "displayName"}

	for _, v := range opts.CreateDomainOptions.ExtraFieldMappings {
		userAttributes = append(userAttributes, v)
	}

	users, err := ldapConn.GetActiveUsers(userAttributes)
	if err != nil {
		return
	}

	groups, err := ldapConn.GetGroups()
	if err != nil {
		return
	}

	domain := CreateDomain(users.Entries, groups.Entries, opts.CreateDomainOptions)

	section := fsxml.Section{Name: "directory", Domains: &[]fsxml.Domain{domain}}

	out, _ := xml.MarshalIndent(section, " ", "  ")
	r := regexp.MustCompile("></[a-zA-Z0-9]*>")
	xmlstring = r.ReplaceAllString(string(out), "/>")
	return
}
