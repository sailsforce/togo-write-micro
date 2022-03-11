package internal

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	kit_logger "github.com/sailsforce/gomicro-kit/logger"
	kit_middleware "github.com/sailsforce/gomicro-kit/middleware"
	"github.com/sailsforce/togo-write-micro/internal/config"
	"github.com/sailsforce/togo-write-micro/internal/rest"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(render.SetContentType(render.ContentTypeJSON),
		middleware.RedirectSlashes,
		middleware.RequestID,
		kit_logger.NewStructuredLogger(config.Logger),
		kit_middleware.Headers,
		middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		r.Mount("/heartbeat", rest.HeartbeatRotues())
	})

	r.Get("/", func(rw http.ResponseWriter, req *http.Request) {
		config.Logger.Debug(rw.Write([]byte("online")))
	})

	return r
}
