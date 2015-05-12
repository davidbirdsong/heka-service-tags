package hekaservicetags

import (
	"github.com/mozilla-services/heka/message"
	"github.com/mozilla-services/heka/pipeline"
)

type ServiceTaggerFilter struct {
	f int64
}

func (*ServiceTaggerFilter) Init(config interface{}) error {
	return nil
}

func (*ServiceTaggerFilter) Run(runner pipeline.FilterRunner, helper pipeline.PluginHelper) error {

	for pack := range runner.InChan() {
		runner.LogError(fmt.Errorf("in"))
		message.NewStringField(pack.Message, "foo", "bar")
		message.NewIntField(pack.Message, "servicetagger", 1, "")
		if !runner.Inject(pack) {
			runner.LogError(fmt.Errorf("failed to inject"))
		}

		//pack.Recycle()

	}
	return nil

}
func init() {
	pipeline.RegisterPlugin("ServiceTaggerFilter", func() interface{} {
		return new(ServiceTaggerFilter)
	})
}
