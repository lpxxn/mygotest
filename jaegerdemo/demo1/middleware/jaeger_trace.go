package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

const spanContextKey = "spanContext"

func SetupRootTrace() gin.HandlerFunc {
	return func(c *gin.Context) {
		tracer := opentracing.GlobalTracer()
		var parentSpan opentracing.Span
		spCtx, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		if err != nil {
			parentSpan = tracer.StartSpan(c.Request.URL.Path)
			defer parentSpan.Finish()
		} else {
			parentSpan = opentracing.StartSpan(
				c.Request.URL.Path,
				opentracing.ChildOf(spCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
				ext.SpanKindRPCServer,
			)
			defer parentSpan.Finish()
		}
		c.Set(spanContextKey, parentSpan)
		c.Next()
	}
}

// 上面那个parentSpan不太好，可以用下面这个
func Middleware() gin.HandlerFunc {
	return func(g *gin.Context) {
		tracer := opentracing.GlobalTracer()
		ctx := g.Request.Context()
		spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(g.Request.Header))
		span := tracer.StartSpan("http: "+g.Request.RequestURI, ext.RPCServerOption(spanCtx))
		g.Request = g.Request.WithContext(opentracing.ContextWithSpan(ctx, span))
		defer span.Finish()
		g.Next()
	}
}

func SpanFromCtx() gin.HandlerFunc {
	return func(c *gin.Context) {
		tracer := opentracing.GlobalTracer()
		spanContext, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		if err != nil {
			c.Next()
			return
		}

		opts := append([]opentracing.StartSpanOption{opentracing.ChildOf(spanContext)})

		span := tracer.StartSpan(c.Request.URL.Path, opts...)
		c.Set(spanContextKey, span)
		defer span.Finish()

		c.Next()
	}
}

// GetSpan extracts span from context.
func GetSpan(ctx *gin.Context) (span opentracing.Span, exists bool) {
	spanI, _ := ctx.Get(spanContextKey)
	span, ok := spanI.(opentracing.Span)
	exists = span != nil && ok
	return
}