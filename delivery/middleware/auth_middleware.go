package middleware

import (
	"net/http"
	"strings"
	"usecase-1/utils/security"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	Authorization string `header:"Authorization"` // Pastikan tag ini sesuai dengan header yang dikirimkan
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var h authHeader
		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "authorization header is required"})
			c.Abort()
			return
		}

		// Ambil token dari header Authorization dan hapus "Bearer "
		if !strings.HasPrefix(h.Authorization, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "authorization header must start with 'Bearer '"})
			c.Abort()
			return
		}
		tokenHeader := strings.TrimSpace(strings.Replace(h.Authorization, "Bearer ", "", 1))

		if tokenHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "token is required"})
			c.Abort()
			return
		}

		// Verifikasi token dan pastikan token valid
		claims, err := security.VerifyAccessToken(tokenHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid or expired token"})
			c.Abort()
			return
		}

		// Pastikan klaim token mengandung 'username'
		if username, ok := claims["username"]; !ok || username == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token payload"})
			c.Abort()
			return
		}

		// Simpan klaim token di context untuk digunakan di handler berikutnya
		c.Set("userClaims", claims)

		// Lanjutkan ke handler berikutnya
		c.Next()
	}
}
