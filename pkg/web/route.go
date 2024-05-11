package web

import (
	"InfoCounter/pkg/config"
	"github.com/gin-gonic/gin"
)

import (
	"InfoCounter/pkg/web/controller"
)

func RunHttp() {
	r := gin.Default()
	//解决跨域
	r.Use(config.CorsConfig())
	//路由组
	appInfoGroup := r.Group("/")
	{
		appInfoGroup.POST("/add-table", controller.NewInfoControllerImpl().AddTable)
		appInfoGroup.GET("/get-table-by-id/:key", controller.NewInfoControllerImpl().GetTableByID)
		appInfoGroup.POST("/create-user", controller.NewInfoControllerImpl().CreateUser)
		appInfoGroup.POST("/get-password", controller.NewInfoControllerImpl().GetPassword)
	}
	r.Run(":" + config.PORT)
}
