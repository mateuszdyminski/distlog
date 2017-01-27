package logs

import (
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
	"testing"
)

func TestCtx(t *testing.T) {
	reqIdKey := "reqId"
	reqId := "RequestID"

	testMetadata := metadata.MD{
		reqIdKey: []string{reqId},
	}

	ctx1 := context.Background()

	ctx2 := metadata.NewContext(ctx1, testMetadata)
	meta, ok := metadata.FromContext(ctx2)

	assert.Equal(t, ok, true, "The metadata should be available.")
	assert.Equal(t, meta[reqIdKey][0], reqId, "Request id should be the same")
}
