package http

import (
	//"fmt"
	//"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/gin/binding"
	//"log"
	//"mavianceutililtyservice/pkg/util/json"
)

const (
	prefix           = "ddd-gin-admin"
	UserIDKey        = prefix + "/user-id"
	ReqBodyKey       = prefix + "/req-body"
	ResBodyKey       = prefix + "/res-body"
	LoggerReqBodyKey = prefix + "/logger-req-body"
)

func GetToken(c *gin.Context) string {

	var token string
	auth := c.GetHeader("Authorization")

	prefix := "Bearer "
	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}
	return token
}

func GetUserID(c *gin.Context) string {
	return c.GetString(UserIDKey)
}

func SetUserID(c *gin.Context, userID string) {
	c.Set(UserIDKey, userID)
}

func GetBody(c *gin.Context) []byte {
	if v, ok := c.Get(ReqBodyKey); ok {
		if b, ok := v.([]byte); ok {
			return b
		}
	}
	return nil
}