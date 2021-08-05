package elements

import (
	"fmt"
	config "github.com/br9k777/logconfig"
	"go.uber.org/zap"
	"os"
	"testing"
)

const (
	testLogsXML = "tests/log_no_xsl.xml"
)

func TestReadLogs(t *testing.T) {
	var logs *Logs
	var err error
	if logs, err = ReadLogsFromFile(testLogsXML); err != nil {
		zap.S().Error(err)
		return
	}
	if err = logs.PrintLogs(os.Stdout, false); err != nil {
		zap.S().Error(err)
		return
	}
}

func init() {
	c := config.GetDefaultZapConfig()
	c.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	c.DisableCaller = false
	logger, err := c.Build()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Fatal error, before set logger %s", err)
	}
	zap.ReplaceGlobals(logger)
}
