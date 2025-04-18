package router

import (
	"toorme-api-golang/internal/handler"
	"toorme-api-golang/internal/middleware"

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

	r.UserRoutes()
	r.AdminRoutes()
	r.DriverRoutes()
	r.CustomerRoutes()
}

func (r *Router) UserRoutes() {
	user := r.e.Group("/user", middleware.JWTMiddleware)
	user.Use(middleware.RoleMiddleware("admin"))

	user.GET("", handler.GetAllUser)
	user.GET("/:id", handler.GetUserById)
	user.POST("", handler.CreateUser)
	user.PUT("/:id", handler.UpdateUser)
	user.DELETE("/:id", handler.DeleteUser)
}

func (r *Router) AdminRoutes() {
	adminGroup := r.e.Group("/admin", middleware.JWTMiddleware)
	adminGroup.Use(middleware.RoleMiddleware("admin"))

	adminGroup.GET("/alldata", handler.FetchAllData)

}

func (r *Router) DriverRoutes() {
	driverGroup := r.e.Group("/driver", middleware.JWTMiddleware)
	driverGroup.Use(middleware.RoleMiddleware("driver", "admin"))

	driverGroup.GET("/trips", handler.DriverTrips)

}

func (r *Router) CustomerRoutes() {
	customerGroup := r.e.Group("/customer", middleware.JWTMiddleware)
	customerGroup.Use(middleware.RoleMiddleware("customer", "admin"))

}
