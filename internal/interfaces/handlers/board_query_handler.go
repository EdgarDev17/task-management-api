package handlers

import (
	"fmt"
	"task-management-api/internal/usecases/services"

	"github.com/gin-gonic/gin"
)

type BoardQueryHandlerI interface {
	GetById(c *gin.Context)
	GetAll(c *gin.Context)
}

// las implementaciones deben ser privadas hacia otros paquetes
type boardQueryHandlerImpl struct {
	service services.BoardQueryServiceI
}

func NewBoardHandler(service services.BoardQueryServiceI) BoardQueryHandlerI {
	return &boardQueryHandlerImpl{
		service: service,
	}
}

func (h *boardQueryHandlerImpl) GetAll(c *gin.Context) {
	boards, err := h.service.GetAll(c)

	if err != nil {
		fmt.Println("error", err)
	}

	// Responde con un mensaje
	c.JSON(200, gin.H{
		"message": "Welcome to the GetById!",
		"boards":  boards,
	})
}

func (h *boardQueryHandlerImpl) GetById(c *gin.Context) {
	id := c.Param("id")

	board, err := h.service.GetById(c, id)

	if err != nil {
		fmt.Println("error", err)
	}

	// Responde con un mensaje
	c.JSON(200, gin.H{
		"message": "Welcome to the GetById!",
		"boardId": id,
		"board":   board,
	})
}
