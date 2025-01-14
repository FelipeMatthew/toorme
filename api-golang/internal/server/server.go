package server

import (
	"log/slog"
	"net/http"
	"strings"
	"toorme-api-golang/config"
	"toorme-api-golang/internal/router"
	"toorme-api-golang/pkg/logger"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Echo *echo.Echo
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewServer() *Server {
	slog.Info("Instantiating server")
	server := echo.New()
	corsAllowedOrigins := strings.Split(config.Config.CORS_ALLOWED_ORIGINS, ",")
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     corsAllowedOrigins,
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))
	server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}]  ${status}  ${method}  ${host}${path} ${latency_human}` + "\n",
	}))
	server.Validator = &CustomValidator{validator: validator.New()}
	r := router.NewRouter(server)
	r.SetupRoutes()
	return &Server{Echo: server}
}

func (s *Server) Start() {
	slog.Info("Starting server")
	err := s.Echo.Start(":8080")
	logger.LogOnError(err, "Failed to start server")
}
