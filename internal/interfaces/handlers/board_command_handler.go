package handlers

import (
	"fmt"
	"net/http"
	"task-management-api/internal/domain/models"
	"task-management-api/internal/usecases/services"

	"github.com/gin-gonic/gin"
)

type BoardCommandHandlerI interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type boardCommandHandlerImpl struct {
	service services.BoardCommandServiceI
}

func NewBoardCommandHandler(service services.BoardCommandServiceI) BoardCommandHandlerI {
	return &boardCommandHandlerImpl{
		service: service,
	}
}

func (h *boardCommandHandlerImpl) Create(c *gin.Context) {
	var board models.Board

	// Primero mapeo el JSON para que funcione con el modelo
	if err := c.ShouldBindJSON(&board); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.Create(c, &board)

	if err != nil {
		fmt.Println("error", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tablero Recibido",
		"board":   board,
	})
}

func (h *boardCommandHandlerImpl) Update(c *gin.Context) {
	var board models.Board

	// Primero mapeo el JSON con el modelo
	if err := c.ShouldBindJSON(&board); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.Update(c, &board)

	if err != nil {
		fmt.Println("error", err)
		return
	}

	// Responde con un mensaje
	c.JSON(200, gin.H{
		"message": "board actualizado con exito",
	})
}

func (h *boardCommandHandlerImpl) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(c, id)

	if err != nil {
		fmt.Println("error", err)
	}

	// Responde con un mensaje
	c.JSON(200, gin.H{
		"message": "creadoo con exito",
	})
}
