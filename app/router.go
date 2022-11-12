package app

import (
	"app/app/auth"
	"encoding/gob"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func IsAuthenticated(ctx *gin.Context) {
	profile := sessions.Default(ctx).Get("profile")
	if profile == nil {
		ctx.Redirect(http.StatusSeeOther, "/auth")
	} else {
		ctx.Next()
	}
}

// NewRouter registers the routes and returns the router.
func NewRouter(directory string, authenticator *auth.Authenticator) *gin.Engine {
	// gin.ReleaseMode = "true"
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))
	router.LoadHTMLGlob("template/*")

	router.GET("/auth", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	})
	router.GET("/login", auth.LoginHandler(authenticator))
	router.GET("/callback", auth.CallbackHandler(authenticator))

	router.Use(IsAuthenticated)
	router.GET("/logout", auth.LogoutHandler)

	files, _ := os.ReadDir(directory)

	for _, file := range files {
		name := file.Name()
		if strings.HasPrefix(name, ".") {
			continue
		}
		if file.IsDir() {
			router.Static("/"+name, fmt.Sprintf("%s/%s", directory, name))
		} else {
			router.StaticFile("/"+name, fmt.Sprintf("%s/%s", directory, name))
		}
	}

	router.StaticFile("/", fmt.Sprintf("%s/index.html", directory))
	return router
}
