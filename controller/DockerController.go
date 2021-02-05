package controller

import (
	"dobo/database/model"
	"dobo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)



// @Title dockerInfo
// @Description 获取docker信息
// @Security ApiKeyAuth
// @Success 200 "获取docker信息成功"
// @Router /docker/dockerInfo [get]
func DockerInfo(c *gin.Context ) {
	imageList, err := model.GetDockerInfo()
	if err != nil {
		c.JSON(http.StatusOK,utils.RespError(http.StatusBadRequest,err))
		return
	}
	c.JSON(http.StatusOK,utils.RespData(http.StatusOK,imageList))
}

// @Title dockerVersion
// @Description 获取docker信息
// @Security ApiKeyAuth
// @Success 200 "获取版本成功"
// @Router /docker/getVersion [get]
func GerVersion(c *gin.Context)  {
	version := model.GetVersion()
	c.JSON(http.StatusOK,utils.RespData(http.StatusOK,version))
}

// @Title ping
// @Description 检测心跳
// @Security ApiKeyAuth
// @Success 200 "心跳检测通过"
// @Router /docker/ping [get]
func Ping(c *gin.Context)  {
	ping, err := model.Ping()
	if err != nil {
		c.JSON(http.StatusOK,utils.RespError(http.StatusBadRequest,err))
		return
	}
	c.JSON(http.StatusOK,utils.RespData(http.StatusOK,ping))
}

// @Title DiskUsage
// @Description 检测磁盘
// @Security ApiKeyAuth
// @Success 200 "磁盘检测通过"
// @Router /docker/diskUsage [get]
func DiskUsage(c *gin.Context)  {
	disk, err := model.Disk()
	if err != nil {
		c.JSON(http.StatusOK,utils.RespError(http.StatusBadRequest, err))
		return
	}
	c.JSON(http.StatusOK, utils.RespData(http.StatusOK, disk))
}
