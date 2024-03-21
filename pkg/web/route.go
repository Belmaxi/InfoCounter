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
		appInfoGroup.POST("/add_table", controller.NewInfoControllerImpl().AddTable)
		appInfoGroup.GET("/get_table_by_id/:key", controller.NewInfoControllerImpl().GetTableByID)
	}
	r.Run("127.0.0.1:" + config.PORT)
}
