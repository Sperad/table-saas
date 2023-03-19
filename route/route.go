package route

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"table-saas/app/controller"
	"table-saas/config"
)

func Start()  {
	r := gin.Default()

	tableV1 := r.Group("/api/table/v1")
	{
		tableV1.GET("/view/get_by_id", controller.TableGetById)
	}

	port := config.App.Port
	err := r.Run(":" + strconv.Itoa(int(port)))
	if err != nil {
		return
	}
}
