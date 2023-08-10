package integration

import (
	"context"
	"strconv"

	"github.com/linode/linodego"
	. "github.com/onsi/gomega"
)

func assertSucceeds(request string) {
	response, err := execCPI(request)
	Expect(err).ToNot(HaveOccurred())
	Expect(response.Error).To(BeNil())
}

func assertFails(request string) error {
	response, _ := execCPI(request)
	Expect(response.Error).ToNot(BeNil())
	return response.Error
}

func assertSucceedsWithResult(request string) interface{} {
	response, err := execCPI(request)
	Expect(err).ToNot(HaveOccurred())
	Expect(response.Error).To(BeNil())
	Expect(response.Result).ToNot(BeNil())
	return response.Result
}

func toStringArray(raw []interface{}) []string {
	strings := make([]string, len(raw), len(raw))
	for i := range raw {
		strings[i] = raw[i].(string)
	}
	return strings
}

func assertValidVM(id string, valFunc func(*linodego.Instance)) {
	linodeId, err := strconv.Atoi(id)
	Expect(err).To(BeNil())
	linodeClient := CreateTestClient()
	instance, err := linodeClient.GetInstance(context.Background(), linodeId)
	Expect(err).To(BeNil())
	valFunc(instance)
}
