package controller

import (
	"dobo/database/model"
	"dobo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)


// @Title dockerImage
// @Description 获取镜像列表
// @Security ApiKeyAuth
// @Success 200 "获取版本成功"
// @Router /images/getImageList [get]
func GetImgesList(c *gin.Context)  {
	info, err := model.GetImageList()
	if err != nil {
		c.JSON(http.StatusOK,utils.RespError(http.StatusBadRequest,err))
		return
	}
	c.JSON(http.StatusOK,utils.RespData(http.StatusOK,info))
}

//imageId string

// @Title dockerImage
// @Description 根据镜像id获取镜像信息
// @Security ApiKeyAuth
// @Success 200 "获取镜像信息成功"
// @Router /images/getImageInfo [get]
func GetImageInfo(c *gin.Context)  {

}


// @Title deleteDocker
// @Description 根据镜像id删除镜像
// @Security ApiKeyAuth
// @Success 200 "删除镜像成功"
// @Router /images/getImageInfo [get]
func DeleteImage()  {

}


// @Title reTagImage
// @Description 重命名tag
// @Security ApiKeyAuth
// @Success 200 "重命名成功"
// @Router /images/reTagImage [post]
func ReTagImage(c *gin.Context)  {

}