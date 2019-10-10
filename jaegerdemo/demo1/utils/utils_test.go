package utils

import (
	"context"
	"testing"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func TestInjectTraceID(t *testing.T) {
	ctx := context.Background()
	span, _ := opentracing.StartSpanFromContext(
		ctx,
		"call Http Get",
		opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
		ext.SpanKindRPCClient,
	)
	defer span.Finish()
}

