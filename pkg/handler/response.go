package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	// поскольку у ручки может быть несколько
	// обработчиков (хэндлеров) ндо прервать выполнение последующих
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
