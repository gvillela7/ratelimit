package routes

import (
	"github.com/gvillela7/ratelimit/configs"
	redis2 "github.com/gvillela7/ratelimit/internal/database/redis"
	"github.com/gvillela7/ratelimit/internal/handler"
	"github.com/gvillela7/ratelimit/internal/middlewares"
	"github.com/gvillela7/ratelimit/internal/util"
	"net/http"
)

func Routes() {
	err := configs.Load(".")
	if err != nil {
		util.Log(true, false, "error", "failed to initialize environment variables:", "error", err)
		panic(err)
	}
	client, err := redis2.NewRedis()
	if err != nil {
		util.Log(true, false, "error", "failed to initialize redis:", "error", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.OKHandler)

	cfg := configs.GetAPIConfig()
	util.Log(true, false, "info", "Listening on:", "port", cfg.Port)

	nextMiddleware := middlewares.Limit(client, mux)

	http.ListenAndServe(":"+cfg.Port, nextMiddleware)
}
