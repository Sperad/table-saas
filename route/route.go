package route

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"table-saas/app/controller"
	"table-saas/config"
)

func Start()  {
	r := gin.Default()

	viewV1 := r.Group("/api/view/v1")
	{
		viewV1.GET("/get_all", controller.ViewGetAll)
		viewV1.GET("/get_by_id", controller.ViewGetById)
		viewV1.POST("/add", controller.ViewAdd)
	}

	port := config.App.Port
	err := r.Run(":" + strconv.Itoa(int(port)))
	if err != nil {
		return
	}
}
