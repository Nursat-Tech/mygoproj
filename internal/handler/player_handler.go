package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nursat/myproj/internal/models"
	"github.com/nursat/myproj/internal/service"
)

type PlayerHandler struct {
	service *service.PlayerService
}

func NewPlayerHandler(service *service.PlayerService) *PlayerHandler {
	return &PlayerHandler{service: service}
}

func (h *PlayerHandler) CreatePlayer(c *gin.Context) {
	var player models.Player
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreatePlayer(&player); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create player"})
		return
	}
	c.JSON(http.StatusCreated, player)
}

func (h *PlayerHandler) GetPlayers(c *gin.Context) {
	players, err := h.service.GetPlayers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get players"})
		return
	}
	c.JSON(http.StatusOK, players)
}

func (h *PlayerHandler) GetPlayerByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}
	player, err := h.service.GetPlayerByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}
	c.JSON(http.StatusOK, player)
}

func (h *PlayerHandler) UpdatePlayer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}
	var player models.Player
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	player.ID = uint(id)
	if err := h.service.UpdatePlayer(&player); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot update player"})
		return
	}
	c.JSON(http.StatusOK, player)
}

func (h *PlayerHandler) DeletePlayer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}
	if err := h.service.DeletePlayer(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot delete player"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Player deleted"})
}
