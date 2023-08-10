package integration

import (
	"fmt"

	"github.com/linode/linodego"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("VM", func() {
	It("can create a VM", func() {
		request := fmt.Sprintf(`{
					"method": "create_vm",
					"arguments": [
						"agent",
						"%v",
						{
							"linode_type": "g6-standard-1",
							"tags": ["tag1", "tag2"],
							"region": "us-east",
							"tags": ["tag1", "integration-delete"]
						}
					]
					}`, "todo-existing-stemcell?")

		vmCID := assertSucceedsWithResult(request).(string)
		Expect(vmCID).ToNot(BeEmpty())
		assertValidVM(vmCID, func(instance *linodego.Instance) {
			Expect(instance.Region).To(Equal("us-east"))
			Expect(instance.Type).To(Equal("g6-standard-1"))
			Expect(instance.Tags).To(ConsistOf("integration-delete", "tag1"))
		})
	})

	It("executes the VM lifecycle", func() {
		var vmCID string
		By("creating a VM")
		request := fmt.Sprintf(`{
			  "method": "create_vm",
			  "arguments": [
				"agent",
				"%v",
				{
				  "linode_type": "g6-standard-1",
				  "label": "integration-label",
				  "region": "us-east",
				  "tags": ["integration-delete"]
				}
			  ]
			}`, existingStemcell)
		vmCID = assertSucceedsWithResult(request).(string)

		By("deleting the VM")
		request = fmt.Sprintf(`{
			  "method": "delete_vm",
			  "arguments": ["%v"]
			}`, vmCID)
		assertSucceeds(request)

	})
})
