package config

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

var cpiRelease string

type LinodeConfig struct {
	UserAgentPrefix string `json:"user_agent_prefix"`
	LinodeToken     string `json:"linode_token"`
}

func (c LinodeConfig) GetUserAgent() string {
	if cpiRelease == "" {
		cpiRelease = "dev"
	}
	userAgent := "bosh-linode-cpi/" + cpiRelease
	if c.UserAgentPrefix == "" {
		return userAgent
	}
	return c.UserAgentPrefix + " " + userAgent
}

func (c LinodeConfig) Validate() error {
	if c.LinodeToken == "" {
		return bosherr.Error("Must provide a non-empty Linode Token")
	}

	return nil
}
