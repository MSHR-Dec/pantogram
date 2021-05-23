package jaeger

import (
	"io"

	"github.com/opentracing/opentracing-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func NewTracer() (opentracing.Tracer, io.Closer, error) {
	cfg, _ := jaegercfg.FromEnv()

	return cfg.NewTracer()
}
