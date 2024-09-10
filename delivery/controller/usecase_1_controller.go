package controller

import (
	"strconv"
	"usecase-1/model"
	"usecase-1/usecase"

	"github.com/gin-gonic/gin"
)

type Usecase1Controller struct {
	usecase1 usecase.Usecase1UseCase
	router   *gin.Engine
}

func (e *Usecase1Controller) createHandler(c *gin.Context) {
	var u1s model.Usecase1Model
	if err := c.ShouldBindJSON(&u1s); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	data, err := e.usecase1.RegisterNewU1(u1s)
	if err != nil {
		c.JSON(400, gin.H{"err": err})
		return
	}
	c.JSON(201, data)
}

func (e *Usecase1Controller) listHandler(c *gin.Context) {
	u1s := e.usecase1.FindAllU1()
	status := map[string]any{
		"code":        200,
		"description": "Get All Data Successfully",
	}
	c.JSON(200, gin.H{
		"status": status,
		"data":   u1s,
	})
}

func (e *Usecase1Controller) listByIdHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	data := e.usecase1.FindByIdU1(id)
	if data.ID == 0 {
		status := map[string]any{
			"code":        404,
			"description": "Get Data By ID:" + strconv.Itoa(id) + " Not Found",
		}
		c.JSON(200, gin.H{
			"status": status,
			"data":   id,
		})
	} else {
		status := map[string]any{
			"code":        200,
			"description": "Get Data By ID:" + strconv.Itoa(id) + " Successfully",
		}
		c.JSON(200, gin.H{
			"status": status,
			"data":   data,
		})
	}

}

func (e *Usecase1Controller) deleteHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	status := map[string]any{
		"code":        200,
		"description": e.usecase1.DeleteByIdU1(id),
	}
	c.JSON(200, gin.H{
		"status": status,
	})
}

func NewU1Controller(usecase usecase.Usecase1UseCase, r *gin.Engine) *Usecase1Controller {
	controller := Usecase1Controller{
		router:   r,
		usecase1: usecase,
	}

	rg := r.Group("/task/")
	rg.POST("/", controller.createHandler)
	rg.GET("/", controller.listHandler)
	rg.GET("/:id", controller.listByIdHandler)
	rg.DELETE("/:id", controller.deleteHandler)
	return &controller
}
