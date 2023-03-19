package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"table-saas/app/service"
)

func TableGetById(c *gin.Context)  {
	viewId, ok := c.GetQuery("viewId")
	if !ok {
		c.JSON(http.StatusBadRequest, "get viewId failure")
		return
	}

	viewIdInt, err := strconv.ParseUint(viewId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, "viewId is not uint64")
		return
	}

	ser := service.ViewLogic{}
	res := ser.GetSetting(viewIdInt)
	c.JSON(http.StatusOK,res)
}