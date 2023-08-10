package config

import (
	"encoding/json"

	boshlinodeconfig "bosh-linode-cpi/linode/config"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
)

type Config struct {
	Cloud Cloud
}

type Cloud struct {
	Plugin     string
	Properties CPIProperties
}

type CPIProperties struct {
	Linode boshlinodeconfig.LinodeConfig
}

func NewConfigFromPath(configFile string, fs boshsys.FileSystem) (Config, error) {
	var config Config

	if configFile == "" {
		return config, bosherr.Errorf("Must provide a config file")
	}

	bytes, err := fs.ReadFile(configFile)
	if err != nil {
		return config, bosherr.WrapErrorf(err, "Reading config file '%s'", configFile)
	}

	if err = json.Unmarshal(bytes, &config); err != nil {
		return config, bosherr.WrapError(err, "Unmarshalling config contents")
	}

	if err = config.Validate(); err != nil {
		return config, bosherr.WrapError(err, "Validating config")
	}

	return config, nil
}

func NewConfigFromString(configString string) (Config, error) {
	var config Config
	var err error
	if configString == "" {
		return config, bosherr.Errorf("Must provide a config")
	}

	if err = json.Unmarshal([]byte(configString), &config); err != nil {
		return config, bosherr.WrapError(err, "Unmarshalling config contents")
	}

	if err = config.Validate(); err != nil {
		return config, bosherr.WrapError(err, "Validating config")
	}

	return config, nil
}

func (c Config) Validate() error {
	if c.Cloud.Plugin != "linode" {
		return bosherr.Errorf("Unsupported cloud plugin type %q", c.Cloud.Plugin)
	}
	if err := c.Cloud.Properties.Linode.Validate(); err != nil {
		return bosherr.WrapError(err, "Validating Linode configuration")
	}

	return nil
}
