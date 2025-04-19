package test

import (
	"context"
	"github.com/gvillela7/ratelimit/configs"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type JsonResp struct {
	StatusCode int           `json:"StatusCode"`
	Message    string        `json:"message,omitempty"`
	Data       []interface{} `json:"data,omitempty"`
}

var (
	i   int
	req *http.Request
	res *http.Response
)

func UnBlock() bool {
	_ = configs.Load("../")
	cfg := configs.GetAPIConfig()
	reteLimitRequestBlock := cfg.RateLimitTimeBlock
	time.Sleep(time.Duration(reteLimitRequestBlock+1) * time.Second)

	req, _ = http.NewRequest("GET", "http://127.0.0.1:"+cfg.Port, nil)

	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Second)
	defer cancel()
	req = req.WithContext(ctx)
	client := http.DefaultClient
	res, _ = client.Do(req)
	return true
}
func TestRateLimitBlock(t *testing.T) {
	err := configs.Load("../")
	assert.NoError(t, err, "Load config fail")
	cfg := configs.GetAPIConfig()
	reteLimitRequest := cfg.RateLimitRequest

	for i = 0; i <= reteLimitRequest; i++ {
		req, err = http.NewRequest("GET", "http://127.0.0.1:"+cfg.Port, nil)
		assert.NoError(t, err, "creating request should not fail")
		ctx, cancel := context.WithTimeout(req.Context(), 1*time.Second)
		defer cancel()
		req = req.WithContext(ctx)
		client := http.DefaultClient
		res, err = client.Do(req)
	}
	assert.Equal(t, http.StatusTooManyRequests, res.StatusCode)
	if i == reteLimitRequest+1 {
		resp := UnBlock()
		assert.True(t, resp)
		assert.Equal(t, http.StatusOK, res.StatusCode)
	}
}
