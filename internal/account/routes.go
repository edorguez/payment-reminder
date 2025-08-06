package account

import (
	"time"

	"github.com/edorguez/payment-reminder/internal/account/handlers"
	"github.com/edorguez/payment-reminder/pkg/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	userHandler handlers.UserHandler
	router      *gin.Engine
}

func NewRoutes(userHandler handlers.UserHandler) *Routes {
	router := gin.Default()
	corsConfig := corsConfig()
	router.Use(cors.New(corsConfig))
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
		userGroup.POST("", r.userHandler.Create)
		userGroup.GET("", r.userHandler.ListOrFind)
		userGroup.GET(":id", r.userHandler.FindById)
		userGroup.PUT(":id", r.userHandler.Update)
		userGroup.DELETE(":id", r.userHandler.Delete)
	}
}

func corsConfig() cors.Config {
	return cors.Config{
		// Allow BOTH forms that the browser may send
		AllowOrigins: []string{
			"http://localhost:5173",
			"http://0.0.0.0:5173",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
}
