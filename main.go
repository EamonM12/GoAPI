package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// building REST API

type todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Extra     string `json:"extra"`
	Completed bool   `json:"completed"`
	Date      string `json:"date"`
}

var todos = []todo{
	{ID: "1", Title: "Clean room", Body: "Make bed aswell as clear my desk", Date: "20/12/2020", Completed: false},
	{ID: "2", Title: "Do the dishes", Body: "Load the dishwasher and clean up cooking utensiles", Date: "21/2/2021", Completed: false},
	{ID: "3", Title: "Go shopping", Body: "Grab cereal, pasta and also some more beer", Date: "12/3/2022", Completed: false},
}

func get_todos(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, todos)
}

func getTodoID(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("Todo not found")
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func NewTodo(context *gin.Context) {
	var newtodo todo

	if err := context.BindJSON(&newtodo); err != nil {
		return
	}
	todos = append(todos, newtodo)
	context.IndentedJSON(http.StatusCreated, newtodo)

}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

// func deletetodos(c *gin.Context) {
// 	id := c.Param("id")
// 	todo, err := getTodoID(id)

// 	if err != nil {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
// 		return
// 	}

// 	todos = remove(todos, todo)

// }

//cors

func main() {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"DELETE", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))
	// router.DELETE("/todos/:id", deletetodos)
	router.GET("/todos", get_todos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", NewTodo)
	router.Run("localhost:9090")

}
