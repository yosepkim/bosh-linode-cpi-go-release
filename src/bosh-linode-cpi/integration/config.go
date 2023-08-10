package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"bosh-linode-cpi/action"
	boshapi "bosh-linode-cpi/api"
	boshdisp "bosh-linode-cpi/api/dispatcher"
	"bosh-linode-cpi/api/transport"
	boshcfg "bosh-linode-cpi/config"

	boshlogger "github.com/cloudfoundry/bosh-utils/logger"
)

var (
	// A stemcell that will be created in integration_suite_test.go
	existingStemcell string

	// Provided by user
	linodeToken = envRequired("LINODE_TOKEN")

	// Configurable defaults
	stemcellFile    = envOrDefault("STEMCELL_FILE", "")
	stemcellVersion = envOrDefault("STEMCELL_VERSION", "")

	cfgContent = fmt.Sprintf(`{
		"cloud": {
			"plugin": "linode",
			"properties": {
				"linode": {
					"linode_token": "%v"
				}
			}
		}
	}`, linodeToken)
)

func execCPI(request string) (boshdisp.Response, error) {
	var err error
	var cfg boshcfg.Config
	var in, out, errOut, errOutLog bytes.Buffer
	var boshResponse boshdisp.Response

	if cfg, err = boshcfg.NewConfigFromString(cfgContent); err != nil {
		return boshResponse, err
	}

	var req boshdisp.Request
	if err = json.Unmarshal([]byte(request), &req); err != nil {
		return boshResponse, err
	}

	// Marshal the modified request back to string
	requestByte, err := json.Marshal(req)
	if err != nil {
		return boshResponse, err
	}
	request = string(requestByte)

	multiWriter := io.MultiWriter(&errOut, &errOutLog)
	logger := boshlogger.NewWriterLogger(boshlogger.LevelDebug, multiWriter)
	multiLogger := boshapi.MultiLogger{Logger: logger, LogBuff: &errOutLog}

	actionFactory := action.NewConcreteFactory(
		cfg,
		multiLogger,
	)

	caller := boshdisp.NewJSONCaller()
	dispatcher := boshdisp.NewJSON(actionFactory, caller, multiLogger)

	in.WriteString(request)
	cli := transport.NewCLI(&in, &out, dispatcher, multiLogger)

	var response []byte

	if err = cli.ServeOnce(); err != nil {
		return boshResponse, err
	}

	if response, err = ioutil.ReadAll(&out); err != nil {
		return boshResponse, err
	}

	if err = json.Unmarshal(response, &boshResponse); err != nil {
		return boshResponse, err
	}
	return boshResponse, nil
}

func envRequired(key string) (val string) {
	if val = os.Getenv(key); val == "" {
		panic(fmt.Sprintf("Could not find required environment variable '%s'", key))
	}
	return
}

func envOrDefault(key, defaultVal string) (val string) {
	if val = os.Getenv(key); val == "" {
		val = defaultVal
	}
	return
}
