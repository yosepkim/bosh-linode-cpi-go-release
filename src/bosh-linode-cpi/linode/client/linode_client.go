package client

import (
	"net/http"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/linode/linodego"
	"golang.org/x/oauth2"

	"bosh-linode-cpi/linode/config"
)

func NewLinodeClient(
	config config.LinodeConfig,
	logger boshlog.Logger, //TODO: What to do with logger?
) (linodego.Client, error) {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: config.LinodeToken})

	oauthTransport := &oauth2.Transport{
		Source: tokenSource,
	}
	oauth2Client := &http.Client{
		Transport: oauthTransport,
	}
	return linodego.NewClient(oauth2Client), nil
}
