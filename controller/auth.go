package controller

import (
	"dobo/database"
	"dobo/database/model"
	"dobo/middleware"
	"dobo/utils"
	"net"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")


type RegisterPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserName string `json:"user_name"`
}
// @Title 注册
// @Description 用户注册
// @Security ApiKeyAuth
// @tags auth
// @Param auth body model.Auth true "用户信息"
// @Success 201 "注册成功"
// @Router /register [post]
func CreateUserAuth(c *gin.Context) {
	db := database.GetDB()
	//var registerPayload RegisterPayload
	auth := model.Auth{}

	createAuth := 0 // default value

	auth.AuthID = middleware.AuthID

	err := c.ShouldBindJSON(&auth)
	if err != nil {
		c.JSON(http.StatusOK, utils.RespErrorMsg(500,"数据格式错误"))
		return
	}

	if isEmailValid(auth.Email) == false {
		createAuth = 1 // invalid email
		c.JSON(http.StatusOK, utils.RespErrorMsg(http.StatusBadRequest,"邮箱格式错误"))
		utils.Logger().Debug("邮箱格式错误")
		return
	}

	if err := db.Where("email = ?", auth.Email).First(&auth).Error; err == nil {
		createAuth = 2 // email is already registered
		c.JSON(http.StatusOK, utils.RespErrorMsg(http.StatusBadRequest,"用户名已存在"))
		utils.Logger().Debug("用户名已存在")
		return
	}

	// one unique email for each account
	if createAuth == 0 {
		tx := db.Begin()
		if err := tx.Create(&auth).Error; err != nil {
			tx.Rollback()
			utils.Logger().Error(err)
			c.JSON(http.StatusOK, utils.RespError(http.StatusInternalServerError,err))
		} else {
			tx.Commit()
			utils.Logger().Info("创建用户成功")
			c.JSON(http.StatusCreated, utils.RespData(http.StatusCreated,auth))
		}
	}
}

// isEmailValid checks if the email provided passes the required structure
// and length test. It also checks the domain has a valid MX record.
// Credit: Edd Turtle
func isEmailValid(e string) bool {
	if len(e) < 3 || len(e) > 254 {
		return false
	}

	if !emailRegex.MatchString(e) {
		return false
	}

	parts := strings.Split(e, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return false
	}

	return true
}
