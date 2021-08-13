package main

import (
	"net/http"
	"time"

	"github.com/cuijxin/blog-service/internal/routers"
)

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":9090",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
