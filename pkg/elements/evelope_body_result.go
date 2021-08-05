package elements

import "encoding/xml"

type Envelope struct {
	XMLName xml.Name
	Body    Body
}

type Body struct {
	XMLName     xml.Name
	GetResponse completeResponse `xml:"WMLS_GetFromStoreResponse"`
}

type completeResponse struct {
	Result *Result `xml:"Result"`
	XMLOut XMLOut  `xml:"XMLout"`
}

type Result struct {
	Type   string `xml:"type,attr"`
	Result int    `xml:",chardata"`
}

type XMLOut struct {
	Logs *Logs `xml:"logs"`
}
