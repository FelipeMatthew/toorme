package router

import (
	"toorme-api-golang/internal/app/handler"
	"toorme-api-golang/internal/app/middleware"

	"github.com/labstack/echo/v4"
)

type Router struct {
	e *echo.Echo
}

func NewRouter(e *echo.Echo) *Router {
	return &Router{e: e}
}

func (r *Router) SetupRoutes() {
	r.e.GET("/ping", handler.Ping)
	r.e.POST("/login", handler.Login)

	r.setupAdminRoutes()
	r.setupDriverRoutes()
	r.setupCustomerRoutes()
}

func (r *Router) setupAdminRoutes() {
	adminGroup := r.e.Group("/admin", middleware.JWTMiddleware)
	adminGroup.Use(middleware.RoleMiddleware("admin"))

	adminGroup.GET("/dashboard", handler.AdminDashboard)
}

func (r *Router) setupDriverRoutes() {
	driverGroup := r.e.Group("/driver", middleware.JWTMiddleware)
	driverGroup.Use(middleware.RoleMiddleware("driver", "admin"))

	driverGroup.GET("/trips", handler.DriverTrips)
}

func (r *Router) setupCustomerRoutes() {
	customerGroup := r.e.Group("/customer", middleware.JWTMiddleware)
	customerGroup.Use(middleware.RoleMiddleware("customer", "admin"))

	customerGroup.GET("/plans", handler.CustumerPlans)
}
