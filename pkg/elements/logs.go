package elements

import "time"

type Logs struct {
	Xmlns               string          `xml:"xmlns,attr"`
	XmlnsXSI            string          `xml:"xsi,attr"`
	XmlnsSchemaLocation string          `xml:"schemaLocation,attr"`
	Version             string          `xml:"version,attr"`
	DocumentInfo        []*DocumentInfo `xml:"documentInfo"`
	Logs                []*Log          `xml:"log"`
}

type DocumentInfo struct {
	DocumentName string `xml:"DocumentName"`
}

type Log struct {
	UIDWell        string         `xml:"uidWell,attr"`
	UIDWellbore    string         `xml:"uidWellbore,attr"`
	UID            string         `xml:"uid,attr"`
	NameWell       string         `xml:"nameWell"`
	NameWellbore   string         `xml:"nameWellbore"`
	Name           string         `xml:"name"`
	ServiceCompany string         `xml:"serviceCompany"`
	RunNumber      string         `xml:"runNumber"`
	CreationDate   string         `xml:"creationDate"`
	Description    string         `xml:"description"`
	IndexType      string         `xml:"indexType"`
	StartIndex     Index          `xml:"startIndex"`
	EndIndex       Index          `xml:"endIndex"`
	StepIncrement  string         `xml:"stepIncrement"`
	Direction      string         `xml:"direction"`
	IndexCurve     IndexCurve     `xml:"indexCurve"`
	NullValue      string         `xml:"nullValue"`
	LogParam       []LogParam     `xml:"logParam"`
	LogCurveInfo   []LogCurveInfo `xml:"logCurveInfo"`
	LogData        []string       `xml:"logData"`
	CommonData     []CommonData   `xml:"commonData"`
}

type Index struct {
	UOM   string  `xml:"uom,attr"`
	Index float32 `xml:",chardata"`
}

type IndexCurve struct {
	Index int    `xml:"columnIndex,attr"`
	Curve string `xml:",chardata"`
}

type LogParam struct {
	Index       string `xml:"index,attr"`
	Name        string `xml:"name,attr"`
	UOM         string `xml:"uom,attr"`
	Description string `xml:"description,attr"`
	Param       string `xml:",chardata"`
}

type LogCurveInfo struct {
	UID              string    `xml:"uid,attr"`
	Mnemonic         string    `xml:"mnemonic"`
	ClassWITSML      string    `xml:"classWitsml"`
	UNIT             string    `xml:"unit"`
	MnemAlias        string    `xml:"mnemAlias"`
	NullValue        string    `xml:"nullValue"`
	MinIndex         Index     `xml:"minIndex"`
	MaxIndex         Index     `xml:"maxIndex"`
	MinDateTimeIndex time.Time `xml:"minDateTimeIndex"`
	MaxDateTimeIndex time.Time `xml:"maxDateTimeIndex"`
	ColumnIndex      int       `xml:"columnIndex"`
	CurveDescription string    `xml:"curveDescription"`
	SensorOffset     string    `xml:"sensorOffset"`
	TraceState       string    `xml:"traceState"`
	TypeLogData      string    `xml:"typeLogData"`
}

type CommonData struct {
	DTimCreation   string `xml:"dTimCreation"`
	DTimLastChange string `xml:"dTimLastChange"`
	ItemState      string `xml:"itemState"`
	Comments       string `xml:"comments"`
}
