package controller

import (
	"net/http"
	"usecase-1/delivery/middleware"
	"usecase-1/model"
	"usecase-1/usecase"
	"usecase-1/utils/common"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Usecase2Controller struct {
	usecase2 usecase.Usecase2UseCase
	router   *gin.Engine
}

func (e *Usecase2Controller) registerHandler(c *gin.Context) {
	var u2s model.Usecase2RegisterModel
	if err := c.ShouldBindJSON(&u2s); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	u2s.ID = common.GenerateID()
	hash, err := bcrypt.GenerateFromPassword([]byte(u2s.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(400, gin.H{"err": err})
		return
	}
	u2s.Password = string(hash)
	data, err := e.usecase2.RegisterNewU1(u2s)
	if err != nil {
		c.JSON(400, gin.H{"err": err})
		return
	}
	c.JSON(201, data)
}

func (e *Usecase2Controller) loginHandler(c *gin.Context) {
	var payload model.Usecase2LoginModel
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	token, err := e.usecase2.LoginNewU1(payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"token": token})
}

func (e *Usecase2Controller) getProfileHandler(c *gin.Context) {
	// Ambil klaim dari context
	claims, exists := c.Get("userClaims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		c.Abort()
		return
	}

	userClaims := claims.(jwt.MapClaims)
	tokenUsername, _ := userClaims["username"].(string) // Dapatkan username dari klaim token

	username := c.Param("username")

	// Periksa apakah pengguna mencoba mengakses profil mereka sendiri
	if tokenUsername != username {
		c.JSON(http.StatusForbidden, gin.H{"message": "forbidden"})
		c.Abort()
		return
	}

	// Ambil data pengguna berdasarkan username
	data := e.usecase2.FindByUsernameU2(username)
	if data.Username == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status": gin.H{
				"code":        http.StatusNotFound,
				"description": "Get Data By Username: " + username + " Not Found",
			},
			"data": username,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": gin.H{
			"code":        http.StatusOK,
			"description": "Get Data By Username: " + username + " Successfully",
		},
		"data": data,
	})
}

func NewU2Controller(usecase usecase.Usecase2UseCase, r *gin.Engine) *Usecase2Controller {
	controller := Usecase2Controller{
		router:   r,
		usecase2: usecase,
	}

	// Grup rute untuk pengguna yang tidak membutuhkan autentikasi
	rg := r.Group("/user")
	rg.POST("/register", controller.registerHandler)
	rg.POST("/login", controller.loginHandler)

	// Grup rute yang membutuhkan autentikasi
	authRg := rg.Group("/profil/")
	authRg.Use(middleware.AuthMiddleware())
	authRg.GET("/:username", controller.getProfileHandler)

	return &controller
}
