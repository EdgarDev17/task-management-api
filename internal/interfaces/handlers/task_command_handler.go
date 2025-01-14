package handlers

import (
	"fmt"
	"net/http"
	"task-management-api/internal/domain/models"
	"task-management-api/internal/usecases/services"

	"github.com/gin-gonic/gin"
)

type TaskCommandHandlerI interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type taskCommandHandlerImpl struct {
	service services.TaskCommandServiceI
}

func NewTaskCommandHandler(service services.TaskCommandServiceI) TaskCommandHandlerI {

	return &taskCommandHandlerImpl{
		service: service,
	}
}

func (h *taskCommandHandlerImpl) Create(c *gin.Context) {
	var task models.Task

	// Primero mapeo el JSON para que funcione con el modelo
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.Create(c, &task)

	if err != nil {
		fmt.Println("error", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "tarea Recibido",
		"task":    task,
	})
}

func (h *taskCommandHandlerImpl) Update(c *gin.Context) {
	var task models.Task

	// Primero mapeo el JSON con el modelo
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.Update(c, &task)

	if err != nil {
		fmt.Println("error", err)
		return
	}

	// Responde con un mensaje
	c.JSON(200, gin.H{
		"message": "task actualizado con exito",
	})
}

func (h *taskCommandHandlerImpl) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(c, id)

	if err != nil {
		fmt.Println("error", err)
	}

	// Responde con un mensaje
	c.JSON(200, gin.H{
		"message": "eliminado con exito",
	})
}
