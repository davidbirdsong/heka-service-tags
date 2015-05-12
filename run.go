package hekaservicetags

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/mozilla-services/heka/message"
	"github.com/mozilla-services/heka/pipeline"
)

type ServiceTaggerFilter struct {
	f        int64
	client   *api.Client
	services map[string]*api.AgentService
}

func (t *ServiceTaggerFilter) update() error {
	var err error
	t.services, err = t.client.Agent().Services()
	return err
}

func (t *ServiceTaggerFilter) Init(config interface{}) error {
	var err error
	t.client, err = api.NewClient(api.DefaultConfig())
	return err
}

func (t *ServiceTaggerFilter) writeMessage(m *message.Message) {
	f := message.NewFieldInit("ServiceTagger", message.Field_STRING, "")

	for _, agentS := range t.services {
		f.AddValue(agentS.Service)
		/*
			for _, t := range agentS.Tags {
				message.NewStringField(m, strings.Join([]string{agentS.Service, t}, "="))
			}
		*/
	}
	m.AddField(f)
	message.NewIntField(m, "_servicetagger", 1, "")
}

func (t *ServiceTaggerFilter) Run(runner pipeline.FilterRunner, helper pipeline.PluginHelper) error {
	var (
		pack    *pipeline.PipelinePack
		running bool
	)
	running = true
	for running {
		select {
		case pack, running = <-runner.InChan():
			t.writeMessage(pack.Message)
			if !runner.Inject(pack) {
				runner.LogError(fmt.Errorf("failed to inject"))
			}
		case <-runner.Ticker():
			t.update()
			runner.LogError(fmt.Errorf("updating"))

			//pack.Recycle()

		}
	}
	return nil

}

func init() {
	pipeline.RegisterPlugin("ServiceTaggerFilter", func() interface{} {
		return new(ServiceTaggerFilter)
	})
}
