package action_test

import (
	. "bosh-linode-cpi/action"
	instancemocks "bosh-linode-cpi/linode/instance_service/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	mock "github.com/stretchr/testify/mock"
)

var _ = Describe("DeleteVM", func() {
	var (
		deleteVM  DeleteVM
		vmService *instancemocks.Service
	)

	BeforeEach(func() {
		vmService = &instancemocks.Service{}
		deleteVM = NewDeleteVM(vmService)
	})

	AfterEach(func() {
		vmService.AssertExpectations(GinkgoT())
	})

	Describe("Run", func() {
		BeforeEach(func() {
			vmService.On("Delete", mock.AnythingOfType("string")).Return(nil)
		})

		It("deletes the vm", func() {
			_, err := deleteVM.Run("fake-vm-id")
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
