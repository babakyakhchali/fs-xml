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

	directory := fsxml.Section{Name: "directory", Domains: &[]fsxml.Domain{domain}}

	configuration := fsxml.Section{Name: "configuration", Configurations: &[]fsxml.Configuration{
		{
			Name: "event_socket.conf",
			Settings: &fsxml.Settings{
				Params: []fsxml.Param{
					{Name: "listen-ip", Value: "127.0.0.1"},
				},
			},
		},
	}}

	dialplan := fsxml.Section{Name: "dialplan"}

	fsdoc := fsxml.FreeswitchDocument{
		XPreProcesses: []fsxml.XPreProcess{{Cmd: "set", Data: "internal_ip=127.0.0.1"}},
		Sections:      []fsxml.Section{directory, configuration, dialplan}, Type: "freeswitch/xml"}

	out, err := xml.MarshalIndent(fsdoc, " ", "  ")
	if err != nil {
		t.Fatalf("failed to unmarshal, error:%s", err)
		return
	}
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
