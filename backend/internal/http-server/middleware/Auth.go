package middleware

import (
	"time"

	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var identityKey = "key"

func GetAuthMW(targetUsername, targetPassword string) (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "admin zone",
		Key:         []byte("secret"),
		Timeout:     time.Hour * 30 * 24,
		MaxRefresh:  time.Hour * 30 * 24,
		IdentityKey: identityKey,

		IdentityHandler: func(c *gin.Context) interface{} {
			claims, exists := c.Get("JWT_PAYLOAD")
			if exists {
				if claimMap, ok := claims.(jwt.MapClaims); ok {
					return &Login{Username: claimMap["username"].(string)}
				}
			}
			return nil
		},

		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals Login
			if err := c.ShouldBindJSON(&loginVals); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}
			username := loginVals.Username
			password := loginVals.Password

			if username == targetUsername && password == targetPassword {
				return &Login{Username: username}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},

		Authorizator: func(data interface{}, c *gin.Context) bool {
			if user, ok := data.(*Login); ok && user.Username == targetUsername {
				return true
			}
			return false
		},

		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{"error": message})
		},

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if user, ok := data.(*Login); ok {
				return jwt.MapClaims{
					"username": user.Username,
					"exp":      time.Now().Add(4 * time.Hour).Unix(),
					"iat":      time.Now().Unix(),
				}
			}
			return jwt.MapClaims{}
		},

		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",

		TimeFunc: time.Now,
	})
}
