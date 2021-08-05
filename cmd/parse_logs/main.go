package main

import (
	"fmt"
	"github.com/br9k777/logconfig"
	"github.com/br9k777/witsml_client/pkg/elements"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"os"
)

func main() {

	logger, err := config.GetStandartLogger("production")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stdout, "Can't create logger %s", err)
		os.Exit(1)
	}
	zap.ReplaceGlobals(logger)

	c := config.GetDefaultZapConfig()
	c.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

	//c.DisableCaller = false
	//logger, err := c.Build()
	//if err != nil {
	//	_, _ = fmt.Fprintf(os.Stderr, "Fatal error, before set logger %s", err)
	//}
	//zap.ReplaceGlobals(logger)

	app := &cli.App{
		Name:    "parse_WITSML_logs",
		Version: "v0.1.0",
		Usage:   "get logs info from WITSML response and print",
		// ArgsUsage: "host port",
		Authors: []*cli.Author{
			{Name: "br9k"},
		},
	}

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "path-xml-file",
			Usage:   "path to xml file for parse",
			EnvVars: []string{"PRINT_WITSML_LOGS_PATH_XML_FILE"},
		},
		&cli.BoolFlag{
			Name:    "simple-out",
			Usage:   "less info for out",
			EnvVars: []string{"PRINT_WITSML_SIMPLE_OUT"},
		},
	}

	app.Action = func(c *cli.Context) (err error) {

		var logs *elements.Logs
		// /tmp/2GetLogs.xml"
		// /tmp/GetLogs.pretty_result.xml"
		if c.String("path-xml-file") == "" {
			return nil
		}
		if logs, err = elements.ReadLogsFromFile(c.String("path-xml-file")); err != nil {
			zap.S().Error(err)
			return
		}
		if err = logs.PrintLogs(os.Stdout, c.Bool("simple-out")); err != nil {
			zap.S().Error(err)
			return
		}

		return nil
	}
	if err = app.Run(os.Args); err != nil {
		zap.S().Fatal(err)
	}
}
