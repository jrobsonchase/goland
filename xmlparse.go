package main

type Protocol struct {
	XMLName    string      `xml:"protocol"`
	Name       string      `xml:"name,attr"`
	Copyright  string      `xml:"copyright"`
	Interfaces []Interface `xml:"interface"`
}

type Interface struct {
	XMLName     string      `xml:"interface"`
	Name        string      `xml:"name,attr"`
	Version     string      `xml:"version,attr"`
	Description Description `xml:"description"`
	Requests    []Request   `xml:"request"`
	Events      []Event     `xml:"event"`
	Enums       []Enum      `xml:"enum"`
}

type Description struct {
	XMLName string `xml:"description"`
	Summary string `xml:"summary,attr"`
	Full    string `xml:",chardata"`
}

type Request struct {
	XMLName     string      `xml:"request"`
	Name        string      `xml:"name,attr"`
	Description Description `xml:"description"`
	Type        string      `xml:"type,attr"`
	Args        []Arg       `xml:"arg"`
}

type Event struct {
	XMLName     string      `xml:"event"`
	Name        string      `xml:"name,attr"`
	Description Description `xml:"description"`
	Type        string      `xml:"type,attr"`
	Args        []Arg       `xml:"arg"`
}

type Arg struct {
	XMLName   string `xml:"arg"`
	Name      string `xml:"name,attr"`
	Type      string `xml:"type,attr"`
	Interface string `xml:"interface,attr"`
	AllowNull bool   `xml:"allow-null,attr"`
}

type Enum struct {
	XMLName     string      `xml:"enum"`
	Name        string      `xml:"name,attr"`
	Description Description `xml:"description"`
	Entries     []EnumEntry `xml:"entry"`
}

type EnumEntry struct {
	XMLName string `xml:"entry"`
	Name    string `xml:"name,attr"`
	Value   string `xml:"value,attr"`
	Summary string `xml:"summary,attr"`
}
