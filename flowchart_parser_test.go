package visibleIdeas

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_flowchartJsonStringToWorkflow_bare_essentials(t *testing.T) {
	RegisterTestingT(t)
	{
		_, err := flowchartJsonStringToWorkflow("")
		Expect(err.Error()).To(ContainSubstring("not a valid json"))
	}
	{
		_, err := flowchartJsonStringToWorkflow("somthin ghta tis not valid json {")
		Expect(err.Error()).To(ContainSubstring("not a valid json string"))
	}
}

func Test_flowchartJsonStringToWorkflow_basics(t *testing.T) {
	RegisterTestingT(t)

}
