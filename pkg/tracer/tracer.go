package tracer

import (
	"github.com/WeCanRun/gin-blog/global"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

func Setup(svcName, agentPort string) {
	tracer, _, err := NewJaegerTracer(svcName, agentPort)
	if err != nil {
		logging.Errorf("NewJaegerTracer fail, err: %v", err)
		return
	}
	global.Tracer = tracer
}

func NewJaegerTracer(svcName, agentPort string) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: svcName,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  agentPort,
		},
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}

	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, nil
}
