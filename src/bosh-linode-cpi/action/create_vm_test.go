package action_test

import (
	. "bosh-linode-cpi/action"
	instancemocks "bosh-linode-cpi/linode/instance_service/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	mock "github.com/stretchr/testify/mock"
)

var _ = Describe("CreateVM", func() {
	var (
		createVM   CreateVM
		vmService  *instancemocks.Service
		cloudProps VMCloudProperties
	)

	BeforeEach(func() {
		vmService = &instancemocks.Service{}
		createVM = NewCreateVM(vmService)
		cloudProps = VMCloudProperties{
			LinodeType: "fake-linode-type",
		}
	})
	AfterEach(func() {
		vmService.AssertExpectations(GinkgoT())
	})

	Describe("Run", func() {
		BeforeEach(func() {
			vmService.On("Create", mock.Anything, mock.AnythingOfType("string")).Return("fake-vm-id", nil)
		})

		It("creates the vm", func() {
			results, err := createVM.Run("fake-agent-id", "fake-stemcell-id", cloudProps)
			Expect(err).NotTo(HaveOccurred())
			Expect(results).To(Equal(VMCID("fake-vm-id")))
		})
	})
})
