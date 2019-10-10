package utils

import (
	"fmt"
	"io"
	"net/http"
	"runtime"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
)

func NewJaegerTracer(serviceName string, jaegerHostPort string) (opentracing.Tracer, io.Closer, error) {
	cfg := &jaegerConfig.Configuration{
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  "const", //固定采样
			Param: 1,       //1=全采样、0=不采样
		},

		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: jaegerHostPort,
		},

		ServiceName: serviceName,
	}

	tracer, closer, err := cfg.NewTracer(jaegerConfig.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, err
}

// StartSpan will start a new span with no parent span.
func StartSpan(operationName, method, path string) opentracing.Span {
	return StartSpanWithParent(nil, operationName, method, path)
}

// StartSpanWithParent will start a new span with a parent span.
// example:
//      span:= StartSpanWithParent(c.Get("tracing-context"),
func StartSpanWithParent(parent opentracing.SpanContext, operationName, method, path string) opentracing.Span {
	options := []opentracing.StartSpanOption{
		opentracing.Tag{Key: ext.SpanKindRPCServer.Key, Value: ext.SpanKindRPCServer.Value},
		opentracing.Tag{Key: string(ext.HTTPMethod), Value: method},
		opentracing.Tag{Key: string(ext.HTTPUrl), Value: path},
		opentracing.Tag{Key: "current-goroutines", Value: runtime.NumGoroutine()},
	}

	if parent != nil {
		options = append(options, opentracing.ChildOf(parent))
	}

	return opentracing.StartSpan(operationName, options...)
}

// InjectTraceID injects the span ID into the provided HTTP header object, so that the
// current span will be propogated downstream to the server responding to an HTTP request.
// Specifying the span ID in this way will allow the tracing system to connect spans
// between servers.
//
//  Usage:
//          // resty example
// 	    r := resty.R()
//	    injectTraceID(span, r.Header)
//	    resp, err := r.Get(fmt.Sprintf("http://localhost:8000/users/%s", bosePersonID))
//
//          // galapagos_clients example
//          c := galapagos_clients.GetHTTPClient()
//          req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8000/users/%s", bosePersonID))
//          injectTraceID(span, req.Header)
//          c.Do(req)
func InjectTraceID(ctx opentracing.SpanContext, header http.Header) error {
	return opentracing.GlobalTracer().Inject(ctx, opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(header))
}
