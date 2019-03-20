package csrf

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

//validateCsrfToken ...
func validateCsrfToken(c *gin.Context) bool {

	tokenInCookie, err := c.Cookie("csrf_token")
	if err != nil {
		return false
	}
	tokenInRequest := c.PostForm("csrf_token")
	if tokenInRequest == "" {
		tokenInRequest = c.Query("csrf_token")
	}
	if tokenInRequest == "" || tokenInCookie == "" {
		return false
	}

	if tokenInRequest != tokenInCookie {
		return false
	}

	return true
}

//GetToken ...
func GetToken(c *gin.Context) string {
	r := rand.Int()
	t := time.Now().UnixNano()
	token := fmt.Sprintf("%d%d", r, t)
	c.SetCookie("csrf_token", token, 2*60, "/", "", false, false)
	c.Set("csrf_token", token)

	return token

}

//CsrfMiddleWare ...
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//if c.Request.Method == "POST" {}
		if !validateCsrfToken(c) {
			//may need to redirect or output some error message
			c.Abort()
		} else {
			c.Next()
		}
	}
}
