package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/suifengtec/gin-csrf"
)

func main() {

	r := gin.Default()

	store := cookie.NewStore([]byte("boofarfar"))
	r.Use(sessions.Sessions("csrfsession", store))
	r.LoadHTMLGlob("templates/*")
	r.GET("/signin", func(c *gin.Context) {

		c.HTML(http.StatusOK, "signin.html", gin.H{
			"title": "登录",
			"token": csrf.GetToken(c),
		})
	})
	r.GET("/dashboard", csrf.MiddleWare(), func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"title":   "后台",
			"success": true,
		})
	})
	r.POST("/dashboard", csrf.MiddleWare(), func(c *gin.Context) {

		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"title":   "后台",
			"success": true,
		})

	})
	// http://127.0.0.1:8102
	r.Run(":8102")

}
