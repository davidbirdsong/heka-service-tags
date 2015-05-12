package hekaservicetags

import (
	"github.com/mozilla-services/heka/message"
	"github.com/mozilla-services/heka/pipeline"
)

type tags struct {
	f int64
}

func (*tags) Run(runner pipeline.FilterRunner, helper pipeline.PluginHelper) error {

	for pack := range runner.InChan() {
		message.NewStringField(pack.Message, "foo", "bar")
		runner.Inject(pack)
		//pack.Recycle()

	}
	return nil

}
func init() {
	pipeline.RegisterPlugin("ServiceTagger", func() interface{} {
		return &tags{}
	})
}
