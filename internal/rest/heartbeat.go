package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	kit_logger "github.com/sailsforce/gomicro-kit/logger"
	kit_models "github.com/sailsforce/gomicro-kit/models"
	kit_utils "github.com/sailsforce/gomicro-kit/utils"

	"github.com/sailsforce/togo-write-micro/internal/config"
)

func HeartbeatRotues() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", GetHeartbeat)

	return router
}

func GetHeartbeat(rw http.ResponseWriter, req *http.Request) {
	logger := kit_logger.GetLogEntry(req)
	dbOnline := false

	logger.Info("pinging database... ", config.RV.DatabaseURL)
	if err := kit_utils.PingDB(config.Psql); err == nil {
		dbOnline = true
	}
	response := kit_models.Heartbeat{
		RequestID:      middleware.GetReqID(req.Context()),
		DatabaseOnline: dbOnline,
		AppName:        config.RV.AppName,
		ReleaseDate:    config.RV.ReleaseDate,
		ReleaseVersion: config.RV.ReleaseVersion,
		Slug:           config.RV.Slug,
		Message:        "togo write service",
	}

	logger.Info("heartbeat request finished.")
	render.JSON(rw, req, response)
}
