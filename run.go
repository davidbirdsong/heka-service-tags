package hekaservicetags

import (
	"github.com/mozilla-services/heka/message"
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
	RegisterPlugin("ServiceTagger", func() interface{} {
		return &tags{}
	})
}
