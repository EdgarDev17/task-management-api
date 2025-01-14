package handlers

import (
	"fmt"
	"task-management-api/internal/usecases/services"

	"github.com/gin-gonic/gin"
)

type TaskQueryHandlerI interface {
	GetById(c *gin.Context)
	GetAll(c *gin.Context)
}

// las implementaciones deben ser privadas hacia otros paquetes
type taskQueryHandlerImpl struct {
	service services.TaskQueryServiceI
}

func NewTaskQueryHandler(service services.TaskQueryServiceI) TaskQueryHandlerI {
	return &taskQueryHandlerImpl{
		service: service,
	}
}

func (h *taskQueryHandlerImpl) GetAll(c *gin.Context) {
	tasks, err := h.service.GetAll(c)

	if err != nil {
		fmt.Println("error", err)
	}

	// Responde con un mensaje
	c.JSON(200, gin.H{
		"error": nil,
		"Tasks": tasks,
	})
}

func (h *taskQueryHandlerImpl) GetById(c *gin.Context) {
	id := c.Param("id")

	task, err := h.service.GetById(c, id)

	if err != nil {
		fmt.Println("error", err)
	}
	// Responde con un mensaje
	c.JSON(200, gin.H{
		"error": nil,
		"Task":  task,
	})
}
