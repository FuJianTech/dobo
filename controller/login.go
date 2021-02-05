package controller

import (
	"dobo/database/model"
	"dobo/middleware"
	"dobo/service"
	"dobo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)




// LoginPayload ...
type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login ...
// @Title 登录
// @Description 用户登录
// @tags auth
// @Param auth body LoginPayload true "Email&Password"
// @Success 200
// @Router /login [post]
func Login(c *gin.Context) {
	var payload LoginPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusOK, utils.RespErrorMsg(http.StatusBadRequest,"用户名或密码错误"))
		utils.Logger().Error(err)
		//log.Error(err)
		return
	}
	v, err := service.GetUserByEmail(payload.Email)
	if err != nil {
		c.JSON(http.StatusOK, utils.RespErrorMsg(http.StatusBadRequest,"用户名错误"))
		utils.Logger().Error(err)
		return
	}

	if v.Password != model.HashPass(payload.Password) {
		c.JSON(http.StatusOK, utils.RespErrorMsg(http.StatusBadRequest,"密码错误"))
		utils.Logger().Error(err)
		return
	}

	jwtValue, err := middleware.GetJWT(v.AuthID, v.Email)
	if err != nil {
		c.JSON(http.StatusOK,utils.RespErrorMsg(http.StatusInternalServerError, "服务器内部错误"))
		utils.Logger().Error(err)
		return
	}
	data := make(map[string]interface{})
	data["Token"] = jwtValue
	c.JSON(http.StatusOK, utils.RespData(http.StatusOK,data))

}



func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true)

}
