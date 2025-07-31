package alert

import (
	"github.com/edorguez/payment-reminder/internal/alert/handlers"
	"github.com/edorguez/payment-reminder/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	alertHandler handlers.AlertHandler
	router       *gin.Engine
}

func NewRoutes(alertHandler handlers.AlertHandler) *Routes {
	router := gin.Default()
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
