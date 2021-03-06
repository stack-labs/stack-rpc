package memory

import (
	"github.com/stack-labs/stack/plugin"
	"github.com/stack-labs/stack/transport"
)

type grpcTransportPlugin struct {
}

func (g *grpcTransportPlugin) Name() string {
	return "memory"
}

func (g *grpcTransportPlugin) Options() []transport.Option {
	return nil
}

func (g *grpcTransportPlugin) New(opts ...transport.Option) transport.Transport {
	return NewTransport(opts...)
}

func init() {
	plugin.TransportPlugins["memory"] = &grpcTransportPlugin{}
}
