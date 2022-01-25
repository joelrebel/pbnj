package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/packethost/pkg/log/logr/v2"
	"github.com/tinkerbell/pbnj/cmd"
	"github.com/tinkerbell/pbnj/test/runner"
)

var (
	cfgData = runner.ConfigFile{
		Server: runner.Server{
			URL:  defaultURL,
			Port: defaultPort,
		},
	}
	defaultPort = randomInt(40000, 50000)
	defaultURL  = "localhost"
	cfg         = flag.String("config", "resources.yaml", "resources yaml file to read")
	logLevel    = flag.String("logLevel", "info", "log level (default: info")
)

func main() {
	fmt.Println("PBnJ Functional Testing")
	flag.Parse()
	if err := cfgData.Config(*cfg); err != nil {
		fmt.Fprintf(os.Stderr, "config error: %v", err)
		os.Exit(1)
	}

	if cfgData.Server.Port == "" && cfgData.Server.URL == "" {
		// start a local internal server
		go func() {
			logFile := "./pbnj.log"
			// remove existing log file
			os.Remove(logFile)
			serverCmd := cmd.NewRootCmd()
			serverCmd.SetArgs([]string{"server", "--port", defaultPort, "--logToFile", logFile})
			_ = serverCmd.Execute()
		}()
	}

	fmt.Println(*logLevel)
	logger, _, _ := logr.NewPacketLogr(logr.WithLogLevel(*logLevel))
	runner.RunTests(logger.Logger, cfgData)
}

// Returns an int >= min, < max.
func randomInt(min, max int) string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(max-min+1) + min)
}
