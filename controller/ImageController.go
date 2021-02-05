package controller

import (
	"dobo/database/model"
	"dobo/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


// @Title dockerImage
// @Description 获取镜像列表
// @tags docker
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

//imageId string @Security ApiKeyAuth

// @Title dockerImage
// @Description 根据镜像id获取镜像信息
// @tags docker
// @Security ApiKeyAuth
// @Param imageId path string true "镜像ID"
// @Success 200 "获取镜像信息成功"
// @Router /images/getImageInfo/{imageId}  [get]
func GetImageInfo(c *gin.Context)  {

	fmt.Printf("%v", c.Param("imageId"))
	info, err := model.GetImageInfo(c.Param("imageId"))
	if err != nil {
		c.JSON(http.StatusOK, utils.RespError(http.StatusBadRequest, err))
		return
	}
	c.JSON(http.StatusOK, utils.RespData(http.StatusOK, info))

}


// @Title deleteDocker
// @Description 根据镜像id删除镜像
// @tags docker
// @Security ApiKeyAuth
// @Param imageId path string true "镜像ID"
// @Success 200 "删除镜像成功"
// @Router /images/deleteImage/{imageId} [get]
func DeleteImage(c *gin.Context)  {
	err := model.DeleteImage(c.Param("imageId"), true)
	if err != nil {
		c.JSON(http.StatusOK, utils.RespError(http.StatusBadRequest, err))
		return
	}
	c.JSON(http.StatusOK, utils.RespSuccess())
}


// @Title reTagImage
// @Description 重命名tag
// @tags docker
// @Param source query string true "镜像原tag"
// @Param tag query string true "镜像新tag"
// @Security ApiKeyAuth
// @Success 200 "重命名成功"
// @Router /images/reTagImage [get]
func ReTagImage(c *gin.Context)  {
	source := c.Request.URL.Query().Get("source")
	tag := c.Request.URL.Query().Get("tag")
	err := model.TagImage(source, tag)
	if err != nil {
		c.JSON(http.StatusOK, utils.RespError(http.StatusBadRequest, err))
		return
	}
	c.JSON(http.StatusOK, utils.RespSuccess())
}


// @Title 保存镜像到本地
// @Description 保存镜像到本地
// @tags docker
// @Security ApiKeyAuth
// @Success 200 "保存成功"
// @Router /images/saveImage [post]
func SaveImage(c *gin.Context)  {

}
//拉去镜像
// @Title 拉去镜像
// @Description 拉去镜像
// @tags docker
// @Security ApiKeyAuth
// @Success 200 "推送成功"
// @Router /images/pullImage [post]
func PullImage()  {

}

//推送镜像
// @Title 推送镜像
// @Description 推送镜像
// @tags docker
// @Security ApiKeyAuth
// @Success 200 "推送成功"
// @Router /images/pushImage [post]
func PushImage()  {

}

//导入镜像
func ImportImage()  {

}