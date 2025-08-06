package alert

import (
	"time"

	"github.com/edorguez/payment-reminder/internal/alert/handlers"
	"github.com/edorguez/payment-reminder/pkg/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	alertHandler handlers.AlertHandler
	router       *gin.Engine
}

func NewRoutes(alertHandler handlers.AlertHandler) *Routes {
	router := gin.Default()
	corsConfig := corsConfig()
	router.Use(cors.New(corsConfig))
	router.Use(middleware.FirebaseAuth())

	routes := &Routes{alertHandler: alertHandler, router: router}
	routes.addAlertRoutes()

	return routes
}

func (r *Routes) Start(address string) error {
	return r.router.Run(address)
}

func (r *Routes) addAlertRoutes() {
	userGroup := r.router.Group("/api/alerts")
	{
		userGroup.POST("/", r.alertHandler.Create)
		userGroup.GET("/:id", r.alertHandler.FindById)
		userGroup.PUT("/:id", r.alertHandler.Update)
		userGroup.DELETE("/:id", r.alertHandler.Delete)
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
