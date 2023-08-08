package container

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sarulabs/di"
	producer "github.com/taranovegor/com.jurnalo/internal/amqp"
	"github.com/taranovegor/com.jurnalo/internal/collector"
	"github.com/taranovegor/com.jurnalo/internal/config"
	"github.com/taranovegor/com.jurnalo/internal/handler/tcp"
	"net"
)

const (
	ScopeApp     = di.App
	ScopeTcp     = ScopeApp // "scope_http"
	Collector    = "collector"
	ListenerTcp  = "listener_tcp"
	HandlerTcp   = "handler_tcp"
	Amqp         = "amqp"
	AmqpChannel  = "amqp_channel"
	AmqpProducer = "amqp_producer"
)

type ServiceContainer interface {
	Get(name string) interface{}
}

type serviceContainer struct {
	ServiceContainer
	container di.Container
}

func Init() (ServiceContainer, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}

	return &serviceContainer{
		container: build(builder),
	}, nil
}

func (sc serviceContainer) Get(name string) interface{} {
	return sc.container.Get(name)
}

func build(builder *di.Builder) di.Container {
	builder.Add(di.Def{
		Name: Amqp,
		Build: func(c di.Container) (interface{}, error) {
			return amqp.Dial(config.GetEnv(config.AmqpDsn))
		},
	})

	builder.Add(di.Def{
		Name: AmqpChannel,
		Build: func(c di.Container) (interface{}, error) {
			return c.Get(Amqp).(*amqp.Connection).Channel()
		},
	})

	builder.Add(di.Def{
		Name: AmqpProducer,
		Build: func(c di.Container) (interface{}, error) {
			return producer.NewProducer(
				config.GetAmqpConfig().Producers,
				c.Get(AmqpChannel).(*amqp.Channel),
			), nil
		},
	})

	buildCollector(builder)
	buildHandler(builder)

	return builder.Build()
}

func buildCollector(builder *di.Builder) {
	builder.Add(di.Def{
		Name: Collector,
		Build: func(ctn di.Container) (interface{}, error) {
			return collector.NewCollector(
				config.GetEnv(config.CollectorExpr),
			), nil
		},
	})
}

func buildHandler(builder *di.Builder) {
	buildHandlerTcp(builder)
}

func buildHandlerTcp(builder *di.Builder) {
	builder.Add(di.Def{
		Name: ListenerTcp,
		Build: func(ctn di.Container) (interface{}, error) {
			return net.Listen("tcp", config.GetEnv(config.TcpDsn))
		},
	})

	builder.Add(di.Def{
		Name: HandlerTcp,
		Build: func(ctn di.Container) (interface{}, error) {
			return tcp.NewHandler(
				ctn.Get(ListenerTcp).(net.Listener),
				ctn.Get(Collector).(*collector.Collector),
			), nil
		},
	})
}
