package logs

import (
	"context"
	"os"
	"path"

	"github.com/uber-go/zap"
	"google.golang.org/grpc/metadata"
)

var logger zap.Logger

func init() {
	// a fallback/root logger for events without context
	logger = zap.New(
		zap.NewJSONEncoder(zap.RFC3339Formatter("time")),
		zap.Fields(zap.Int("pid", os.Getpid()),
			zap.String("exe", path.Base(os.Args[0]))),
	)
}

// WithRqId returns a context which knows its request ID
func WithRqId(ctx context.Context, rqId string) context.Context {
	return metadata.NewContext(ctx, metadata.MD{
		"id": []string{rqId},
	})
}

// Logger returns a zap logger with as much context as possible
func Logger(ctx context.Context) zap.Logger {
	newLogger := logger
	if ctx != nil {
		if ctxMetadata, ok := metadata.FromContext(ctx); ok {
			if rqIds, ok := ctxMetadata["id"]; ok {
				if len(rqIds) == 1 {
					newLogger = newLogger.With(zap.String("id", rqIds[0]))
				}
			}
		}
	}
	return newLogger
}
