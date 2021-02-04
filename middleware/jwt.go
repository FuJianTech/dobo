package middleware

import (
	"dobo/utils"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// MySigningKey ...
var MySigningKey []byte

// JWTExpireTime ...
var JWTExpireTime int

// MyCustomClaims ...
type MyCustomClaims struct {
	ID    uint   `json:"Id"`
	Email string `json:"Email"`
	jwt.StandardClaims
}

// AuthID - Access details
var AuthID uint

// JWT ...
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		val := c.Request.Header.Get("Authorization")
		if len(val) == 0  {
			//log.Println("no Authorization found")
			utils.Logger().Error("no Authorization found")
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.RespErrorMsg(http.StatusUnauthorized,"no Authorization found"))
			return
		}

		token, err := jwt.ParseWithClaims(val, &MyCustomClaims{}, validateJWT)

		if err != nil {
			utils.Logger().Error("error parsing JWT", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.RespErrorMsg(http.StatusUnauthorized,"Token错误"))
			return
		}
		if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
			AuthID = claims.ID
		}
	}
}

// validateJWT ...
func validateJWT(token *jwt.Token) (interface{}, error) {
	//log.Println("try to parse the JWT")
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		//log.Println("error parsing JWT")
		utils.Logger().Error("error parsing JWT", ok)
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return MySigningKey, nil
}

// GetJWT ...
func GetJWT(id uint, email string) (string, error) {
	// Create the Claims
	claims := MyCustomClaims{
		id,
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(JWTExpireTime)).Unix(),
			Issuer:    "GoRest API",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtValue, err := token.SignedString(MySigningKey)
	if err != nil {
		return "", err
	}
	return jwtValue, nil
}
