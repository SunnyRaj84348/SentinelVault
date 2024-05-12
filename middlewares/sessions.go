package middlewares

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Sessions() gin.HandlerFunc {
	store := cookie.NewStore([]byte(os.Getenv("AUTH_KEY")))
	store.Options(sessions.Options{
		MaxAge:   60 * 60 * 24 * 30,
		HttpOnly: true,
	})

	return sessions.Sessions("session_user", store)
}
