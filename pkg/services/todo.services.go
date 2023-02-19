package services

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"

	"github.com/GaurKS/crud-app/pkg/dtos"
	"github.com/GaurKS/crud-app/pkg/models"
	"github.com/gin-gonic/gin"
)

func (db Handler) CreateTodo(c *gin.Context) {
	newTodo := dtos.CreateTodo{}

	if err := c.BindJSON(&newTodo); err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		fmt.Println(err.Error())
		c.Abort()
		return
	}

	todo := models.Todo{}
	todo.Title = newTodo.Title
	todo.TodoStatus = newTodo.TodoStatus
	todo.Description = newTodo.Description
	todo.CreatedBy = newTodo.CreatedBy

	if result := db.DB.Model(&todo).Create(&todo); result.Error != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": result.Error,
			},
		)
		fmt.Println(result.Error)
		c.Abort()
		return
	}

	c.IndentedJSON(
		http.StatusCreated,
		gin.H{
			"data": &todo,
		},
	)
}

func (db Handler) GetAllTodos(c *gin.Context) {
	var todos []models.Todo
	if result := db.DB.Find(&todos); result.Error != nil {
		fmt.Println(result.Error)
	}
	c.IndentedJSON(
		http.StatusOK,
		gin.H{
			"message": "SUCCESS",
			"resource": &todos,
		},
	)
}

func (db Handler) GetTodoById(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := db.DB.First(&todo, id); err.Error != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{
				"error": err.Error,
			},
		)
		c.Abort()
		return
	}
	c.IndentedJSON(
		http.StatusOK,
		gin.H{
			"data": &todo,
		},
	)
}

func (db Handler) UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	editTodo := dtos.CreateTodo{}
	var todo models.Todo

	if result := db.DB.First(&todo, id); result.Error != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{
				"error": result.Error,
			},
		)
		c.Abort()
		return
	}

	if err := c.ShouldBindJSON(&editTodo); err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		fmt.Println(err.Error())
		c.Abort()
		return
	}

	todo.Title = editTodo.Title
	todo.TodoStatus = editTodo.TodoStatus
	todo.Description = editTodo.Description
	todo.CreatedBy = editTodo.CreatedBy

	if result := db.DB.Save(&todo); result.Error != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": result.Error,
			},
		)
		fmt.Println(result.Error)
		c.Abort()
		return
	}

	c.IndentedJSON(
		http.StatusCreated,
		gin.H{
			"data": &todo,
		},
	)
}

func (db Handler) DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	var todo models.Todo

	if result := db.DB.First(&todo, id); result.Error != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{
				"error": result.Error,
			},
		)
		c.Abort()
		return
	}

	if result := db.DB.Delete(&todo); result.Error != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": result.Error,
			},
		)
		fmt.Println(result.Error)
		c.Abort()
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		gin.H{
			"message": "SUCCESS",
		},
	)
}

func (db Handler) ParseCsv(c *gin.Context) {
	file, err := c.FormFile("csv")
	if err != nil {
		c.IndentedJSON(
			http.StatusBadRequest, 
			gin.H {
				"error": "File not found",
			},
		)
		c.Abort()
		return
	}

	f, err := file.Open()
	if err != nil {
		c.IndentedJSON(
			http.StatusBadRequest, 
			gin.H {
				"error": "Failed to open file",
			},
		)
		c.Abort()
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)
	var todos []dtos.CreateTodo
	// todoModel := models.Todo{} 

	for {
		row, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			c.IndentedJSON(
				http.StatusBadRequest, 
				gin.H {
					"error": "Failed to parse CSV file",
				},
			)
			c.Abort()
			return
		}

		todo := dtos.CreateTodo{
			Title: row[0],
			TodoStatus: row[1],
			Description: row[2],
			CreatedBy: row[3],
		}

		todos = append(todos, todo)
	}

	var builder strings.Builder
	builder.WriteString("+--------------------------+------------------+------------------------------------+--------------------+\n")
	for _, record := range todos {
			builder.WriteString(fmt.Sprintf("| %-24s | %-16s | %-34s | %-18s |\n", 
				record.Title, 
				record.TodoStatus, 
				record.Description, 
				record.CreatedBy,
			))
	}
	builder.WriteString("+--------------------------+------------------+------------------------------------+--------------------+\n")

	// Return the formatted table
	c.String(http.StatusOK, builder.String())
}