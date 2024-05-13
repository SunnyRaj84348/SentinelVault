package middlewares

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth(ctx *gin.Context) {
	session := sessions.Default(ctx)

	userid := session.Get("user")
	if userid == nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("userid", userid)
}
