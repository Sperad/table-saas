package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"table-saas/app/dto/req"
	"table-saas/app/service"
)

func ViewGetAll(c *gin.Context) {
	uid, ok := c.GetQuery("uid")
	if !ok {
		c.JSON(http.StatusBadRequest, "get uid failure")
		return
	}
	groupKey, ok := c.GetQuery("groupKey")
	if !ok {
		c.JSON(http.StatusBadRequest, "get groupKey failure")
		return
	}


	uidInt, err := strconv.ParseUint(uid, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, "uid is not uint64")
		return
	}

	ser := service.ViewLogic{}
	res := ser.GetAll(uidInt, groupKey)
	c.JSON(http.StatusOK,res)
}

// ViewGetById /**
func ViewGetById(c *gin.Context)  {
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

// ViewAdd /**
func ViewAdd(c *gin.Context)  {
	var viewAddReq = req.ViewAddReq{}
	err := c.ShouldBindJSON(viewAddReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, viewAddReq)
}