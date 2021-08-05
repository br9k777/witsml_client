package elements

import (
	"encoding/xml"
	"fmt"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"strings"
	"time"
)

func ReadXML(interfaceXML interface{}, xmlData []byte) error {
	var err error
	if err = xml.Unmarshal(xmlData, &interfaceXML); err != nil {
		zap.S().Errorf("Fail read from XML:\n%s", string(xmlData))
		return err
	}
	return nil
}

func ReadLogsFromFile(fileXML string) (*Logs, error) {

	OutObj := new(Envelope)
	var err error

	var xmlData []byte
	if xmlData, err = ioutil.ReadFile(fileXML); err != nil {
		zap.S().Errorf("Unable to read %v. Err: %v.", fileXML, err)
		return nil, err
	}
	var logs *Logs
	if err = ReadXML(OutObj, xmlData); err != nil {
		return nil, err
	}
	if OutObj == nil || OutObj.Body.GetResponse.Result == nil {
		logs = new(Logs)
		if err = ReadXML(logs, xmlData); err != nil {
			return nil, err
		}
		return logs, nil
	}
	zap.S().Infof("Result %d", OutObj.Body.GetResponse.Result.Result)
	if OutObj.Body.GetResponse.Result.Result != 1 {
		return nil, fmt.Errorf("%v", OutObj.Body.GetResponse)
	}
	zap.S().Debugf("Out object %#v", OutObj)
	logs = OutObj.Body.GetResponse.XMLOut.Logs

	//logs = new(Logs)
	return logs, nil
}

func (logs *Logs) PrintLogs(writer io.Writer, simple bool) (err error) {
	if logs == nil || len(logs.Logs) == 0 {
		zap.S().Debugf("%#v", logs)
		return nil
	}
	for _, log := range logs.Logs {
		fmt.Printf("uidWell: %-20s\t uidWellbore: %-20s\t uid: %-s\tIndex Type: %s\n", log.UIDWell, log.UIDWellbore, log.UID, log.IndexType)
		if !simple {
			if log.NameWell == "" || log.NameWellbore == "" {
				if _, err = fmt.Fprintf(writer, "Log: %-s\n", log.Name); err != nil {
					return err
				}
			} else {
				if _, err = fmt.Fprintf(writer, "Well: %-20s\t Wellbore: %-20s\t Log: %-s\n", log.NameWell, log.NameWellbore, log.Name); err != nil {
					return err
				}
			}
			if log.ServiceCompany != "" {
				if _, err = fmt.Fprintf(writer, "Service Company: %s\t", log.ServiceCompany); err != nil {
					return err
				}
			}
			if log.RunNumber != "" {
				if _, err = fmt.Fprintf(writer, "Run Number: %s\t", log.RunNumber); err != nil {
					return err
				}
			}
			if log.CreationDate != "" {
				if _, err = fmt.Fprintf(writer, "Creation Date: %s\n", log.CreationDate); err != nil {
					return err
				}
			} else {
				if _, err = fmt.Fprintf(writer, "\n"); err != nil {
					return err
				}
			}
			if log.Description != "" {
				if _, err = fmt.Fprintf(writer, "Description %s\n", log.Description); err != nil {
					return err
				}
			}
			if log.IndexCurve.Curve != "" {
				if _, err = fmt.Fprintf(writer, "IndexCurve: %s\tuom: %s\n", log.IndexCurve.Curve, log.IndexCurve.Curve); err != nil {
					return err
				}
			}
		}
		if err = log.PrintLog(writer, simple); err != nil {
			return err
		}
	}
	return nil
}

func (log *Log) PrintLog(writer io.Writer, simple bool) (err error) {

	if len(log.LogCurveInfo) == 0 {
		return nil
	}
	var printMinIDX, printMaxIDX func(LogCurveInfo) string
	if log.IndexType == "date time" {
		printMinIDX = func(curveInfo LogCurveInfo) string {
			return curveInfo.MinDateTimeIndex.Format(time.RFC3339)
		}
		printMaxIDX = func(curveInfo LogCurveInfo) string {
			return curveInfo.MaxDateTimeIndex.Format(time.RFC3339)
		}
	} else {
		printMinIDX = func(curveInfo LogCurveInfo) string {
			return fmt.Sprintf("%6.2f %s", curveInfo.MinIndex.Index, curveInfo.MinIndex.UOM)
		}
		printMaxIDX = func(curveInfo LogCurveInfo) string {
			return fmt.Sprintf("%6.2f %s", curveInfo.MaxIndex.Index, curveInfo.MaxIndex.UOM)
		}
	}
	if simple {
		if _, err = fmt.Fprintf(writer, "\t%5s\t%20s\t%10s\t%8s\t%-26s\t%-26s\t%20s\t%s\n", "Index", "UID", "Mnemonic", "unit", "min_idx",
			"max_idx", "MnemAlias", "classWitsml"); err != nil {
			return err
		}
		for _, curveInfo := range log.LogCurveInfo {
			if _, err = fmt.Fprintf(writer, "\t%5d\t%20s\t%10s\t%8s\t%-26s\t%-26s\t%20s\t%s\n", curveInfo.ColumnIndex, curveInfo.UID, curveInfo.Mnemonic,
				curveInfo.UNIT, printMinIDX(curveInfo), printMaxIDX(curveInfo), curveInfo.MnemAlias, curveInfo.ClassWITSML); err != nil {
				return err
			}
		}
	} else {
		if _, err = fmt.Fprintf(writer, "\t%5s\t%20s\t%10s\t%8s\t%-26s\t%-26s\t%40s\t%s\n", "Index", "UID", "Mnemonic", "unit", "min_idx",
			"max_idx", "Description", "Type"); err != nil {
			return err
		}
		for _, curveInfo := range log.LogCurveInfo {
			if _, err = fmt.Fprintf(writer, "\t%5d\t%20s\t%10s\t%8s\t%-26s\t%-26s\t%40s\t%s\n", curveInfo.ColumnIndex, curveInfo.UID, curveInfo.Mnemonic,
				curveInfo.UNIT, printMinIDX(curveInfo), printMaxIDX(curveInfo), strings.Replace(curveInfo.CurveDescription, "\n", " ", -1), curveInfo.TypeLogData); err != nil {
				return err
			}
		}
	}
	return nil
}
