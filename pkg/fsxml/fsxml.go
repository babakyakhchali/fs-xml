package fsxml

import "encoding/xml"

type Section struct {
	XMLName xml.Name `xml:"section" json:"section,omitempty"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
	Name    string   `xml:"name,attr" json:"name,omitempty"`
	Domain  []Domain `xml:"domain" json:"domain,omitempty"`
}

type Domain struct {
	XMLName   xml.Name   `xml:"domain" json:"domain,omitempty"`
	Name      string     `xml:"name,attr" json:"name,omitempty"`
	Params    []Param    `xml:"params>param,omitempty" json:"params,omitempty"`
	Variables []Variable `xml:"variables>variable,omitempty" json:"variables,omitempty"`
	Groups    []Group    `xml:"groups>group" json:"groups,omitempty"`
}

type Group struct {
	XMLName xml.Name `xml:"group" json:"group,omitempty"`
	Name    string   `xml:"name,attr" json:"name,omitempty"`
	Users   []User   `xml:"users>user" json:"users,omitempty"`
}
type User struct {
	XMLName   xml.Name   `xml:"user" json:"user,omitempty"`
	ID        string     `xml:"id,attr" json:"id,omitempty"`
	Params    []Param    `xml:"params>param,omitempty" json:"params,omitempty"`
	Variables []Variable `xml:"variables>variable,omitempty" json:"variables,omitempty"`
	Type      string     `xml:"type,attr,omitempty" json:"type,omitempty"`
}

type Param struct {
	XMLName xml.Name `xml:"param" json:"param,omitempty"`
	Name    string   `xml:"name,attr" json:"name,omitempty"`
	Value   string   `xml:"value,attr" json:"value,omitempty"`
}

type Variable struct {
	XMLName xml.Name `xml:"variable" json:"variable,omitempty"`
	Name    string   `xml:"name,attr" json:"name,omitempty"`
	Value   string   `xml:"value,attr" json:"value,omitempty"`
}
