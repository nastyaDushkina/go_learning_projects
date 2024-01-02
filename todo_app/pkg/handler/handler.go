package handler

import (
	"todo_app/pkg/service"

	"github.com/gin-gonic/gin"

	_ "todo_app/docs" // пустой импорт

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// имплементируем хэндлеры

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// web http framework
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// router.Run(":8080")

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIndentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList) // добавили методы хэндлеров
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById) // любое значение, к которому обращаемся по имени параметра
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
			}
		}

		// для этих запросов не нужен list_id. Поэтому выделяем их в отдельную группу
		items := api.Group("/items")
		{
			items.GET("/:id", h.getItemById)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
		}
	}

	return router
}
