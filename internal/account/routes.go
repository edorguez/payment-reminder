package account

import (
	"github.com/edorguez/payment-reminder/internal/account/handlers"
	"github.com/edorguez/payment-reminder/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	userHandler handlers.UserHandler
	router      *gin.Engine
}

func NewRoutes(userHandler handlers.UserHandler) *Routes {
	router := gin.Default()
	router.Use(middleware.FirebaseAuth())

	routes := &Routes{userHandler: userHandler, router: router}
	routes.addUserRoutes()

	return routes
}

func (r *Routes) Start(address string) error {
	return r.router.Run(address)
}

func (r *Routes) addUserRoutes() {
	userGroup := r.router.Group("/api/users")
	{
		userGroup.POST("/", r.userHandler.Create)
		userGroup.GET("/", r.userHandler.ListOrFind)
		userGroup.GET("/:id", r.userHandler.FindById)
		userGroup.PUT("/:id", r.userHandler.Update)
		userGroup.DELETE("/:id", r.userHandler.Delete)
	}
}
