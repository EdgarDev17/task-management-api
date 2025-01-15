package handlers

import (
	"fmt"
	"net/http"
	"task-management-api/internal/domain/models"
	"task-management-api/internal/infrastructure/delivery/response"
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
		c.JSON(http.StatusBadRequest, &response.ResponseError{
			Code:    http.StatusBadRequest,
			Message: "Datos invalidos",
		})
		return
	}

	newBoard, err := h.service.Create(c, &board)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &response.ResponseError{
			Code:    http.StatusInternalServerError,
			Message: "Error al crear el tablero: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &response.ResponseSuccess{
		Code:    http.StatusOK,
		Message: "Tablero creado con exito",
		Data:    newBoard,
	})
}

func (h *boardCommandHandlerImpl) Update(c *gin.Context) {
	var board models.Board

	// Primero mapeo el JSON con el modelo
	if err := c.ShouldBindJSON(&board); err != nil {
		c.JSON(http.StatusInternalServerError, &response.ResponseError{
			Code:    http.StatusInternalServerError,
			Message: "Error al actualizar el tablero: " + err.Error(),
		})
		return
	}

	err := h.service.Update(c, &board)

	if err != nil {
		fmt.Println("error", err)
		return
	}

	// Responde con un mensaje
	// Primero mapeo el JSON con el modelo
	if err := c.ShouldBindJSON(&board); err != nil {
		c.JSON(http.StatusInternalServerError, &response.ResponseSuccess{
			Code:    http.StatusOK,
			Message: "Tablero actualizado con exito",
			Data:    board.ID,
		})
		return
	}

}

func (h *boardCommandHandlerImpl) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(c, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &response.ResponseError{
			Code:    http.StatusInternalServerError,
			Message: "Error al eliminar el tablero: " + err.Error(),
		})
	}

	// Responde con un mensaje
	c.JSON(http.StatusOK, &response.ResponseSuccess{
		Code:    http.StatusOK,
		Message: "Tablero eliminado con exito",
	})
}
