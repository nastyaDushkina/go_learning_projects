package handler

import (
	"net/http"
	todo "todo_app"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signIn(c *gin.Context) {
	// регистрация

}

func (h *Handler) signUp(c *gin.Context) {
	// аутентификация
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// далее мы должны передать валидированные данные на слой ниже в сервисы (репозиторий)
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
