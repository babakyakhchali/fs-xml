package fsxml_test

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"testing"

	"github.com/babakyakhchali/go-fsxml/pkg/fsxml"
)

// fsxml.Param{Name: "password", Value: "1234"}
func TestToXML(t *testing.T) {
	user1000 := fsxml.User{
		ID:     "1000",
		Params: []fsxml.Param{{Name: "password", Value: "1234"}, {Name: "vm-password", Value: "1234"}},
	}

	user1001 := fsxml.User{
		ID:     "1001",
		Params: []fsxml.Param{{Name: "password", Value: "1234"}, {Name: "vm-password", Value: "1234"}},
	}

	group := fsxml.Group{Name: "default", Users: []fsxml.User{user1000, user1001}}

	domain := fsxml.Domain{Name: "$${local_ip_v4}", Groups: []fsxml.Group{group}}

	section := fsxml.Section{Name: "directory", Domain: []fsxml.Domain{domain}}

	out, _ := xml.MarshalIndent(section, " ", "  ")
	r := regexp.MustCompile("></[a-zA-Z0-9]*>")
	ns := r.ReplaceAllString(string(out), "/>")
	fmt.Println(ns)
}

var testXML = `
<section name="directory">
    <domain name="$${internal_ip}">
      <params>
        <param name="dial-string" value="${sofia_contact(*/${dialed_user}@${dialed_domain})}"/>
      </params>
      <groups>
        <group name="local">
          <users>
            <user id="1001">
              <params>
                <param name="password" value="1001"/>
              </params>
            </user>
            <user id="1002">
              <params>
                <param name="password" value="1002"/>
              </params>
            </user>
            <user id="1003">
              <params>
                <param name="password" value="1003"/>
              </params>
            </user>
          </users>
        </group>
      </groups>
    </domain>
  </section>
`

func TestFromXML(t *testing.T) {
	section := fsxml.Section{}
	err := xml.Unmarshal([]byte(testXML), &section)
	if err != nil {
		t.Fatalf("failed to unmarshal, error:%s", err)
		return
	}
}
