package fsxml

import "encoding/xml"

type FreeswitchDocument struct {
	XMLName       xml.Name `xml:"document" json:"document,omitempty"`
	XPreProcesses []XPreProcess
	Sections      []Section //`xml:"section" json:"section,omitempty"`
	Type          string    `xml:"type,attr" json:"type,omitempty"`
}

type Configuration struct {
	XMLName  xml.Name  `xml:"configuration" json:"configuration,omitempty"`
	Name     string    `xml:"name,attr" json:"name,omitempty"`
	Settings *Settings //`xml:"settings" json:"settings,omitempty"`
}

type Context struct {
	XMLName xml.Name `xml:"context" json:"context,omitempty"`
	Name    string   `xml:"name,attr" json:"name,omitempty"`
}

type Section struct {
	XMLName        xml.Name         `xml:"section" json:"section,omitempty"`
	Text           string           `xml:",chardata" json:"text,omitempty"`
	Name           string           `xml:"name,attr" json:"name,omitempty"`
	Domains        *[]Domain        //`xml:"domain" json:"domain,omitempty"`
	Configurations *[]Configuration //`xml:"configuration" json:"configuration,omitempty"`
	Contexts       *[]Context       `xml:"context" json:"context,omitempty"`
}

type Settings struct {
	XMLName xml.Name `xml:"settings" json:"settings,omitempty"`
	Params  []Param  `xml:"params" json:"params,omitempty"`
}

type Variables struct {
	XMLName   xml.Name   `xml:"variables" json:"variables,variables"`
	Variables []Variable `xml:"variable" json:"variable,variable"`
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

// <X-PRE-PROCESS cmd="set" data="internal_ip=127.0.0.1"/>
type XPreProcess struct {
	XMLName xml.Name `xml:"X-PRE-PROCESS" json:"X-PRE-PROCESS,omitempty"`
	Cmd     string   `xml:"cmd,attr" json:"cmd,omitempty"`
	Data    string   `xml:"data,attr" json:"data,omitempty"`
}
