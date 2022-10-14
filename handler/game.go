package handler

import (
	"net/http"
	"test-coding/game"
	"test-coding/helper"

	"github.com/gin-gonic/gin"
)

type gameHandler struct {
	gameService game.Service
}

func NewGameHandler(gameService game.Service) *gameHandler {
	return &gameHandler{gameService}
}

func (h *gameHandler) InsertData(c *gin.Context) {
	var input game.InputDataGame

	err := c.ShouldBind(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("insert data failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newData, err := h.gameService.InserData(input)
	if err != nil {
		response := helper.APIResponse("insert data failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formater := game.FormatDataGame(newData)
	response := helper.APIResponse("Data Game Successfully inserted", http.StatusOK, "success", formater)
	c.JSON(http.StatusOK, response)

}

func (h *gameHandler) ShowData(c *gin.Context) {

	data, err := h.gameService.ShowDataGame()
	if err != nil {
		response := helper.APIResponse("Error to get data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// response := helper.APIResponse("list of campaigns", http.StatusOK, "ok", data)
	c.JSON(http.StatusOK, data)

}
