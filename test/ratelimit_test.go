package test

import (
	"context"
	"github.com/gvillela7/ratelimit/configs"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRateLimit(t *testing.T) {
	err := configs.Load("../")
	assert.NoError(t, err, "Load config fail")
	cfg := configs.GetAPIConfig()
	reteLimitRequest := cfg.RateLimitRequest

	for i := 0; i < reteLimitRequest; i++ {
		req, err := http.NewRequest("GET", "http://app:"+cfg.Port, nil)
		ctx, cancel := context.WithTimeout(req.Context(), 1*time.Second)
		defer cancel()
		req = req.WithContext(ctx)
		client := http.DefaultClient
		res, err := client.Do(req)

		assert.NoError(t, err, "creating request should not fail")
		assert.Equal(t, http.StatusOK, res.StatusCode)
	}
}
