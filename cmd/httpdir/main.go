package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

type Config struct {
	Port int
}

func getConfig() (cfg Config) {
	flag.IntVar(&cfg.Port, "port", 8080, "port to listen on")
	flag.Parse()
	return cfg
}

func run() error {
	cfg := getConfig()

	lc := zap.NewDevelopmentConfig()
	lc.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	log, err := lc.Build()
	if err != nil {
		return fmt.Errorf("create zap logger: %w", err)
	}
	//nolint:all
	defer log.Sync()

	r := mux.NewRouter()
	r.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Info("request", zap.String("method", r.Method), zap.String("path", r.URL.Path))
			h.ServeHTTP(w, r)
		})
	})
	r.Use(mux.CORSMethodMiddleware(r))
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(".")))

	log.Info("serving", zap.Int("port", cfg.Port))

	srv := http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Port),
		Handler:           r,
		ReadHeaderTimeout: time.Minute,
	}

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("http.ListenAndServe: %w", err)
	}

	return nil
}
