package integration

import (
	boshapi "bosh-linode-cpi/api"
	boshlinodeclient "bosh-linode-cpi/linode/client"
	"bosh-linode-cpi/linode/config"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	"github.com/linode/linodego"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

var _ = SynchronizedBeforeSuite(func() []byte {
	// Clean any straggler VMs
	cleanVMs()
	data := ""
	return []byte(data)
}, func(data []byte) {
	// Required env vars
	Expect(linodeToken).ToNot(Equal(""), "LINODE_TOKEN must be set")
})

var _ = SynchronizedAfterSuite(func() {}, func() {
	cleanVMs()
})

func CreateTestClient() *linodego.Client {
	var logBuff bytes.Buffer
	multiWriter := io.MultiWriter(os.Stderr, &logBuff)
	logger := boshlog.NewWriterLogger(boshlog.LevelDebug, multiWriter)
	multiLogger := boshapi.MultiLogger{Logger: logger, LogBuff: &logBuff}

	linodeClient, err := boshlinodeclient.NewLinodeClient(config.LinodeConfig{LinodeToken: linodeToken}, multiLogger)
	Expect(err).To(BeNil())
	return &linodeClient
}

func cleanVMs() {
	// Clean up any VMs left behind from failed tests. Instances with the 'integration-delete' tag will be deleted.
	GinkgoWriter.Write([]byte("Looking for Linodes with 'integration-delete' tag. Matches will be deleted\n"))

	linodeClient := CreateTestClient()
	Expect(linodeClient).ToNot(BeNil())
	instances, err := linodeClient.ListInstances(context.Background(), linodego.NewListOptions(0, `{"tags": "integration-delete"}`))
	Expect(err).To(BeNil())
	for _, instance := range instances {
		Expect(instance.Tags).To(ContainElement("integration-delete"))
		GinkgoWriter.Write([]byte(fmt.Sprintf("Deleting Linode %v\n", instance.ID)))
		err := linodeClient.DeleteInstance(context.Background(), instance.ID)
		Expect(err).ToNot(HaveOccurred())
	}
}
