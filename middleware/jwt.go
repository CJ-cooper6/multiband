package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"multiband/errmsg"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte("123")

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 定义错误
var (
	TokenExpired     error = errors.New("Token已过期,请重新登录")
	TokenNotValidYet error = errors.New("Token无效,请重新登录")
	TokenMalformed   error = errors.New("Token不正确,请重新登录")
	TokenInvalid     error = errors.New("这不是一个token,请重新登录")
)

func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			code := errmsg.Error_Token_Exist
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.Geterrmsg(code),
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			code := errmsg.Error_Token_Type_Wrong
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.Geterrmsg(code),
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它

		// 解析token
		claims, err := CheckToken(parts[1])
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"status":  errmsg.Error,
					"message": "token授权已过期,请重新登录",
					"data":    nil,
				})
				c.Abort()
				return
			}
			// 其他错误
			c.JSON(http.StatusOK, gin.H{
				"status":  errmsg.Error,
				"message": err.Error(),
				"data":    nil,
			})
			c.Abort()
			return
		}

		c.Set("username", claims)
		c.Next()
	}
}

//生成Token
func SetToken(username string) (string, int) {

	expireTime := time.Now().Add(10 * time.Hour)
	SetClaims := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 过期时间
			Issuer:    "my-project",      // 签发人
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := t.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.Error
	}
	return token, errmsg.Success

}

//验证Token

func CheckToken(token string) (*MyClaims, error) {
	// 解析token
	settoken, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return JwtKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if settoken != nil {
		if claims, ok := settoken.Claims.(*MyClaims); ok && settoken.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	}

	return nil, TokenInvalid
}
