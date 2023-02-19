package routes

import (
	"github.com/GaurKS/crud-app/pkg/services"
	"github.com/gin-gonic/gin"
)

func TodoRouter(r *gin.RouterGroup, h *services.Handler) {
	r.POST("/create", h.CreateTodo)
	r.GET("/read/all", h.GetAllTodos)
	r.GET("/read/:id", h.GetTodoById)
	r.PATCH("/update/:id", h.UpdateTodo)
	r.DELETE("/delete/:id", h.DeleteTodo)
	r.POST("/parse/csv", h.ParseCsv)
	r.GET("/health", h.HealthCheck)
}